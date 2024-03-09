package sintax

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type GlobalModifier func(any, []any) (any, error)

var BuiltinFunctions = map[string]GlobalModifier{
	"default": toDefault,
	"json":    toJSON,
	"yaml":    toYAML,
	"sexy":    toSexy,
}

func toSexy(val any, params []any) (any, error) {
	sexyBear := `
	 ʕ•ᴥ•ʔ
	/\o-o/\
	 | ᴥ |
	 \_|_/
	`
	return sexyBear, nil
}

func isParam(params []any, index int, name string) bool {
	if len(params) <= index {
		return false
	}

	return params[index] == name
}

func toDefault(varValue any, params []any) (any, error) {
	if varValue == nil {
		return params[0], nil
	}

	return varValue, nil
}

func toJSON(val any, params []any) (any, error) {
	if len(params) > 0 && isParam(params, 0, "pretty") {
		jsonBytes, err := json.MarshalIndent(val, "", "  ")
		if err != nil {
			return "", fmt.Errorf("failed to apply json filter: %w", err)
		}

		return string(jsonBytes), nil
	}
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("failed to apply json filter: %w", err)
	}

	return string(jsonBytes), nil
}

func toYAML(val any, params []any) (any, error) {
	yamlBytes, err := yaml.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("failed to apply yaml filter: %w", err)
	}

	return string(yamlBytes), nil
}
