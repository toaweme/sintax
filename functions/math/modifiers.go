package math

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine.
var (
	addModifier      = functions.WrapOne(Add)
	subtractModifier = functions.WrapOne(Subtract)
	multiplyModifier = functions.WrapOne(Multiply)
	divideModifier   = functions.WrapOne(Divide)
)

// Modifiers returns the arithmetic modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameAdd):      addModifier,
		string(ModifierNameSubtract): subtractModifier,
		string(ModifierNameMultiply): multiplyModifier,
		string(ModifierNameDivide):   divideModifier,
	}
}
