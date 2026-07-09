package query

import "github.com/toaweme/sintax/functions"

// Modifiers returns the path-reading modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameDirname):        Dirname,
		string(ModifierNameFilename):       Filename,
		string(ModifierNameFilenameExt):    FilenameExt,
		string(ModifierNameFilenameExtDot): FilenameExtDot,
	}
}
