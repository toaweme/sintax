package edit

import "github.com/toaweme/sintax/functions"

var (
	extPrependModifier = functions.WrapOne(FilenamePrependExt)
	extTrimModifier    = functions.Wrap(FilenameTrimExt)
)

// Modifiers returns the path extension-rewriting modifiers keyed by their
// template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFilenamePrependExt): extPrependModifier,
		string(ModifierNameFilenameTrimExt):    extTrimModifier,
	}
}
