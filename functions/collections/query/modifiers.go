package query

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine. has stays a plain
// modifier: what it matches depends on both the value shape (slice, slice of
// maps, or map) and the param count, so it does not fit a typed clause.
var (
	filterModifier = functions.WrapTwo(Filter)
	isModifier     = functions.WrapVariadic(Is)
)

// Modifiers returns the collection query modifiers keyed by their template
// names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFilter): filterModifier,
		string(ModifierNameHas):    Has,
		string(ModifierNameIs):     isModifier,
	}
}
