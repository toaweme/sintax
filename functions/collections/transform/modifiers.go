package transform

import "github.com/toaweme/sintax/functions"

// Modifiers returns the collection transform modifiers keyed by their template
// names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameMap):     Map,
		string(ModifierNameSort):    Sort,
		string(ModifierNameMerge):   Merge,
		string(ModifierNameSum):     Sum,
		string(ModifierNameFlatten): Flatten,
	}
}
