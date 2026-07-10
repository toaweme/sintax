package serialize

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameYAML is the template name for the YAML modifier.
const ModifierNameYAML functions.ModifierName = "yaml"

// YAML serializes a value as YAML. The library ships only a stub, because
// binding a YAML codec is a choice the consumer should make (which library,
// which options), so the core does not force one on every build. Applications
// that want this modifier inject their own implementation by overriding the
// modifier entry. Until injected, calling it returns an error.
func YAML(value any) (string, error) {
	return "", errors.New("yaml function needs to be injected")
}
