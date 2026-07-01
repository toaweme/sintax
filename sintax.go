package sintax

import "fmt"

type sintax struct {
	parser Parser
	render Renderer
}

var _ Sintax = (*sintax)(nil)

// New creates a Sintax using funcs as the set of available global modifiers.
func New(funcs map[string]GlobalModifier) *sintax { //nolint:revive // Sintax is the public contract; the concrete type is never named by callers
	return &sintax{
		parser: NewStringParser(),
		render: NewStringRenderer(funcs),
	}
}

func (s *sintax) Render(template string, vars map[string]any) (any, error) {
	tokens, err := s.parser.Parse(template)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	result, err := s.render.Render(tokens, vars)
	if err != nil {
		return nil, fmt.Errorf("failed to render template: %w", err)
	}

	return result, nil
}

// Render parses and renders template against vars in one call, using funcs as
// the set of available modifiers.
func Render(template string, vars map[string]any, funcs map[string]GlobalModifier) (any, error) {
	return New(funcs).Render(template, vars)
}
