package query

import "github.com/toaweme/sintax/functions"

// Modifiers returns the collection query modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFilter): Filter,
		string(ModifierNameHas):    Has,
		string(ModifierNameIs):     Is,
	}
}
