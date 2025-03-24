package functions

import (
	"encoding/json"
	"fmt"
)

func From(value any, params []any) (any, error) {
	if len(params) > 0 && isParam(params, 0, "json") {
		v := make(map[string]any)
		if _, ok := value.(string); !ok {
			return nil, fmt.Errorf("from function expected string, got %T", v)
		}

		err := json.Unmarshal([]byte(value.(string)), &v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert JSON to map: %w", err)
		}
		return v, nil
	}

	return nil, fmt.Errorf("unsupported format in from function")
}
