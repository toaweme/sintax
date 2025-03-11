package sintax

import (
	"fmt"
	"strings"
)

type Renderer interface {
	RenderString(tokens []Token, vars map[string]any) (string, error)
	Render(tokens []Token, vars map[string]any) (any, error)
}

type Func struct {
	Name string
	Args []any
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

// RenderString renders the template based on the parsed tokens.
func (r *StringRenderer) RenderString(tokens []Token, vars map[string]any) (string, error) {
	var sb strings.Builder

	for _, token := range tokens {
		switch token.Type() {
		case TextToken:
			sb.WriteString(token.Raw())
		case VariableToken:
			variable, err := r.renderVariable(token, vars)
			if err != nil {
				return "", fmt.Errorf("failed to render variable: %w", err)
			}
			val := fmt.Sprintf("%v", variable)
			sb.WriteString(val)
		case IfToken:
			// Implementation for conditional rendering will go here
		case ElseToken, IfEndToken:
			// Handle else and end tokens as part of conditionals
		case FilteredVariableToken:
			// Filters handling will be implemented here
		default:
		}
	}

	return sb.String(), nil
}

func (r *StringRenderer) Render(tokens []Token, vars map[string]any) (any, error) {
	// log.Trace().Interface("tokens", tokens).Interface("vars", vars).Msg("rendering tokens")
	var str strings.Builder
	for _, token := range tokens {
		switch token.Type() {
		case TextToken:
			str.WriteString(token.Raw())
		case VariableToken, FilteredVariableToken:
			variable, err := r.renderVariable(token, vars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable token '%s': %w", token.Raw(), err)
			}
			if val, ok := variable.(string); ok {
				str.WriteString(val)
				continue
			}
			return variable, nil
		case IfToken:
			// Implementation for conditional rendering will go here
		case ElseToken, IfEndToken:
			// Handle else and end tokens as part of conditionals
		default:
		}
	}

	return str.String(), nil
}

// RenderVariable renders a single variable token.
func (r *StringRenderer) renderVariable(token Token, vars map[string]any) (any, error) {
	if token.Type() == TextToken {
		return token.Raw(), nil
	}

	if token.Type() != VariableToken && token.Type() != FilteredVariableToken {
		return nil, fmt.Errorf("invalid token type: %d: %s", token.Type(), token.Raw())
	}

	if token.Type() == VariableToken {
		varValue, ok := vars[token.Raw()]
		if !ok {
			return nil, fmt.Errorf("plain variable '%s' not found", token.Raw())
		}

		return varValue, nil
	}

	// handle filtered variable token
	varName, funcs := r.getVarAndFunctions(token)
	hasFunctionsToApply := len(funcs) > 0

	// get the variable value on which the function will be applied
	varValue, varExists := vars[varName]
	if !varExists {
		// only return an error if there are no functions to apply
		// if there are functions to apply, we can assume that the variable can be optional e.g. using "default" function
		if !hasFunctionsToApply {
			return nil, fmt.Errorf("variable '%s' does not exist", varName)
		}
	}

	if !hasFunctionsToApply {
		return varValue, nil
	}

	for _, fn := range funcs {
		// get the function
		function, ok := r.funcs[fn.Name]
		if !ok {
			err := fmt.Errorf("function '%s' not found", fn.Name)
			// log.Err(err).Interface("args", fn.Args).Msg("")
			return nil, err
		}

		// log.Trace().
		// 	Str("function", fn.Name).
		// 	Str("variable", varName).
		// 	Any("value", varValue).
		// 	Interface("args", fn.Args).
		// 	Msg("applying function on variable value")

		// apply the function
		newVarValueAfterFunctions, err := function(varValue, fn.Args)
		if err != nil {
			err = fmt.Errorf("failed to apply function '%s': %w", fn.Name, err)
			// log.Err(err).Interface("args", fn.Args).Msg("")
			return nil, err
		}

		varValue = newVarValueAfterFunctions
	}

	return varValue, nil
}

func (r *StringRenderer) getVarAndFunctions(token Token) (string, []Func) {
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
		args := splitRespectingQuotes(argsStr, ",")
		// spew.Dump(args)
		for i, arg := range args {
			// unquote and unescape arguments, but only once and only if they are quoted with the same character
			// "'arg'" -> 'arg'
			// '"arg"' -> "arg"

			if isQuotedWith(arg, `"`) {
				// log.Trace().Str("arg", arg).Msg("unquoting double")
				args[i] = unquote(arg, `"`)
				continue
			}
			if isQuotedWith(arg, `'`) {
				// log.Trace().Str("arg", arg).Msg("unquoting single")
				args[i] = unquote(arg, `'`)
				continue
			}
		}

		funcs = append(funcs, Func{Name: fn, Args: castToAny(args)})
	}

	return varName, funcs
}

func castToAny(args []string) []any {
	if args == nil || len(args) == 0 {
		return make([]any, 0)
	}
	var result []any
	for _, arg := range args {
		result = append(result, arg)
	}
	return result
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
