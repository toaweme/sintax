package functions

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func YAML(value any, params []any) (any, error) {
	yamlBytes, err := yaml.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("failed to apply yaml filter: %w", err)
	}

	return string(yamlBytes), nil
}
