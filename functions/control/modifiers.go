package control

import "github.com/toaweme/sintax/functions"

var (
	defaultModifier = functions.WrapOne(Default)
	whenModifier    = functions.WrapTwo(When)
)

// Modifiers returns the value-resolution control modifiers keyed by their
// template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameDefault): defaultModifier,
		string(ModifierNameWhen):    whenModifier,
	}
}
