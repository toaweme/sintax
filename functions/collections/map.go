package collections

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

func Map(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("in requires at least one parameter")
	}

	key, err := functions.ParamString(params, 0)
	if err != nil {
		return nil, err
	}

	switch slice := value.(type) {
	case []map[string]any:
		n := make(map[string]map[string]any)
		for _, m := range slice {
			if v, ok := m[key]; ok {
				if s, ok := v.(string); ok {
					n[s] = m
				} else {
					return nil, fmt.Errorf("expected string value in map at key %q, got %T", key, v)
				}
			}
		}
		return n, nil
	}

	return nil, fmt.Errorf("expected slice of maps, got %T", value)
}
