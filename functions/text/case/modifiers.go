package casing

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine. Each is wrapped in
// AsText: these are text-first, so a scalar value (a number or bool) is handed
// in as its string form rather than rejected.
var (
	lowerModifier      = functions.AsText(functions.Wrap(ToLower))
	upperModifier      = functions.AsText(functions.Wrap(ToUpper))
	slugModifier       = functions.AsText(functions.Wrap(Slug))
	titleModifier      = functions.AsText(functions.WrapVariadic(Title))
	titleModelModifier = functions.AsText(functions.Wrap(ModelTitle))
)

// Modifiers returns the case-transformation and slug modifiers keyed by their
// template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameToLower):    lowerModifier,
		string(ModifierNameToUpper):    upperModifier,
		string(ModifierNameSlug):       slugModifier,
		string(ModifierNameTitle):      titleModifier,
		string(ModifierNameModelTitle): titleModelModifier,
	}
}
