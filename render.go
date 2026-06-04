package sintax

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/toaweme/sintax/functions"
)

type Arg struct {
	// Value is the value of the argument, which can be a literal value or a variable name
	Value any
	// Var tells us if the value is a variable name or a literal value
	Var bool
}

type Func struct {
	Name string
	Args []Arg
}

// StringRenderer handles rendering templates with a given context.
type StringRenderer struct {
	funcs map[string]GlobalModifier
}

var _ Renderer = (*StringRenderer)(nil)

// NewStringRenderer creates a new instance of StringRenderer with the provided context.
func NewStringRenderer(funcs map[string]GlobalModifier) *StringRenderer {
	return &StringRenderer{
		funcs: funcs,
	}
}

// Render processes the provided tokens and variables, returning a rendered string or any value
func (r *StringRenderer) Render(tokens []Token, vars map[string]any) (any, error) {
	out, _, err := r.renderRange(tokens, 0, len(tokens), vars, true)
	return out, err
}

// renderRange renders tokens[start:end] with the given vars. returns the rendered
// value, the index at which rendering stopped (one past the last token consumed),
// and any error.
//
// if allowDirect is true and the only meaningful token in the range is a single
// non-string variable, the value is returned directly (legacy passthrough
// behaviour). otherwise everything is stringified.
func (r *StringRenderer) renderRange(tokens []Token, start, end int, vars map[string]any, allowDirect bool) (any, int, error) {
	var str strings.Builder
	i := start
	for i < end {
		token := tokens[i]
		switch token.Type() {
		case TextToken:
			str.WriteString(token.Raw())
			i++
		case VariableToken, FilteredVariableToken:
			variable, err := r.renderVariable(token, vars)
			if err != nil {
				return nil, i, fmt.Errorf("failed to render variable token '%s': %w", token.Name(), err)
			}
			if val, ok := variable.(string); ok {
				str.WriteString(val)
				i++
				continue
			}
			if val, ok := variable.(bool); ok {
				str.WriteString(fmt.Sprintf("%t", val))
				i++
				continue
			}
			// passthrough: a single non-string variable returns the value directly
			if allowDirect && start == i && i+1 == end && str.Len() == 0 {
				return variable, i + 1, nil
			}
			str.WriteString(fmt.Sprint(variable))
			i++
		case IfToken:
			out, next, err := r.renderIf(tokens, i, end, vars)
			if err != nil {
				return nil, i, err
			}
			str.WriteString(out)
			i = next
		case ForToken:
			out, next, err := r.renderFor(tokens, i, end, vars)
			if err != nil {
				return nil, i, err
			}
			str.WriteString(out)
			i = next
		case ElseToken, IfEndToken, ForEndToken:
			// caller should have stopped before this; reaching here means a stray closer
			return nil, i, fmt.Errorf("unexpected control token: %s", controlName(token.Type()))
		default:
			i++
		}
	}
	return str.String(), i, nil
}

func controlName(t TokenType) string {
	switch t {
	case IfToken:
		return "if"
	case ElseToken:
		return "else"
	case IfEndToken:
		return "endif"
	case ForToken:
		return "for"
	case ForEndToken:
		return "endfor"
	}
	return "?"
}

// findIfEnd locates the matching `endif` for the IfToken at index `start`. it also
// records the index of the (top-level) `else` if one exists. nested ifs are
// counted correctly.
func findIfEnd(tokens []Token, start, end int) (elseIdx, endIdx int, err error) {
	elseIdx = -1
	depth := 0
	for j := start + 1; j < end; j++ {
		switch tokens[j].Type() {
		case IfToken:
			depth++
		case IfEndToken:
			if depth == 0 {
				return elseIdx, j, nil
			}
			depth--
		case ElseToken:
			if depth == 0 && elseIdx == -1 {
				elseIdx = j
			}
		}
	}
	return -1, -1, fmt.Errorf("unterminated if block (missing endif)")
}

// findForEnd locates the matching `endfor` for the ForToken at index `start`.
func findForEnd(tokens []Token, start, end int) (int, error) {
	depth := 0
	for j := start + 1; j < end; j++ {
		switch tokens[j].Type() {
		case ForToken:
			depth++
		case ForEndToken:
			if depth == 0 {
				return j, nil
			}
			depth--
		}
	}
	return -1, fmt.Errorf("unterminated for block (missing endfor)")
}

