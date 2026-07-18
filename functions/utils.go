package functions

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ValueString asserts the value is a string, returning ErrInvalidValueType otherwise.
func ValueString(v any) (string, error) {
	if vv, ok := v.(string); ok {
		return vv, nil
	}
	return "", fmt.Errorf("%w: expected string, got %T", ErrInvalidValueType, v)
}

// ValueSlice asserts the value is a slice or array, returning ErrInvalidValueType otherwise.
func ValueSlice(v any) ([]any, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		result := make([]any, rv.Len())
		for i := range rv.Len() {
			result[i] = rv.Index(i).Interface()
		}
		return result, nil
	}
	return nil, fmt.Errorf("%w: expected slice or array, got %T", ErrInvalidValueType, v)
}

// ValueNumber coerces the value to a float64 across the numeric kinds, returning
// ErrInvalidValueType if it is not numeric (nil is treated as zero).
func ValueNumber(v any) (float64, error) {
	switch vv := v.(type) {
	case float64:
		return vv, nil
	case float32:
		return float64(vv), nil
	case int:
		return float64(vv), nil
	case int8:
		return float64(vv), nil
	case int16:
		return float64(vv), nil
	case int32:
		return float64(vv), nil
	case int64:
		return float64(vv), nil
	case uint:
		return float64(vv), nil
	case uint8:
		return float64(vv), nil
	case uint16:
		return float64(vv), nil
	case uint32:
		return float64(vv), nil
	case uint64:
		return float64(vv), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("%w: expected number, got %T", ErrInvalidValueType, v)
	}
}

// ValueInt coerces the value to an int across the numeric kinds, reporting false for a
// non-numeric value or for a float carrying a fractional part, since silently
// truncating would hide a caller mistake. It backs Wrap's int slot, where a
// string must not satisfy an int param, so unlike the lenient index coercion in
// collections/access it neither parses strings nor truncates.
func ValueInt(v any) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, true
	case int8:
		return int(n), true
	case int16:
		return int(n), true
	case int32:
		return int(n), true
	case int64:
		return int(n), true
	case uint:
		return int(n), true
	case uint8:
		return int(n), true
	case uint16:
		return int(n), true
	case uint32:
		return int(n), true
	case uint64:
		return int(n), true
	case float32:
		if float32(int(n)) == n {
			return int(n), true
		}
	case float64:
		if float64(int(n)) == n {
			return int(n), true
		}
	}
	return 0, false
}

// ParseNumber coerces the value to a float64 across the numeric kinds, additionally
// parsing a numeric string and treating nil as zero. It is the lenient companion
// to ValueNumber (which rejects strings), for modifiers such as decimal and sum
// whose value may legitimately arrive as text.
func ParseNumber(v any) (float64, error) {
	if s, ok := v.(string); ok {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse %q as a number: %w", s, ErrInvalidValueType)
		}
		return f, nil
	}
	return ValueNumber(v)
}

// ParamStringList asserts every element of params is a string, returning
// ErrInvalidParamType at the first mismatch.
func ParamStringList(params []any) ([]string, error) {
	if len(params) == 0 {
		return []string{}, nil
	}

	var result []string
	for i, param := range params {
		v, ok := param.(string)
		if !ok {
			return nil, fmt.Errorf("%w: expected string at index %d, got %T", ErrInvalidParamType, i, param)
		}
		result = append(result, v)
	}

	return result, nil
}

// ParamString returns params[index] as a string, returning ErrMissingParam if
// index is out of range or ErrInvalidParamType if the value isn't a string.
func ParamString(params []any, index int) (string, error) {
	if len(params) <= index {
		return "", fmt.Errorf("%w: missing parameter at index %d", ErrMissingParam, index)
	}

	v, ok := params[index].(string)
	if !ok {
		return "", fmt.Errorf("%w: expected string at index %d, got %T", ErrInvalidParamType, index, params[index])
	}

	return v, nil
}

// ParamAny returns params[index], or ErrMissingParam if index is out of range.
func ParamAny(params []any, index int) (any, error) {
	if len(params) <= index {
		return nil, fmt.Errorf("%w: missing parameter at index %d", ErrMissingParam, index)
	}

	return params[index], nil
}

// ParamInt returns params[index] as an int, returning ErrMissingParam if
// index is out of range or ErrInvalidParamType if the value isn't an int.
func ParamInt(params []any, index int) (int, error) {
	if len(params) <= index {
		return 0, fmt.Errorf("%w: missing parameter at index %d", ErrMissingParam, index)
	}

	v, ok := params[index].(int)
	if !ok {
		return 0, fmt.Errorf("%w: expected int at index %d, got %T", ErrInvalidParamType, index, params[index])
	}

	return v, nil
}

// IsParam reports whether params[index] equals name.
func IsParam(params []any, index int, name string) bool {
	if len(params) <= index {
		return false
	}

	return params[index] == name
}

// ConditionIsTrue reports whether condition should be treated as truthy in a
// template `if`: booleans and numbers by their value, strings by "true"/
// "false" or non-empty content, and collections by non-zero length.
func ConditionIsTrue(condition any) bool {
	if condition == nil {
		return false
	}
	switch v := condition.(type) {
	case bool:
		return v
	case string:
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
	case []any:
		return len(v) > 0
	case error:
		return v != nil
	case map[string]any:
		return len(v) > 0
	default:
		return false
	}
}

// ConvertNumbersJSON recursively replaces json.Number values in the input (as
// produced by a decoder configured with UseNumber) with float64 or int64, so template
// modifiers see native numeric types instead of json.Number.
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
		}
		i, err := vv.Int64()
		if err == nil {
			return i
		}
		return vv

	default:
		return v
	}
}
