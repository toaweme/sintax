package escape

import "github.com/toaweme/sintax/functions"

var (
	escapeHTMLModifier = functions.Wrap(HTML)
	escapeURLModifier  = functions.Wrap(URL)
	escapeJSModifier   = functions.Wrap(JS)

	escapeBackslashModifier   = functions.WrapVariadic(Backslash)
	unescapeBackslashModifier = functions.WrapVariadic(Unbackslash)
)

// Modifiers returns the context-escaping modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameHTML): escapeHTMLModifier,
		string(ModifierNameURL):  escapeURLModifier,
		string(ModifierNameJS):   escapeJSModifier,

		string(ModifierNameBackslash):   escapeBackslashModifier,
		string(ModifierNameUnbackslash): unescapeBackslashModifier,
	}
}
