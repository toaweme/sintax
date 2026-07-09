package trim

import "github.com/toaweme/sintax/functions"

// Modifiers returns the trimming modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameTrim):       Trim,
		string(ModifierNameTrimPrefix): TrimPrefix,
		string(ModifierNameTrimSuffix): TrimSuffix,
	}
}
