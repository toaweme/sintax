package query

import "github.com/toaweme/sintax/functions"

var (
	dirnameModifier  = functions.Wrap(Dirname)
	filenameModifier = functions.Wrap(Filename)
	extModifier      = functions.Wrap(FilenameExt)
	extDotModifier   = functions.Wrap(FilenameExtDot)
)

// Modifiers returns the path-reading modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameDirname):        dirnameModifier,
		string(ModifierNameFilename):       filenameModifier,
		string(ModifierNameFilenameExt):    extModifier,
		string(ModifierNameFilenameExtDot): extDotModifier,
	}
}
