package collections

import (
	"fmt"
	
	"github.com/toaweme/sintax/functions"
)

func Find(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("in requires at least one parameter")
	}
	
	keyValue, err := functions.ParamAny(params, 1)
	if err != nil {
		return nil, err
	}
	
	key, err := functions.ParamString(params, 0)
	if err != nil {
		return nil, err
	}
	
	switch slice := value.(type) {
	case []map[string]any:
		for _, m := range slice {
			if v, ok := m[key]; ok {
				if s, ok := v.(string); ok && s == keyValue {
					return m, nil
				}
			}
		}
		return nil, fmt.Errorf("key %q with value %q not found in slice", key, keyValue)
	case map[string]any:
		if v, ok := slice[key]; ok {
			if s, ok := v.(string); ok && s == keyValue {
				return slice, nil
			}
		}
		return nil, fmt.Errorf("key %q with value %q not found in map", key, keyValue)
	}
	
	return nil, fmt.Errorf("expected slice of maps, got %T", value)
}
