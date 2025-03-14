package sintax

import (
	"fmt"
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
	resolvedVars := make(map[string]any)
	missingInterpolatedVars := make(map[string]any)
	dependencyGraph := make(map[string][]string)

	// build dependency graph
	for varName, value := range vars {
		if strVal, ok := value.(string); ok {
			tokens, err := sm.parser.ParseVariable(strVal)
			if err != nil {
				return nil, fmt.Errorf("failed to parse variable '%s': %w", varName, err)
			}

			for _, token := range tokens {
				if token.Type() == VariableToken || token.Type() == FilteredVariableToken {
					variable := token.Name()
					if _, inVars := vars[variable]; inVars {
						dependencyGraph[varName] = append(dependencyGraph[varName], variable)
					} else {
						missingInterpolatedVars[variable] = nil
					}
				}
			}
		}
	}

	// error out if there are missing interpolated vars
	if len(missingInterpolatedVars) > 0 {
		for varName := range missingInterpolatedVars {
			return nil, fmt.Errorf("%w: %s", ErrVariableNotFound, varName)
		}
	}

	// topological sort to determine resolution order
	sortedVars, err := topologicalSort(dependencyGraph)
	if err != nil {
		return nil, err
	}

	// resolve variables in order
	for _, varName := range sortedVars {
		value := vars[varName]

		if strVal, ok := value.(string); ok {
			tokens, err := sm.parser.ParseVariable(strVal)
			if err != nil {
				return nil, fmt.Errorf("failed to parse variable '%s': %w", varName, err)
			}

			// optimization: skip rendering if it's a plain text variable
			if len(tokens) == 1 && tokens[0].Type() == TextToken {
				resolvedVars[varName] = strVal
				continue
			}

			renderedValue, err := sm.render.Render(tokens, vars)
			if err != nil {
				return nil, fmt.Errorf("failed to render variable '%s': %w", varName, err)
			}

			resolvedVars[varName] = renderedValue
			vars[varName] = renderedValue
		} else {
			// directly copy non-string values
			resolvedVars[varName] = value
			vars[varName] = value
		}
	}

	// ensure non-interpolated vars are copied over
	for varName, value := range vars {
		if _, exists := resolvedVars[varName]; !exists {
			resolvedVars[varName] = value
		}
	}

	return resolvedVars, nil
}

func (sm *Sintax) Render(input string, vars map[string]any) (string, error) {
	tokens, err := sm.parser.Parse(input)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrParseFailed, err)
	}

	render, err := sm.render.RenderString(tokens, vars)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrRenderFailed, err)
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
			return fmt.Errorf("%w: %s", ErrCircularDependency, node)
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
