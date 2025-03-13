package sintax

import (
	"fmt"
	
	"github.com/toaweme/log"
)

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
	
	log.Debug("system", "vars", systemVars)
	log.Debug("config", "vars", configVars)
	log.Debug("action", "vars", actionVars)
	log.Debug("output", "vars", outputVars)
	
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
	allVars := mergeMaps(systemVars, vars)
	resolvedVars := make(map[string]any)
	dependencyGraph := make(map[string][]string)
	
	// Step 1: Build dependency graph only for vars defined in `vars`
	for varName, value := range vars {
		if strVal, ok := value.(string); ok {
			tokens, err := sm.parser.ParseVariable(strVal)
			if err != nil {
				return nil, fmt.Errorf("failed to parse variable '%s': %w", varName, err)
			}
			for _, token := range tokens {
				if token.Type() == VariableToken || token.Type() == FilteredVariableToken {
					dependency := token.Raw()
					if _, inVars := vars[dependency]; inVars {
						dependencyGraph[varName] = append(dependencyGraph[varName], dependency)
					}
				}
			}
		}
	}
	
	// topological sort to determine resolution order
	sortedVars, err := topologicalSort(dependencyGraph)
	if err != nil {
		return nil, err
	}
	
	// resolve variables based on type (string needing interpolation, others copied directly)
	for _, varName := range sortedVars {
		value := vars[varName]
		switch val := value.(type) {
		case string:
			tokens, err := sm.parser.ParseVariable(val)
			if err != nil {
				return nil, fmt.Errorf("failed to parse variable '%s': %w", varName, err)
			}
			renderedValue, err := sm.render.Render(tokens, allVars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable '%s': %w", varName, err)
			}
			resolvedVars[varName] = renderedValue
			allVars[varName] = renderedValue
		default:
			// directly copy values that don't require interpolation
			resolvedVars[varName] = val
			allVars[varName] = val
		}
	}
	
	// Step 4: Include variables that were not part of the graph (no dependencies)
	for varName, value := range vars {
		if _, alreadyResolved := resolvedVars[varName]; !alreadyResolved {
			resolvedVars[varName] = value
			allVars[varName] = value
		}
	}
	
	finalVars := mergeMaps(systemVars, resolvedVars)
	return finalVars, nil
}

func topologicalSort(graph map[string][]string) ([]string, error) {
	var sorted []string
	visited := make(map[string]bool)
	tempMarked := make(map[string]bool)
	
	var visit func(node string) error
	visit = func(node string) error {
		if tempMarked[node] {
			return fmt.Errorf("%w: at variable '%s'", ErrCircularDependency, node)
		}
		if !visited[node] {
			tempMarked[node] = true
			for _, dep := range graph[node] {
				if err := visit(dep); err != nil {
					return err
				}
			}
			tempMarked[node] = false
			visited[node] = true
			sorted = append(sorted, node)
		}
		return nil
	}
	
	for node := range graph {
		if !visited[node] {
			if err := visit(node); err != nil {
				return nil, err
			}
		}
	}
	
	return sorted, nil
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
