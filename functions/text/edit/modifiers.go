package edit

import "github.com/toaweme/sintax/functions"

// Modifiers returns the string-editing modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameShorten):        Shorten,
		string(ModifierNameConcat):         Concat,
		string(ModifierNameReplace):        Replace,
		string(ModifierNameReplacePattern): ReplacePattern,
		string(ModifierNameReverse):        Reverse,
		string(ModifierNameWrap):           Wrap,
	}
}
