package functions

import "fmt"

func First(value any, params []any) (any, error) {
	switch v := value.(type) {
	case string:
		if len(v) > 0 {
			return string(v[0]), nil
		}
	case []byte:
		if len(v) > 0 {
			return v[0], nil
		}
	case []any:
		if len(v) > 0 {
			return v[0], nil
		}
	}

	return nil, fmt.Errorf("first function expected a non-empty string or bytes or slice, got %T", value)
}
