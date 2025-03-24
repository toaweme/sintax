package functions

import (
	"encoding/json"
	"fmt"
)

func JSON(value any, params []any) (any, error) {
	if len(params) > 0 && isParam(params, 0, "pretty") {
		jsonBytes, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			return "", fmt.Errorf("failed to apply json filter: %w", err)
		}

		return string(jsonBytes), nil
	}
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("failed to apply json filter: %w", err)
	}

	return string(jsonBytes), nil
}
