package sintax

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
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
	var out any
	for _, token := range tokens {
		switch token.Type() {
		case TextToken:
			return token.Raw(), nil
		case VariableToken, FilteredVariableToken:
			variable, err := r.renderVariable(token, vars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable token '%s': %w", token.Raw(), err)
			}
			return variable, nil
		case IfToken:
			// Implementation for conditional rendering will go here
		case ElseToken, IfEndToken:
			// Handle else and end tokens as part of conditionals
		default:
		}
	}

	return out, nil
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
			return nil, fmt.Errorf("variable '%s' not found", token.Raw())
		}

		return varValue, nil
	}

	// handle filtered variable token
	split := strings.Split(token.Raw(), "|")
	varName := strings.TrimSpace(split[0])

	funcs := make([]Func, 0)
	for _, fnWithArgs := range split[1:] {
		fnWithArgs = strings.TrimSpace(fnWithArgs)
		// summary:255,300 or summary
		splitFnWithArgs := strings.Split(fnWithArgs, ":")
		if len(splitFnWithArgs) > 1 {
			fn := splitFnWithArgs[0]
			args := strings.Split(splitFnWithArgs[1], ",")

			funcs = append(funcs, Func{Name: fn, Args: castToAny(args)})
			continue
		}

		fn := splitFnWithArgs[0]
		args := make([]any, 0)

		funcs = append(funcs, Func{Name: fn, Args: args})
	}

	var hasFunctionsToApply = len(funcs) > 0

	// get the variable value on which the function will be applied
	varValue, ok := vars[varName]
	if !ok {
		// only return an error if there are no functions to apply
		if !hasFunctionsToApply {
			return nil, fmt.Errorf("variable '%s' not found", varName)
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
			log.Err(err).Interface("args", fn.Args).Msg("")
			return nil, err
		}

		log.Trace().
			Str("function", fn.Name).
			Str("variable", varName).
			Any("value", varValue).
			Interface("args", fn.Args).
			Msg("applying function on variable value")

		// apply the function
		newVarValueAfterFunctions, err := function(varValue, fn.Args)
		if err != nil {
			err = fmt.Errorf("failed to apply function '%s': %w", fn.Name, err)
			log.Err(err).Interface("args", fn.Args).Msg("")
			return nil, err
		}

		varValue = newVarValueAfterFunctions
	}

	return varValue, nil
}

func castToAny(val []string) []any {
	anyVal := make([]any, len(val))
	for i, v := range val {
		anyVal[i] = v
	}

	return anyVal
}
