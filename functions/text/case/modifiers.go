package casing

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine.
var (
	lowerModifier      = functions.Wrap(ToLower)
	upperModifier      = functions.Wrap(ToUpper)
	slugModifier       = functions.Wrap(Slug)
	titleModifier      = functions.WrapVariadic(Title)
	titleModelModifier = functions.Wrap(ModelTitle)
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
