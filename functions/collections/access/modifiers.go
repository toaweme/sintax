package access

import "github.com/toaweme/sintax/functions"

// Modifiers returns the element and field access modifiers keyed by their
// template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFirst): First,
		string(ModifierNameLast):  Last,
		string(ModifierNameKey):   Key,
		string(ModifierNamePluck): Pluck,
		string(ModifierNameFind):  Find,
	}
}