func (r *StringRenderer) renderIf(tokens []Token, start, end int, vars map[string]any) (string, int, error) {
	elseIdx, endIdx, err := findIfEnd(tokens, start, end)
	if err != nil {
		return "", start, err
	}
	cond, err := r.evalCondition(tokens[start].Raw(), vars)
	if err != nil {
		return "", start, err
	}
	var bodyStart, bodyEnd int
	if cond {
		bodyStart = start + 1
		if elseIdx >= 0 {
			bodyEnd = elseIdx
		} else {
			bodyEnd = endIdx
		}
	} else if elseIdx >= 0 {
		bodyStart = elseIdx + 1
		bodyEnd = endIdx
	} else {
		return "", endIdx + 1, nil
	}
	out, _, err := r.renderRange(tokens, bodyStart, bodyEnd, vars, false)
	if err != nil {
		return "", start, err
	}
	s, _ := out.(string)
	return s, endIdx + 1, nil
}

func (r *StringRenderer) renderFor(tokens []Token, start, end int, vars map[string]any) (string, int, error) {
	endIdx, err := findForEnd(tokens, start, end)
	if err != nil {
		return "", start, err
	}
	tok := tokens[start]
	spec := tok.Name()
	expr := tok.LoopExpr()
	if spec == "" || expr == "" {
		return "", start, fmt.Errorf("invalid for expression: %q", tok.Raw())
	}
	keyName, loopVar := "", spec
	if idx := strings.IndexByte(spec, ','); idx >= 0 {
		keyName = spec[:idx]
		loopVar = spec[idx+1:]
	}

	iterable, err := r.evalExpr(expr, vars)
	if err != nil {
		return "", start, err
	}
	if iterable == nil {
		return "", endIdx + 1, nil
	}

	var sb strings.Builder
	rv := reflect.ValueOf(iterable)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return "", endIdx + 1, nil
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		n := rv.Len()
		for i := 0; i < n; i++ {
			child := childScope(vars)
			child[loopVar] = rv.Index(i).Interface()
			if keyName != "" {
				// "for i, v in xs" — bind index under the user-chosen name
				child[keyName] = i
			}
			child[loopVar+"_index"] = i
			child[loopVar+"_first"] = i == 0
			child[loopVar+"_last"] = i == n-1
			out, _, err := r.renderRange(tokens, start+1, endIdx, child, false)
			if err != nil {
				return "", start, err
			}
			s, _ := out.(string)
			sb.WriteString(s)
		}
	case reflect.Map:
		keys := rv.MapKeys()
		sort.Slice(keys, func(a, b int) bool {
			return fmt.Sprint(keys[a].Interface()) < fmt.Sprint(keys[b].Interface())
		})
		n := len(keys)
		for i, k := range keys {
			child := childScope(vars)
			child[loopVar] = rv.MapIndex(k).Interface()
			if keyName != "" {
				child[keyName] = k.Interface()
			} else {
				child[loopVar+"_key"] = k.Interface()
			}
			child[loopVar+"_index"] = i
			child[loopVar+"_first"] = i == 0
			child[loopVar+"_last"] = i == n-1
			out, _, err := r.renderRange(tokens, start+1, endIdx, child, false)
			if err != nil {
				return "", start, err
			}
			s, _ := out.(string)
			sb.WriteString(s)
		}
	default:
		return "", start, fmt.Errorf("for: %q is not iterable (got %s)", expr, rv.Kind())
	}

	return sb.String(), endIdx + 1, nil
}

// childScope returns a shallow copy of parent. loop bindings are added to the
// returned map; lookups in the child still resolve parent values for any keys
// not shadowed.
func childScope(parent map[string]any) map[string]any {
	child := make(map[string]any, len(parent)+4)
	for k, v := range parent {
		child[k] = v
	}
	return child
}

