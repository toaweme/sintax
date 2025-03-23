package functions

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func YAML(val any, params []any) (any, error) {
	yamlBytes, err := yaml.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("failed To apply yaml filter: %w", err)
	}

	return string(yamlBytes), nil
}
