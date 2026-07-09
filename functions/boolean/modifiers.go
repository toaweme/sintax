package boolean

import "github.com/toaweme/sintax/functions"

// Modifiers returns the boolean comparison modifiers keyed by their template
// names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameNot): Not,
		string(ModifierNameGt):  Gt,
		string(ModifierNameGte): Gte,
		string(ModifierNameEq):  Eq,
	}
}
