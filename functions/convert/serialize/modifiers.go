package serialize

import "github.com/toaweme/sintax/functions"

var (
	jsonModifier = functions.Overload(
		functions.Wrap(JSON),
		functions.WrapOne(JSONMode),
	)
	yamlModifier     = functions.Wrap(YAML)
	markdownModifier = functions.Wrap(Markdown)
)

// Modifiers returns the serialization modifiers keyed by their template names.
// json is an Overload so its optional mode param accepts either no argument
// (compact) or the literal 'pretty' (indented). YAML and Markdown are deliberately
// stubbed to return an error so the package stays free of third-party
// dependencies; a consumer injects a real codec by overriding the entry.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameJSON):     jsonModifier,
		string(ModifierNameYAML):     yamlModifier,
		string(ModifierNameMarkdown): markdownModifier,
	}
}
