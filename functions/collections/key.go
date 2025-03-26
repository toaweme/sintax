package collections

import (
	"fmt"
	"strconv"
	"strings"
)

func Key(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("key function requires a key parameter")
	}

	switch v := value.(type) {
	case map[string]any:
		return findKeyInMap(params, v)
	case []any:
		indexStr, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("key function: index for slice must be a string")
		}
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return nil, fmt.Errorf("key function: invalid index %q for slice: %v", indexStr, err)
		}
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("key function: index %d out of range", index)
		}
		return v[index], nil
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
