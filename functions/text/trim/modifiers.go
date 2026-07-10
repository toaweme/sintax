package trim

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Each is an Overload over the string and []byte value shapes and
// the present/absent cutset param; the cutset clause is listed before the
// no-param clause so the no-param whitespace trim acts as the fallback.
var (
	trimModifier = functions.Overload(
		functions.WrapOne(TrimStringSet),
		functions.Wrap(TrimString),
		functions.WrapOne(TrimBytesSet),
		functions.Wrap(TrimBytes),
	)
	trimPrefixModifier = functions.Overload(
		functions.WrapOne(TrimPrefixStringArg),
		functions.Wrap(TrimPrefixString),
		functions.WrapOne(TrimPrefixBytesArg),
		functions.Wrap(TrimPrefixBytes),
	)
	trimSuffixModifier = functions.Overload(
		functions.WrapOne(TrimSuffixStringArg),
		functions.Wrap(TrimSuffixString),
		functions.WrapOne(TrimSuffixBytesArg),
		functions.Wrap(TrimSuffixBytes),
	)
)

// Modifiers returns the trimming modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameTrim):       trimModifier,
		string(ModifierNameTrimPrefix): trimPrefixModifier,
		string(ModifierNameTrimSuffix): trimSuffixModifier,
	}
}
