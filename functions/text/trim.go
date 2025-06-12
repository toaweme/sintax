package text

import (
	"fmt"
	"strings"
)

var Trim = func(value any, params []any) (any, error) {
	switch v := value.(type) {
	case string:
		if len(params) == 0 {
			return strings.TrimSpace(v), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}
		return strings.Trim(v, chars), nil
	case []byte:
		if len(params) == 0 {
			return []byte(strings.TrimSpace(string(v))), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}
		return []byte(strings.Trim(string(v), chars)), nil
	default:
		return nil, fmt.Errorf("expected string or []byte, got %T", value)
	}
}

var TrimPrefix = func(value any, params []any) (any, error) {
	cutset := "\n \t"

	switch v := value.(type) {
	case string:
		if len(params) == 0 {
			return strings.TrimLeft(v, cutset), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}
		return strings.TrimPrefix(v, chars), nil
	case []byte:
		if len(params) == 0 {
			return []byte(strings.TrimLeft(string(v), cutset)), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}
		return []byte(strings.TrimPrefix(string(v), chars)), nil
	default:
		return nil, fmt.Errorf("expected string or []byte, got %T", value)
	}
}
