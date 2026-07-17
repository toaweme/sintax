package parse

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFromYAML is the template name for the FromYAML modifier.
const ModifierNameFromYAML functions.ModifierName = "from_yaml"

// FromYAML parses a YAML document into a map, so a config file or a manifest
// becomes data that later template steps can index into.
//
// It ships as a stub that returns an error until you inject a codec. Parsing
// YAML needs a real codec, and the core stays free of third-party dependencies
// rather than forcing one on every consumer, the same way the yaml serializer
// does. Injecting one is a map entry.
//
//	func parseYAML(doc string) (map[string]any, error) {
//		var out map[string]any
//		err := yaml.Unmarshal([]byte(doc), &out)
//		return out, err
//	}
//
//	mods := parse.Modifiers()
//	mods[string(parse.ModifierNameFromYAML)] = functions.Wrap(parseYAML)
//	out, err := sintax.New(mods).Render(tpl, vars)
func FromYAML(doc string) (map[string]any, error) {
	return nil, errors.New("from_yaml function needs to be injected")
}
