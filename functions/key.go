package functions

import (
	"fmt"
	"strings"
)

func Key(value any, params []any) (any, error) {
	// We only handle maps in this function
	m, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("key function expected map, got %T", value)
	}

	// We need at least one parameter for the key name
	if len(params) == 0 {
		return nil, fmt.Errorf("key function requires a key parameter")
	}

	// spew.Dump("Key", val, params)

	// The key param must be a string (potentially "nested.key.name")
	keyPath, ok := params[0].(string)
	if !ok {
		return nil, fmt.Errorf("key function requires a string key parameter")
	}

	parts := strings.Split(keyPath, ".")

	current := m
	for i, part := range parts {
		// Look up the part in the current map
		nextVal, exists := current[part]
		if !exists {
			return nil, fmt.Errorf("key function: path segment %q not found in map", part)
		}

		// If this is the last part in the path, we're done
		if i == len(parts)-1 {
			// spew.Dump("KeyRES.val", val)
			// spew.Dump("KeyRES.params", params)
			// spew.Dump("KeyRES.keyPath", keyPath)
			// spew.Dump("KeyRES.parts", parts)
			// spew.Dump("KeyRES.final", nextVal)
			return nextVal, nil
		}

		// Otherwise, we need to keep traversing; ensure the nextVal is another map
		subMap, ok := nextVal.(map[string]any)
		if !ok {
			return nil, fmt.Errorf(
				"key function: path segment %q is not a map; cannot continue nested lookup",
				part,
			)
		}
		current = subMap
	}

	// In theory, we never get here if parts has at least one element, but just in case:
	return nil, fmt.Errorf("key function: unexpected end of path resolution")
}
