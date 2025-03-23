package functions

import "strings"

var Trim = func(s any, _ []any) (any, error) {
	switch v := s.(type) {
	case string:
		return strings.TrimSpace(v), nil
	case []byte:
		return []byte(strings.TrimSpace(string(v))), nil
	}
	return s, nil
}
