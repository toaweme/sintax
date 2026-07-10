package splitjoin

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. lines is an Overload over a nil passthrough plus the string and
// []byte shapes; join is an Overload over the with-separator and
// default-separator arities.
var (
	linesModifier = functions.Overload(
		linesNil,
		functions.Wrap(LinesString),
		functions.Wrap(LinesBytes),
	)
	joinModifier = functions.Overload(
		functions.WrapOne(JoinAny),
		functions.Wrap(JoinDefault),
	)
	splitModifier = functions.WrapOne(Split)
)

// Modifiers returns the split and join modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameLines): linesModifier,
		string(ModifierNameJoin):  joinModifier,
		string(ModifierNameSplit): splitModifier,
	}
}
