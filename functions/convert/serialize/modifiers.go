package serialize

import "github.com/toaweme/sintax/functions"

// Modifiers returns the serialization modifiers keyed by their template names.
// YAML and Markdown are deliberately stubbed to return an error so the package
// stays free of third-party dependencies.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameJSON):     JSON,
		string(ModifierNameYAML):     YAML,
		string(ModifierNameMarkdown): Markdown,
	}
}
