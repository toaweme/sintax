package parse

import "github.com/toaweme/sintax/functions"

// Modifiers returns the parsing modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFrom): From,
	}
}
