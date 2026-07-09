// Package render provides modifiers that render a value through the template engine.
package render

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameTemplate is the template name for the Template modifier.
const ModifierNameTemplate functions.ModifierName = "template"

// Template renders its incoming string value as a nested sintax template, so a
// value loaded from a file (or any string variable) can itself contain
// {{ ... }} markup. It is the partial/include primitive and composes with the
// file modifier: {{ "partial.tpl" | file | template }}.
//
// With no params it renders against the parent's variables (inherited scope).
// With a single map param it renders against only that map (isolated scope, the
// parent scope is disabled).
//
// render is supplied by the engine and re-enters it with the same modifiers and
// recursion guard, so this modifier holds no renderer state of its own.
//
// value: string (the template source)
// returns: the rendered value
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
