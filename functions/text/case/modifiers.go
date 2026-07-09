package casing

import "github.com/toaweme/sintax/functions"

// Modifiers returns the case-transformation and slug modifiers keyed by their
// template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameToLower):    ToLower,
		string(ModifierNameToUpper):    ToUpper,
		string(ModifierNameSlug):       Slug,
		string(ModifierNameTitle):      Title,
		string(ModifierNameModelTitle): ModelTitle,
	}
}
