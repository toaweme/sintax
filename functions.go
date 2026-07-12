package sintax

import (
	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/functions/render"
)

// GlobalModifier is a stateless modifier that transforms a piped value using its
// call params, with no access to render state.
type GlobalModifier = functions.GlobalModifier

// ContextualModifier is a modifier that needs live render state, the current
// variables and a re-entrant renderer, rather than only its piped value.
type ContextualModifier = functions.ContextualModifier

// builtinContextualModifiers returns the contextual modifiers wired into every
// renderer. Unlike global modifiers these need live render state, so the engine
// wires them itself. The render package is tiny (no reflect, no encoding), so
// keeping it in the engine's import graph costs nothing.
func builtinContextualModifiers() map[string]ContextualModifier {
	return render.ContextualModifiers()
}
