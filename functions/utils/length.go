package utils

import (
	"fmt"
)

var Length = func(value any, _ []any) (any, error) {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("%d", len(v)), nil
	case []byte:
		return fmt.Sprintf("%d", len(v)), nil
	case []any:
		return fmt.Sprintf("%d", len(v)), nil
	}

	return nil, fmt.Errorf("length function expected string or bytes, got %T", value)
}
