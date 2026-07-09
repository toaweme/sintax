package fs

import "github.com/toaweme/sintax/functions"

// Modifiers returns the filesystem modifiers keyed by their template names.
// safeDirs is the allowlist of directories the `file` modifier may read from;
// pass nil to leave file reads disabled.
func Modifiers(safeDirs []string) map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFile): File(safeDirs),
	}
}
