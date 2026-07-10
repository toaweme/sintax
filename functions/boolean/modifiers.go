package boolean

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine.
var (
	notModifier = functions.Wrap(Not)
	gtModifier  = functions.WrapOne(Gt)
	gteModifier = functions.WrapOne(Gte)
	eqModifier  = functions.Overload(
		eqNilGuard,
		functions.WrapOne(EqNumber),
		functions.WrapOne(EqString),
		functions.WrapOne(EqAny),
	)
)

// Modifiers returns the boolean comparison modifiers keyed by their template
// names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameNot): notModifier,
		string(ModifierNameGt):  gtModifier,
		string(ModifierNameGte): gteModifier,
		string(ModifierNameEq):  eqModifier,
	}
}
