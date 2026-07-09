package edit

import "github.com/toaweme/sintax/functions"

// Modifiers returns the path extension-rewriting modifiers keyed by their
// template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFilenamePrependExt): FilenamePrependExt,
		string(ModifierNameFilenameTrimExt):    FilenameTrimExt,
	}
}
