package escape

import "github.com/toaweme/sintax/functions"

// Modifiers returns the context-escaping modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameHTML): HTML,
		string(ModifierNameURL):  URL,
		string(ModifierNameJS):   JS,
	}
}
