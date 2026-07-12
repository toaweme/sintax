// Package render provides modifiers that render a value through the template engine.
package render

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameTemplate is the template name for the Template modifier.
const ModifierNameTemplate functions.ModifierName = "template"

// Template renders its incoming string value as a nested sintax template, so a
// string loaded from a file (or any variable) can itself contain sintax markup
// that is expanded in a second pass. It is the partial/include primitive and
// composes with the file modifier, e.g. {{ "partial.tpl" | file | template }}
// reads a file and then renders it.
//
// With no params the nested template renders against the parent's variables, so
// the partial sees everything the outer template sees (inherited scope). Passing
// a single map param switches to isolated scope, where the nested template sees
// only that map and the parent variables are hidden. Isolation is total, not a
// merge, so any variable the partial needs must be present in the passed map.
//
// The nested render uses the same modifiers as the outer template. A template
// that renders itself, directly or through a chain, is stopped at a fixed
// maximum nesting depth with ErrMaxDepthExceeded rather than recursing forever.
// A non-string value, or a non-map extra param, is an error.
//
// value: string (the template source)
// param:0: map (optional; the isolated scope, replacing the parent variables)
// returns: the rendered value
//
// example: render a string variable that itself contains markup, inheriting parent vars
// in:  tpl = "Hi {{ name }}"
// in:  name = "Bob"
// tpl: {{ tpl | template }}
// out: Hi Bob
//
// example: read a partial from disk and render it against the parent scope
// in:  who = "World"
// tpl: {{ "p.tpl" | file | template }}
// out: Hello World
//
// example: isolated scope hides parent vars and exposes only the extra map
// in:  tpl = "{{ name | default:'?' }}/{{ city }}"
// in:  name = "parent"
// in:  extra = {"city": "Vilnius"}
// tpl: {{ tpl | template:extra }}
// out: ?/Vilnius
//
// example: a self-referential template is stopped by the depth guard
// in:  self = "{{ self | template }}"
// tpl: {{ self | template }}
// out: <error: max template nesting depth exceeded>
func Template(render func(template string, vars map[string]any) (any, error), vars map[string]any, value any, params []any) (any, error) {
	src, err := functions.ValueString(value)
	if err != nil {
		return nil, fmt.Errorf("failed to read template source: %w", err)
	}

	scope := vars
	if len(params) > 0 && params[0] != nil {
		m, ok := params[0].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("failed to apply template modifier: extra vars must be a map, got %T", params[0])
		}
		scope = m
	}

	out, err := render(src, scope)
	if err != nil {
		return nil, fmt.Errorf("failed to render nested template: %w", err)
	}
	return out, nil
}
