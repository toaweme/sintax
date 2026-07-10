package edit

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. shorten is an Overload so its length accepts either an int
// (shorten:30) or a numeric string (shorten:'30').
var (
	shortenModifier = functions.Overload(
		functions.WrapOne(Shorten),
		functions.WrapOne(ShortenParse),
	)
	concatModifier         = functions.WrapVariadic(Concat)
	replaceModifier        = functions.WrapTwo(Replace)
	replacePatternModifier = functions.WrapTwo(ReplacePattern)
	reverseModifier        = functions.Wrap(Reverse)
	wrapModifier           = functions.WrapOne(Wrap)
)

// Modifiers returns the string-editing modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameShorten):        shortenModifier,
		string(ModifierNameConcat):         concatModifier,
		string(ModifierNameReplace):        replaceModifier,
		string(ModifierNameReplacePattern): replacePatternModifier,
		string(ModifierNameReverse):        reverseModifier,
		string(ModifierNameWrap):           wrapModifier,
	}
}
