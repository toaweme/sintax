package transform

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine. sort and sum are
// Overloads: sort over a nil passthrough plus the default and directioned
// arities, sum over the field and whole-slice arities, and clean over a nil
// clause that misses plus the typed clause.
var (
	mapModifier   = functions.WrapOne(Map)
	mergeModifier = functions.WrapOne(Merge)
	sortModifier  = functions.Overload(
		sortNil,
		functions.WrapOne(SortDir),
		functions.Wrap(SortAsc),
	)
	sumModifier = functions.Overload(
		functions.WrapOne(SumField),
		functions.Wrap(SumElements),
	)
	flattenModifier = functions.Wrap(Flatten)
	cleanModifier   = functions.Overload(
		cleanNil,
		functions.Wrap(Clean),
	)
)

// Modifiers returns the collection transform modifiers keyed by their template
// names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameMap):     mapModifier,
		string(ModifierNameMerge):   mergeModifier,
		string(ModifierNameSort):    sortModifier,
		string(ModifierNameSum):     sumModifier,
		string(ModifierNameFlatten): flattenModifier,
		string(ModifierNameClean):   cleanModifier,
	}
}
