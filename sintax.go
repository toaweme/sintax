package sintax

import (
	"fmt"
	"time"
)

type Syntax interface {
	ResolveVariables(systemVars map[string]any, configVars map[string]any, actionVars map[string]any, previousOutputVars map[string]any) (map[string]any, error)
	Render(input string, vars map[string]any) (string, error)
}

type Sintax struct {
	parser Parser
	render Renderer
}

func New(funcs map[string]GlobalModifier) *Sintax {
	tplParser := NewStringParser()
	tplRender := NewStringRenderer(funcs)

	return NewWith(tplParser, tplRender)
}

var _ Syntax = (*Sintax)(nil)

func NewWith(parser Parser, render Renderer) *Sintax {
	return &Sintax{parser: parser, render: render}
}

func (sm *Sintax) Render(input string, vars map[string]any) (string, error) {
	tokens, err := sm.parser.Parse(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	render, err := sm.render.RenderString(tokens, vars)
	if err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return render, err
}

func (sm *Sintax) ResolveVariables(systemVars map[string]any, configVars map[string]any, actionVars map[string]any, outputVars map[string]any) (map[string]any, error) {
	systemVars["now"] = time.Now()
	resolvedConfigVars, err := sm.resolveVariables(systemVars, configVars)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve config variables: %w", err)
	}
	
	all := mergeMaps(systemVars, resolvedConfigVars, outputVars)

	// log.Trace().Interface("all", all).Msg("all vars")

	resolvedActionVars, err := sm.resolveVariables(all, actionVars)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve action variables: %w", err)
	}

	// log.Trace().Interface("vars", resolvedActionVars).Msg("action vars")
	variables := make(map[string]any)
	for k, v := range resolvedConfigVars {
		variables[k] = v
	}
	for k, v := range resolvedActionVars {
		variables[k] = v
	}

	return variables, nil
}

func (sm *Sintax) resolveVariables(systemVars map[string]any, vars map[string]any) (map[string]any, error) {
	resolvedVars := make(map[string]any)
	for k, v := range systemVars {
		resolvedVars[k] = v
	}
	// initially resolve all variables to their raw or directly resolvable values
	for k, v := range vars {
		resolvedVars[k] = v
	}

	// helper function to resolve a variable value, supporting 1 level of recursion
	var resolveVarValue func(key string, value any) (any, error)
	resolveVarValue = func(key string, value any) (any, error) {
		switch val := value.(type) {
		case string:
			tokens, err := sm.parser.ParseVariable(val)
			if err != nil {
				return nil, fmt.Errorf("failed to parse template for variable '%s': %w", key, err)
			}
			variableValue, err := sm.render.Render(tokens, resolvedVars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable '%s': %w", key, err)
			}
			return variableValue, nil
		case map[string]any:
			resolvedMap, err := sm.resolveVariables(systemVars, val) // Recursive call for nested maps
			if err != nil {
				return nil, fmt.Errorf("failed to resolve map variable '%s': %w", key, err)
			}
			return resolvedMap, nil
		default:
			return val, nil
		}
	}

	// Second pass: resolve any variable references
	for varName, varValue := range vars {
		resolvedValue, err := resolveVarValue(varName, varValue)
		if err != nil {
			return nil, err
		}
		resolvedVars[varName] = resolvedValue
	}

	return resolvedVars, nil
}

func (sm *Sintax) resolveVariables2(systemVars map[string]any, vars map[string]any) (map[string]any, error) {
	resolvedVars := make(map[string]any)
	for k, v := range systemVars {
		resolvedVars[k] = v
	}
	for k, v := range vars {
		// we add raw values to be able to reference them in same group vars
		// TODO: maybe a bad idea(?)
		// but otherwise since maps have no order, we can't reference a value from a previous key consistently
		resolvedVars[k] = v
	}

	for varName, varValue := range vars {
		switch val := varValue.(type) {
		case string:
			tokens, err := sm.parser.ParseVariable(val)
			if err != nil {
				return nil, fmt.Errorf("failed to parse template for variable '%s': %w", varName, err)
			}

			variableValue, err := sm.render.Render(tokens, resolvedVars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable '%s': %w", varName, err)
			}

			resolvedVars[varName] = variableValue
			// log.Trace().
			// 	Any("input", varValue).
			// 	Str("input-type", reflect.TypeOf(varValue).String()).
			// 	Str("var", varName).
			// 	Any("value", variableValue).
			// 	Msg("resolved template variable")

		case map[string]any:
			resolvedMap, err := sm.resolveVariables(systemVars, val)
			if err != nil {
				return nil, fmt.Errorf("failed to resolve map variable '%s': %w", varName, err)
			}
			resolvedVars[varName] = resolvedMap
		default:
			resolvedVars[varName] = val
		}
	}

	return resolvedVars, nil
}

func mergeMaps(maps ...map[string]any) map[string]any {
	merged := make(map[string]any)
	for _, m := range maps {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}
