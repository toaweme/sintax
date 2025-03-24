package functions

import "strings"

var Wrap = func(s any, params []any) (any, error) {

	switch v := s.(type) {
	case string:
		return strings.TrimSpace(v), nil
	case []byte:
		return []byte(strings.TrimSpace(string(v))), nil
	}
	return s, nil
}
