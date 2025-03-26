package convert

import (
	"encoding/json"
	"fmt"

	"github.com/toaweme/sintax/functions"
)

func JSON(value any, params []any) (any, error) {
	if len(params) > 0 && functions.IsParam(params, 0, "pretty") {
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