// evalCondition renders a single expression and returns its truthiness via
// functions.ConditionIsTrue.
func (r *StringRenderer) evalCondition(expr string, vars map[string]any) (bool, error) {
	val, err := r.evalExpr(expr, vars)
	if err != nil {
		return false, err
	}
	return functions.ConditionIsTrue(val), nil
}

// evalExpr parses a bare expression (the body of a {{ ... }} without the
// braces) and returns its rendered value.
func (r *StringRenderer) evalExpr(expr string, vars map[string]any) (any, error) {
	p := NewStringParser()
	tokens, err := p.Parse("{{ " + strings.TrimSpace(expr) + " }}")
	if err != nil {
		return nil, fmt.Errorf("failed to parse expression %q: %w", expr, err)
	}
	if len(tokens) == 0 {
		return nil, nil
	}
	// expect a single variable/filtered token (whitespace text tokens shouldn't
	// appear because we trimmed). pick the first non-text token.
	for _, t := range tokens {
		if t.Type() == VariableToken || t.Type() == FilteredVariableToken {
			return r.renderVariable(t, vars)
		}
	}
	return nil, fmt.Errorf("expression %q did not parse to a variable", expr)
}

// renderVariable renders a single variable token
func (r *StringRenderer) renderVariable(token Token, vars map[string]any) (any, error) {
	if token.Type() == TextToken {
		return token.Raw(), nil
	}

	if token.Type() != VariableToken && token.Type() != FilteredVariableToken {
		return nil, fmt.Errorf("%w: %d: %s", ErrInvalidTokenType, token.Type(), token.Raw())
	}

	if token.Type() == VariableToken {
		varValue, ok := vars[token.Name()]
		if !ok {
			return nil, fmt.Errorf("simple %w: %s", ErrVariableNotFound, token.Name())
		}

		switch val := varValue.(type) {
		case string:
			return val, nil
		case bool:
			if val {
				return "true", nil
			} else {
				return "false", nil
			}
		case int:
			return fmt.Sprintf("%d", val), nil
		}

		return varValue, nil
	}

	// handle filtered variable token
	varName, funcs := getVarAndFunctions(token)
	hasFunctionsToApply := len(funcs) > 0

	// get the value on which the function will be applied. a quoted head is a
	// literal string (e.g. {{ "path.tpl" | file }}) rather than a variable name.
	var varValue any
	var varExists bool
	if isQuotedWith(varName, `"`) {
		varValue, varExists = unquote(varName, `"`), true
	} else if isQuotedWith(varName, `'`) {
		varValue, varExists = unquote(varName, `'`), true
	} else {
		varValue, varExists = vars[varName]
	}
	if !varExists {
		// only return an error if there are no functions to apply
		// if there are functions to apply, we can assume that the variable can be optional e.g. using "default" function
		if !hasFunctionsToApply {
			return nil, fmt.Errorf("complex %w: %s", ErrVariableNotFound, varName)
		}
	}

	if !hasFunctionsToApply {
		return varValue, nil
	}

	usesDefaultFunction := hasDefaultFunction(funcs)
	for _, fn := range funcs {
		function, ok := r.funcs[fn.Name]
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrFunctionNotFound, fn.Name)
		}

		args := make([]any, len(fn.Args))
		for i, arg := range fn.Args {
			if arg.Var {
				argValue, ok := vars[arg.Value.(string)]
				if !ok {
					return nil, fmt.Errorf("function arg: %w: %s", ErrVariableNotFound, arg.Value)
				}
				args[i] = argValue
			} else {
				args[i] = arg.Value
			}
		}

		newVarValueAfterFunctions, err := function(varValue, args)
		if err != nil {
			if !usesDefaultFunction {
				return nil, fmt.Errorf("%w: %w", ErrFunctionApplyFailed, err)
			}
			if !errors.Is(err, functions.ErrAllowsDefaultFunc) {
				return nil, fmt.Errorf("%w: %w", ErrFunctionApplyFailed, err)
			}
		}

		varValue = newVarValueAfterFunctions
	}

	return varValue, nil
}

func hasDefaultFunction(funcs []Func) bool {
	for _, fn := range funcs {
		if fn.Name == "default" {
			return true
		}
	}
	return false
}

