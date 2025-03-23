package functions

import (
	"encoding/json"
	"fmt"
)

func From(val any, params []any) (any, error) {
	if len(params) > 0 && isParam(params, 0, "json") {
		value := make(map[string]any)
		if _, ok := val.(string); !ok {
			return nil, fmt.Errorf("from function expected string, got %T", val)
		}

		err := json.Unmarshal([]byte(val.(string)), &value)
		if err != nil {
			return nil, fmt.Errorf("failed to convert JSON to map: %w", err)
		}
		return value, nil
	}

	return nil, fmt.Errorf("unsupported format in from function")
}
