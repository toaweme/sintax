package functions

import (
	"fmt"
)

var Length = func(val any, _ []any) (any, error) {
	switch v := val.(type) {
	case string:
		return fmt.Sprintf("%d", len(v)), nil
	case []byte:
		return fmt.Sprintf("%d", len(v)), nil
	}

	return nil, fmt.Errorf("length function expected string or bytes, got %T", val)
}
