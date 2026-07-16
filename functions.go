package sintax

import "github.com/toaweme/sintax/functions"

// GlobalModifier is a stateless modifier that transforms a piped value using its
// call params, with no access to render state.
type GlobalModifier = functions.GlobalModifier

// ContextualModifier is a modifier that needs live render state, the current
// variables and a re-entrant renderer, rather than only its piped value.
type ContextualModifier = functions.ContextualModifier
