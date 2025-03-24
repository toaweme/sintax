package functions

import "strings"

var Trim = func(value any, _ []any) (any, error) {
	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v), nil
	case []byte:
		return []byte(strings.TrimSpace(string(v))), nil
	}
	return value, nil
}
