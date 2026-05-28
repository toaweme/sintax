package convert

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameYAML is the template name for the YAML modifier.
const ModifierNameYAML functions.ModifierName = "yaml"

// YAML serializes or parses a value as YAML.
// Must be injected by the consumer; returns an error by default.
//
// value: any
// returns: string, map
//
// example: serialize a config map to YAML
// in:  config = {"region": "eu-west-1", "debug": false}
// tpl: {{ config | yaml }}
// out: region: eu-west-1
// out: debug: false
func YAML(value any, params []any) (any, error) {
	return nil, fmt.Errorf("yaml function needs to be injected")
}
