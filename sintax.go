package sintax

import (
	"fmt"
	"maps"
)

// Option configures an engine. Options are applied in order, so a later option
// wins over an earlier one that sets the same thing.
type Option func(*config)

// config is the engine configuration an Option writes to.
type config struct {
	funcs    map[string]GlobalModifier
	ctxFuncs map[string]ContextualModifier
	maxDepth int
}

// newConfig resolves opts over an empty modifier set. The zero configuration
// knows no modifiers at all, so an engine only ever calls what the caller
// passed it and importing sintax links no modifier code on its own.
func newConfig(opts []Option) *config {
	cfg := &config{
		funcs:    make(map[string]GlobalModifier),
		ctxFuncs: make(map[string]ContextualModifier),
		maxDepth: defaultMaxTemplateDepth,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithModifiers adds global modifiers keyed by their template names. It merges
// rather than replaces, so several groups compose in one New call and a later
// group overrides an earlier one that registered the same name.
func WithModifiers(funcs map[string]GlobalModifier) Option {
	return func(c *config) { maps.Copy(c.funcs, funcs) }
}

// WithContextualModifiers adds contextual modifiers keyed by their template
// names, merging on the same terms as WithModifiers. These need live render
// state rather than only their piped value, so they are wired separately from
// the global set. Pass none and a template calling one fails to resolve it.
func WithContextualModifiers(funcs map[string]ContextualModifier) Option {
	return func(c *config) { maps.Copy(c.ctxFuncs, funcs) }
}

// WithMaxDepth bounds how deeply the `template` modifier may re-enter the
// engine before ErrMaxDepthExceeded, guarding against self-referential
// templates that would otherwise recurse forever. Depths below 1 are ignored,
// since an engine that cannot render once is not a useful configuration.
func WithMaxDepth(depth int) Option {
	return func(c *config) {
		if depth >= 1 {
			c.maxDepth = depth
		}
	}
}

// WithOptions bundles opts into a single Option, so a package can hand out a
// whole preconfigured engine setup as one value that callers can still layer
// their own options on top of. See defaults.All.
func WithOptions(opts ...Option) Option {
	return func(c *config) {
		for _, opt := range opts {
			opt(c)
		}
	}
}

type sintax struct {
	parser Parser
	render Renderer
}

var _ Sintax = (*sintax)(nil)

// New creates a Sintax configured by opts. It starts from nothing, so pass at
// least a modifier set: defaults.All() for the whole battery, or WithModifiers
// with the groups you want.
func New(opts ...Option) *sintax { //nolint:revive // Sintax is the public contract
	cfg := newConfig(opts)
	return &sintax{
		parser: NewStringParser(),
		render: newTokenRenderer(cfg),
	}
}

// Render parses template and renders it against vars, returning the rendered
// value. A template that is a single variable or modifier pipeline yields that
// value's own Go type, while anything with surrounding text renders to a string.
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

// Render parses and renders template against vars in one call, using an engine
// configured by opts. Prefer New when rendering more than once, so the engine
// is built once rather than per call.
func Render(template string, vars map[string]any, opts ...Option) (any, error) {
	return New(opts...).Render(template, vars)
}