func getVarAndFunctions(token Token) (string, []Func) {
	// first, split the input based on '|' while respecting quoted sections
	split := splitRespectingQuotes(token.Raw(), "|")
	varName := strings.TrimSpace(split[0])

	funcs := make([]Func, 0)
	for _, fnWithArgs := range split[1:] {
		fnWithArgs = strings.TrimSpace(fnWithArgs)

		// find the first ':' not within quotes to split function name from args
		indexOfColon := strings.IndexFunc(fnWithArgs, func(r rune) bool {
			return r == ':' && !strings.ContainsAny(string(r), `"'`)
		})
		var fn string
		var argsStr string
		if indexOfColon != -1 {
			fn = strings.TrimSpace(fnWithArgs[:indexOfColon])
			argsStr = fnWithArgs[indexOfColon+1:]
		} else {
			fn = fnWithArgs
		}

		// split args respecting quotes
		rawArgs := splitRespectingQuotes(argsStr, ",")
		args := make([]Arg, len(rawArgs))

		for i, arg := range rawArgs {
			// unquote and unescape arguments, but only once and only if they are quoted with the same character
			// "'arg'" -> 'arg'
			// '"arg"' -> "arg"

			if isQuotedWith(arg, `"`) {
				args[i] = Arg{Value: unquote(arg, `"`)}
				continue
			}
			if isQuotedWith(arg, `'`) {
				args[i] = Arg{Value: unquote(arg, `'`)}
				continue
			}

			if num, ok := isInt(arg); ok {
				args[i] = Arg{Value: num}
				continue
			}

			if num, ok := isFloat(arg); ok {
				args[i] = Arg{Value: num}
				continue
			}

			if b, ok := isBool(arg); ok {
				args[i] = Arg{Value: b}
				continue
			}

			args[i] = Arg{Value: arg, Var: true}
		}

		funcs = append(funcs, Func{Name: fn, Args: args})
	}

	return varName, funcs
}

func isInt(s string) (int, bool) {
	i, err := strconv.Atoi(s)
	return i, err == nil
}

func isFloat(s string) (float64, bool) {
	f, err := strconv.ParseFloat(s, 64)
	return f, err == nil
}

func isBool(s string) (bool, bool) {
	if s == "true" || s == "yes" {
		return true, true
	}
	if s == "false" || s == "no" {
		return false, true
	}
	return false, false
}

// unquote removes surrounding quotes from a string and unescapes internal quotes
func unquote(str, quoteChar string) string {
	if isQuotedWith(str, quoteChar) {
		str = str[1 : len(str)-1]                                // remove surrounding quotes
		str = strings.ReplaceAll(str, "\\"+quoteChar, quoteChar) // unescape quotes
	}
	return str
}

func isQuotedWith(str string, quoteChar string) bool {
	return strings.HasPrefix(str, quoteChar) && strings.HasSuffix(str, quoteChar)
}

func splitRespectingQuotes(s, sep string) []string {
	var parts []string
	var currentPart strings.Builder
	inQuotes := false
	quoteChar := byte(0)

	for i := 0; i < len(s); i++ {
		currentChar := s[i]

		if inQuotes {
			if currentChar == quoteChar {
				// check if the quote is escaped by counting the backslashes before it
				backslashCount := 0
				for j := i - 1; j >= 0 && s[j] == '\\'; j-- {
					backslashCount++
				}
				// if the number of backslashes is even (including zero), it's not an escaped quote
				if backslashCount%2 == 0 {
					inQuotes = false
					quoteChar = 0
				}
			}
			currentPart.WriteByte(currentChar)
		} else {
			// if the current character is a quote, start a quoted section
			if currentChar == '"' || currentChar == '\'' {
				inQuotes = true
				quoteChar = currentChar
				// skip the quote character
			} else if strings.HasPrefix(s[i:], sep) {
				parts = append(parts, strings.TrimSpace(currentPart.String()))
				currentPart.Reset()
				i += len(sep) - 1 // skip the separator
				continue
			}
			currentPart.WriteByte(currentChar)
		}
	}

	if currentPart.Len() > 0 {
		parts = append(parts, strings.TrimSpace(currentPart.String()))
	}

	return parts
}
