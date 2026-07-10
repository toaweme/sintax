package escape

import "github.com/toaweme/sintax/functions"

var (
	escapeHtmlModifier = functions.Wrap(HTML)
	escapeUrlModifier  = functions.Wrap(URL)
	escapeJsModifier   = functions.Wrap(JS)
)

// Modifiers returns the context-escaping modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameHTML): escapeHtmlModifier,
		string(ModifierNameURL):  escapeUrlModifier,
		string(ModifierNameJS):   escapeJsModifier,
	}
}
