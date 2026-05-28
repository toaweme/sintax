package sintax

import "fmt"

type sintax struct {
	parser Parser
	render Renderer
}

var _ Sintax = (*sintax)(nil)

func New(funcs map[string]GlobalModifier) *sintax {
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

func Render(template string, vars map[string]any, funcs map[string]GlobalModifier) (any, error) {
	return New(funcs).Render(template, vars)
}
