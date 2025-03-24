package functions

import (
	"encoding/json"
	"fmt"
	"strings"
)

func From(value any, params []any) (any, error) {
	if len(params) > 0 && isParam(params, 0, "json") {
		val, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("from function expected string for json, got %T", value)
		}

		dec := json.NewDecoder(strings.NewReader(val))
		dec.UseNumber()

		var raw map[string]any
		if err := dec.Decode(&raw); err != nil {
			return nil, fmt.Errorf("failed to convert JSON to map: %w", err)
		}

		return convertNumbers(raw), nil
	}

	return nil, fmt.Errorf("unsupported format in from function")
}

func convertNumbers(v any) any {
	switch vv := v.(type) {
	case map[string]any:
		for k, val := range vv {
			vv[k] = convertNumbers(val)
		}
		return vv

	case []any:
		for i, val := range vv {
			vv[i] = convertNumbers(val)
		}
		return vv

	case json.Number:
		s := vv.String()
		if strings.ContainsAny(s, ".eE") {
			f, err := vv.Float64()
			if err == nil {
				return f
			}
			return vv
		} else {
			i, err := vv.Int64()
			if err == nil {
				return i
			}
			return vv
		}

	default:
		// Non-number types remain unchanged
		return v
	}
}
