package sintax

import (
	"fmt"
	
	"github.com/toaweme/log"
)

type Sintax struct {
	parser Parser
	render Renderer
}

var _ Syntax = (*Sintax)(nil)

func New(funcs map[string]GlobalModifier) *Sintax {
	tplParser := NewStringParser()
	tplRender := NewStringRenderer(funcs)
	
	return NewWith(tplParser, tplRender)
}

func NewWith(parser Parser, render Renderer) *Sintax {
	return &Sintax{parser: parser, render: render}
}

func (sm *Sintax) ResolveVariables(vars map[string]any) (map[string]any, error) {
	// holds the final resolved variables
	resolvedVars := make(map[string]any)
	
	// dependency graph variables
	dependencyGraph := make(map[string][]string)
	
	for varName, value := range vars {
		log.Debug("var", "name", varName, "value", value)
		if strVal, ok := value.(string); ok {
			tokens, err := sm.parser.ParseVariable(strVal)
			if err != nil {
				return nil, fmt.Errorf("failed to parse variable '%s': %w", varName, err)
			}
			// log.Debug("tokens", "tokens", tokens)
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
			renderedValue, err := sm.render.Render(tokens, vars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable '%s': %w", varName, err)
			}
			resolvedVars[varName] = renderedValue
			vars[varName] = renderedValue
		default:
			// directly copy values that don't require interpolation
			resolvedVars[varName] = val
			vars[varName] = val
		}
	}
	
	// include variables that were not part of the graph (no dependencies or depending on system vars)
	for varName, value := range vars {
		if _, alreadyResolved := resolvedVars[varName]; alreadyResolved {
			continue
		}
		
		switch val := value.(type) {
		case string:
			tokens, err := sm.parser.ParseVariable(val)
			if err != nil {
				return nil, fmt.Errorf("failed to parse variable '%s': %w", varName, err)
			}
			
			// optimisation: skip rendering if the variable is a plain text
			if len(tokens) == 1 && tokens[0].Type() == TextToken {
				vars[varName] = val
				resolvedVars[varName] = val
				continue
			}
			
			renderedValue, err := sm.render.Render(tokens, vars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable '%s': %w", varName, err)
			}
			
			resolvedVars[varName] = renderedValue
			vars[varName] = renderedValue
		default:
			resolvedVars[varName] = val
			vars[varName] = val
		}
	}
	
	return resolvedVars, nil
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
	
	return render, nil
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
