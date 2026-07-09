package splitjoin

import "github.com/toaweme/sintax/functions"

// Modifiers returns the split and join modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameLines): Lines,
		string(ModifierNameJoin):  Join,
		string(ModifierNameSplit): Split,
	}
}
