// Package parse provides modifiers that parse a serialized value into data.
package parse

import "github.com/toaweme/sintax/functions"

var (
	fromJSONModifier = functions.Wrap(FromJSON)
	fromCSVModifier  = functions.Wrap(FromCSV)
	fromYAMLModifier = functions.Wrap(FromYAML)
)

// Modifiers returns the parsing modifiers keyed by their template names. There
// is one modifier per format rather than a single modifier taking the format as
// a parameter, because a format parameter has to be dispatched on at run time.
// That costs every caller a typed signature, since the return shape would depend
// on a string argument, and turns an unknown format into a render-time error. A
// name per format declares its return type up front, so from_json is a map and
// from_csv is a list of rows.
//
// from_yaml is registered alongside them but ships as a stub, since parsing YAML
// needs a codec this dependency-free package will not pick. It is named and
// documented here rather than left out so the seam is visible.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFromJSON): fromJSONModifier,
		string(ModifierNameFromCSV):  fromCSVModifier,
		string(ModifierNameFromYAML): fromYAMLModifier,
	}
}
