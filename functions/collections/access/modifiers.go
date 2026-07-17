package access

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine. first, last, and find
// are Overloads over their value shapes. key stays a plain modifier: it is
// forgiving by design (every lookup failure renders as nil rather than an error)
// and dispatches on both the value shape and the param type, so it does not fit
// a typed clause.
var (
	firstModifier = functions.Overload(
		functions.Wrap(FirstString),
		functions.Wrap(FirstBytes),
		functions.Wrap(FirstSlice),
	)
	lastModifier = functions.Overload(
		functions.Wrap(LastString),
		functions.Wrap(LastBytes),
		functions.Wrap(LastSlice),
	)
	pluckModifier = functions.WrapOne(Pluck)
	findModifier  = functions.Overload(
		functions.WrapTwo(FindSlice),
		functions.WrapTwo(FindMap),
	)
)

// Modifiers returns the element and field access modifiers keyed by their
// template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFirst): firstModifier,
		string(ModifierNameLast):  lastModifier,
		string(ModifierNameKey):   Key,
		string(ModifierNamePluck): pluckModifier,
		string(ModifierNameFind):  findModifier,
	}
}
