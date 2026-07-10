package serialize

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameYAML is the template name for the YAML modifier.
const ModifierNameYAML functions.ModifierName = "yaml"

// YAML serializes or parses a value as YAML. The library ships only a stub,
// because binding a YAML codec is a choice the consumer should make (which
// library, which options), so the core does not force one on every build.
// Applications that want this modifier must inject their own implementation.
// Until injected, calling it returns the error "yaml function needs to be
// injected". The example below shows the behavior of a typical injected codec,
// not the stub.
//
// value: any
// returns: string, map
//
// example: serialize a config map to YAML (requires an injected codec)
// in:  config = {"region": "eu-west-1", "debug": false}
// tpl: {{ config | yaml }}
// out: region: eu-west-1
// out: debug: false
func YAML(value any, params []any) (any, error) {
	return nil, errors.New("yaml function needs to be injected")
}
