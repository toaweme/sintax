package serialize

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameYAML is the template name for the YAML modifier.
const ModifierNameYAML functions.ModifierName = "yaml"

// YAML serializes a value as YAML, so a map assembled in a template can be
// written out as a config file or a manifest.
//
// It ships as a stub that returns an error until you inject a codec. Which YAML
// library to bind, and with which options, is a choice that belongs to the
// application rather than to this package, and the core stays free of
// third-party dependencies by not making it for you. Injecting one is a map
// entry.
//
//	func marshalYAML(v any) (string, error) {
//		out, err := yaml.Marshal(v)
//		return string(out), err
//	}
//
//	mods := serialize.Modifiers()
//	mods[string(serialize.ModifierNameYAML)] = functions.Wrap(marshalYAML)
//	out, err := sintax.New(mods).Render(tpl, vars)
//
// It takes any value rather than only a string, since anything renderable is
// serializable.
func YAML(value any) (string, error) {
	return "", errors.New("yaml function needs to be injected")
}
