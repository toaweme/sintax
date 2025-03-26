package text

import (
	"fmt"
	"strings"
)

func Split(value any, params []any) (any, error) {
	switch v := value.(type) {
	case string:
		sep, err := getParamString(params, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to get separator: %w", err)
		}
		return strings.Split(v, sep), nil
	}

	return nil, fmt.Errorf("join function expected array of strings, got %T", value)
}

func getParamString(params []any, index int) (string, error) {
	if len(params) <= index {
		return "", fmt.Errorf("missing parameter")
	}

	value := params[index]

	switch v := value.(type) {
	case string:
		return v, nil
	}

	return "", fmt.Errorf("expected string, got %T", value)
}
