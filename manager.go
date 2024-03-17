package sintax

import (
	"fmt"
	"reflect"
	"time"

	"github.com/rs/zerolog/log"
)

type Manager interface {
	ResolveVariables(systemVars map[string]any, configVars map[string]any, actionVars map[string]any, previousOutputVars map[string]any) (map[string]any, error)
	Render(input string, vars map[string]any) (string, error)
}

type StringManager struct {
	parser Parser
	render Renderer
}

func NewManager(funcs map[string]GlobalModifier) *StringManager {
	allFuncs := make(map[string]GlobalModifier)
	for k, v := range BuiltinFunctions {
		allFuncs[k] = v
	}
	for k, v := range funcs {
		allFuncs[k] = v
	}
	tplParser := NewStringParser()
	tplRender := NewStringRenderer(allFuncs)

	return NewStringManager(tplParser, tplRender)
}

func NewStringManager(parser Parser, render Renderer) *StringManager {
	return &StringManager{parser: parser, render: render}
}

var _ Manager = (*StringManager)(nil)

func (sm *StringManager) Render(input string, vars map[string]any) (string, error) {
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

func (sm *StringManager) ResolveVariables(systemVars map[string]any, configVars map[string]any, actionVars map[string]any, outputVars map[string]any) (map[string]any, error) {
	systemVars["now"] = time.Now()
	resolvedConfigVars, err := sm.resolveVariables(systemVars, configVars)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve config variables: %w", err)
	}

	// log.Trace().Interface("res", resolvedConfigVars).Msg("resolved config vars")

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

func (sm *StringManager) resolveVariables(systemVars map[string]any, vars map[string]any) (map[string]any, error) {
	resolvedVars := make(map[string]any)
	for k, v := range systemVars {
		resolvedVars[k] = v
	}

	for varName, varValue := range vars {
		switch val := varValue.(type) {
		case string:
			tokens, err := sm.parser.ParseVariable(val)
			if err != nil {
				return nil, fmt.Errorf("failed to parse template for variable '%s': %w", varName, err)
			}
			if varName == "prompt" {
				log.Trace().Interface("tokens", tokens).Msg("prompt tokens")
			}

			variableValue, err := sm.render.Render(tokens, resolvedVars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable '%s': %w", varName, err)
			}

			resolvedVars[varName] = variableValue
			log.Trace().
				Any("input", varValue).
				Str("input-type", reflect.TypeOf(varValue).String()).
				Str("var", varName).
				Any("value", variableValue).
				Msg("resolved template variable")
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
