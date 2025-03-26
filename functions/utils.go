package functions

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ParamString(params []any, index int) (string, error) {
	if len(params) <= index {
		return "", fmt.Errorf("missing parameter at index %d", index)
	}

	v, ok := params[index].(string)
	if !ok {
		return "", fmt.Errorf("expected string at index %d, got %T", index, params[index])
	}

	return v, nil
}

func ParamAny(params []any, index int) (any, error) {
	if len(params) <= index {
		return nil, fmt.Errorf("missing parameter at index %d", index)
	}

	return params[index], nil
}

func ParamInt(params []any, index int) (int, error) {
	if len(params) <= index {
		return 0, fmt.Errorf("missing parameter at index %d", index)
	}

	v, ok := params[index].(int)
	if !ok {
		return 0, fmt.Errorf("expected int at index %d, got %T", index, params[index])
	}

	return v, nil
}

func IsParam(params []any, index int, name string) bool {
	if len(params) <= index {
		return false
	}

	return params[index] == name
}

func ConditionIsTrue(condition any) bool {
	if condition == nil {
		return false
	}
	switch v := condition.(type) {
	case bool:
		return v == true
	case string:
		// condition can be rendered before being passed
		// in cases where variable is boolean, we render true/false
		if v == "false" {
			return false
		}
		if v == "true" {
			return true
		}
		return len(v) > 0
	case int:
		return v > 0
	case int8:
		return v > 0
	case int16:
		return v > 0
	case int32:
		return v > 0
	case int64:
		return v > 0
	case uint:
		return v > 0
	case uint8:
		return v > 0
	case uint16:
		return v > 0
	case uint32:
		return v > 0
	case uint64:
		return v > 0
	case float32:
		return v > 0
	case float64:
		return v > 0
	default:
		return false
	}
}

func ConvertNumbersJSON(v any) any {
	switch vv := v.(type) {
	case map[string]any:
		for k, val := range vv {
			vv[k] = ConvertNumbersJSON(val)
		}
		return vv

	case []any:
		for i, val := range vv {
			vv[i] = ConvertNumbersJSON(val)
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
