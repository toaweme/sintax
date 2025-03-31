package collections

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

func Key(value any, params []any) (any, error) {
	val, err := key(value, params)
	if err != nil {
		// maybe this shouldn't be an error
		return nil, nil
	}
	return val, err
}

func key(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("key function requires a key parameter")
	}
	switch v := value.(type) {
	case map[string]any:
		inMap, err := findKeyInMap(params, v)
		if err != nil {
			return nil, fmt.Errorf("%w: key function: %w", functions.ErrAllowsDefaultFunc, err)
		}
		return inMap, nil
	case []any:
		index, err := functions.ParamInt(params, 0)
		if err != nil {
			return nil, fmt.Errorf("key function: index for slice must be an int")
		}
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: key function: index %d out of range", functions.ErrAllowsDefaultFunc, index)
		}
		return v[index], nil
	case nil:
		return nil, functions.ErrAllowsDefaultFunc
	default:
		return nil, fmt.Errorf("key function expected map or slice, got %T", value)
	}
}

func findKeyInMap(params []any, m map[string]any) (any, error) {
	parts, err := keyParts(params)
	if err != nil {
		return nil, err
	}

	current := m
	for i, part := range parts {
		nextVal, exists := current[part]
		if !exists {
			return nil, fmt.Errorf("key function: path segment %q not found in map", part)
		}

		if i == len(parts)-1 {
			return nextVal, nil
		}

		subMap, ok := nextVal.(map[string]any)
		if !ok {
			return nil, fmt.Errorf(
				"key function: path segment %q is not a map; cannot continue nested lookup",
				part,
			)
		}
		current = subMap
	}
	return nil, nil
}

func keyParts(params []any) ([]string, error) {
	keyPath, ok := params[0].(string)
	if !ok {
		return nil, fmt.Errorf("key function requires a string key parameter")
	}

	parts := strings.Split(keyPath, ".")
	return parts, nil
}
