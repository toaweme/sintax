package render

import "github.com/toaweme/sintax/functions"

// ContextualModifiers returns the render-state-aware modifiers keyed by their
// template names. These need the live renderer rather than only their piped
// value, so they are wired separately from the global modifiers.
func ContextualModifiers() map[string]functions.ContextualModifier {
	return map[string]functions.ContextualModifier{
		string(ModifierNameTemplate): Template,
	}
}
