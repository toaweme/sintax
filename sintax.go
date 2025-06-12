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

func (sm *Sintax) ExtractDependencies(vars map[string]any) ([]string, error) {
	missingInterpolatedVars, dependencyGraph, err := sm.buildDependencyGraph(vars)
	if err != nil {
		return nil, err
	}

	// spew.Dump("missingInterpolatedVars", missingInterpolatedVars)
	// spew.Dump("dependencyGraph", dependencyGraph)

	// topological sort to determine resolution order
	sortedVars, err := topologicalSort(dependencyGraph)
	if err != nil {
		return nil, err
	}

	for varName := range missingInterpolatedVars {
		sortedVars = append(sortedVars, varName)
	}

	return sortedVars, nil
}

func (sm *Sintax) ResolveVariables(vars map[string]any) (map[string]any, error) {
	missingInterpolatedVars, dependencyGraph, err := sm.buildDependencyGraph(vars)
	if err != nil {
		return nil, err
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

	resolvedVars, vars, err := sm.resolveVarsInOrder(vars, sortedVars)
	if err != nil {
		return nil, err
	}

	// ensure non-interpolated vars are copied over
	for varName, value := range vars {
		if _, exists := resolvedVars[varName]; !exists {
			resolvedVars[varName] = value
		}
	}

	return resolvedVars, nil
}

// only the modifications below, everything else remains unchanged.

func (sm *Sintax) buildDependencyGraph(vars map[string]any) (map[string]any, map[string][]string, error) {
	missingInterpolatedVars := make(map[string]any)
	dependencyGraph := make(map[string][]string)

	// build dependency graph
	for varName, value := range vars {
		switch val := value.(type) {
		case string:
			tokens, err := sm.parser.Parse(val)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse variable '%s': %w", varName, err)
			}
			for _, token := range tokens {
				if token.Type() == VariableToken || token.Type() == FilteredVariableToken {
					variable := token.Name()
					if _, inVars := vars[variable]; inVars {
						dependencyGraph[varName] = append(dependencyGraph[varName], variable)
					} else {
						missingInterpolatedVars[variable] = nil
					}

					params := token.Params()
					if len(params) > 0 {
						for _, param := range params {
							if _, inVars := vars[param]; inVars {
								dependencyGraph[variable] = append(dependencyGraph[variable], param)
							} else {
								missingInterpolatedVars[param] = nil
							}
						}
					}
				}
			}
		case map[string]any:
			// if the variable is a map, traverse it and collect dependencies
			if err := sm.parseMapDependencies(varName, val, vars, dependencyGraph, missingInterpolatedVars); err != nil {
				return nil, nil, err
			}
		}
	}

	return missingInterpolatedVars, dependencyGraph, nil
}

// parseMapDependencies traverses a map looking for string fields referencing other variables.
func (sm *Sintax) parseMapDependencies(
	parentVarName string,
	mapValue map[string]any,
	vars map[string]any,
	dependencyGraph map[string][]string,
	missing map[string]any,
) error {
	for _, v := range mapValue {
		switch subVal := v.(type) {
		case string:
			tokens, err := sm.parser.Parse(subVal)
			if err != nil {
				return fmt.Errorf("failed to parse nested variable in '%s': %w", parentVarName, err)
			}
			for _, token := range tokens {
				if token.Type() == VariableToken || token.Type() == FilteredVariableToken {
					varName := token.Name()
					if _, inVars := vars[varName]; inVars {
						dependencyGraph[parentVarName] = append(dependencyGraph[parentVarName], varName)
					} else {
						missing[varName] = nil
					}
					params := token.Params()
					if len(params) > 0 {
						for _, param := range params {
							if _, inVars := vars[param]; inVars {
								dependencyGraph[parentVarName] = append(dependencyGraph[parentVarName], param)
							} else {
								missing[param] = nil
							}
						}
					}
				}
			}
		case map[string]any:
			// recurse deeper
			if err := sm.parseMapDependencies(parentVarName, subVal, vars, dependencyGraph, missing); err != nil {
				return err
			}
		}
	}
	return nil
}

func (sm *Sintax) resolveVarsInOrder(vars map[string]any, sortedVars []string) (map[string]any, map[string]any, error) {
	resolvedVars := make(map[string]any)
	// resolve variables in order
	for _, varName := range sortedVars {
		value := vars[varName]

		switch val := value.(type) {
		case string:
			tokens, err := sm.parser.Parse(val)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse variable '%s': %w", varName, err)
			}

			// optimization: skip rendering if it's a plain text variable
			if len(tokens) == 1 && tokens[0].Type() == TextToken {
				resolvedVars[varName] = val
				continue
			}

			renderedValue, err := sm.render.Render(tokens, vars)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to render variable '%s': %w", varName, err)
			}

			resolvedVars[varName] = renderedValue
			vars[varName] = renderedValue

		case map[string]any:
			// recursively resolve maps
			resolvedMap, err := sm.resolveMap(val, vars)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to resolve map variable '%s': %w", varName, err)
			}
			resolvedVars[varName] = resolvedMap
			vars[varName] = resolvedMap

		default:
			// directly copy non-string values
			resolvedVars[varName] = val
			vars[varName] = val
		}
	}

	return resolvedVars, vars, nil
}

// resolveMap recursively resolves string fields in a map[string]any
func (sm *Sintax) resolveMap(m map[string]any, vars map[string]any) (map[string]any, error) {
	res := make(map[string]any)
	for k, v := range m {
		switch subVal := v.(type) {
		case string:
			tokens, err := sm.parser.Parse(subVal)
			if err != nil {
				return nil, fmt.Errorf("failed to parse nested field '%s': %w", k, err)
			}
			if len(tokens) == 1 && tokens[0].Type() == TextToken {
				res[k] = subVal
				continue
			}
			renderedValue, err := sm.render.Render(tokens, vars)
			if err != nil {
				return nil, fmt.Errorf("failed to render nested field '%s': %w", k, err)
			}
			res[k] = renderedValue
		case map[string]any:
			nestedMap, err := sm.resolveMap(subVal, vars)
			if err != nil {
				return nil, err
			}
			res[k] = nestedMap
		default:
			// directly copy non-string
			res[k] = subVal
		}
	}
	return res, nil
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

func (sm *Sintax) Render(input string, vars map[string]any) (any, error) {
	resolvedVars, err := sm.ResolveVariables(vars)
	if err != nil {
		return "", err
	}

	tokens, err := sm.parser.Parse(input)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrParseFailed, err)
	}

	render, err := sm.render.Render(tokens, resolvedVars)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrRenderFailed, err)
	}

	return render, nil
}
