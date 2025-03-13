package sintax

import (
	"errors"
	"fmt"
	
	"github.com/toaweme/log"
)

type Syntax interface {
	// ResolveVariables resolves all variables in the given system, config, and action variables.
	// systemVars are variables that are always available to the pipeline e.g. env, now, etc.
	// configVars are variables defined in the pipeline config
	// actionVars are variables defined in the action
	// previousOutputVars are variables that were output from the previous action
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

// ResolveVariables resolves all variables in the given system, config, and action variables.
func (sm *Sintax) ResolveVariables(systemVars map[string]any, configVars map[string]any, actionVars map[string]any, outputVars map[string]any) (map[string]any, error) {
	resolvedConfigVars, err := sm.resolveVariables(systemVars, configVars)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve config variables: %w", err)
	}
	
	all := mergeMaps(systemVars, resolvedConfigVars, outputVars)
	
	resolvedActionVars, err := sm.resolveVariables(all, actionVars)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve action variables: %w", err)
	}
	
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
	
	log.Debug("combined", "vars", resolvedVars)
	
	// var tokenizedVars = make(map[string][]Token)
	
	// helper function to resolve a variable value, supporting 1 level of recursion
	var resolveVarValue = func(key string, value any) (any, error) {
		switch val := value.(type) {
		case string:
			log.Debug("resolving", "string", val)
			tokens, err := sm.parser.ParseVariable(val)
			if err != nil {
				return nil, fmt.Errorf("failed to parse template for variable '%s': %w", key, err)
			}
			log.Debug("tokens", "tokens", tokens)
			variableValue, err := sm.render.Render(tokens, resolvedVars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable '%s': %w", key, err)
			}
			return variableValue, nil
		case map[string]any:
			log.Debug("resolving", "map", val)
			resolvedMap, err := sm.resolveVariables(systemVars, val) // Recursive call for nested maps
			if err != nil {
				return nil, fmt.Errorf("failed to resolve map variable '%s': %w", key, err)
			}
			return resolvedMap, nil
		default:
			return val, nil
		}
	}
	
	// second pass: resolve any variable references
	for varName, varValue := range vars {
		resolvedValue, err := resolveVarValue(varName, varValue)
		if err != nil {
			return nil, err
		}
		resolvedVars[varName] = resolvedValue
	}
	
	return resolvedVars, nil
}

// topoSort performs a topological sort on the dependency graph.
// It returns an error if a circular dependency is detected.
func topoSort(graph map[string][]string) ([]string, error) {
	var order []string
	temp := make(map[string]bool)
	perm := make(map[string]bool)
	
	var visit func(string) error
	visit = func(node string) error {
		if perm[node] {
			return nil
		}
		if temp[node] {
			return errors.New("circular dependency detected")
		}
		temp[node] = true
		for _, dep := range graph[node] {
			// Only visit dependencies that are part of the graph.
			if _, exists := graph[dep]; exists {
				if err := visit(dep); err != nil {
					return err
				}
			}
		}
		temp[node] = false
		perm[node] = true
		order = append(order, node)
		return nil
	}
	
	for node := range graph {
		if !perm[node] {
			if err := visit(node); err != nil {
				return nil, err
			}
		}
	}
	return order, nil
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
