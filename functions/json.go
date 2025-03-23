package functions

import (
	"encoding/json"
	"fmt"
)

func JSON(val any, params []any) (any, error) {
	if len(params) > 0 && isParam(params, 0, "pretty") {
		jsonBytes, err := json.MarshalIndent(val, "", "  ")
		if err != nil {
			return "", fmt.Errorf("failed To apply json filter: %w", err)
		}

		return string(jsonBytes), nil
	}
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("failed To apply json filter: %w", err)
	}

	return string(jsonBytes), nil
}
