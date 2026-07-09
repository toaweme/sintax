package sintax

import (
	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/functions/render"
)

// GlobalModifier is a stateless modifier that transforms a piped value given
// its call params, independent of any render context. It aliases the type in
// functions so callers can keep referring to sintax.GlobalModifier while the
// concrete modifiers live in leaf packages the engine does not import.
type GlobalModifier = functions.GlobalModifier

// ContextualModifier is a modifier that needs live render state - the current
// variables and a re-entrant renderer - rather than only its piped value.
type ContextualModifier = functions.ContextualModifier

// builtinContextualModifiers returns the contextual modifiers wired into every
// renderer. Unlike global modifiers these need live render state, so the engine
// wires them itself. The render package is tiny (no reflect, no encoding), so
// keeping it in the engine's import graph costs nothing.
func builtinContextualModifiers() map[string]ContextualModifier {
	return render.ContextualModifiers()
}
