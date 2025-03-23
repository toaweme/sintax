package functions

import (
	"fmt"
	"strings"
)

func Join(val any, params []any) (any, error) {
	switch v := val.(type) {
	case []string:
		return strings.Join(v, "\n"), nil
	}

	return nil, fmt.Errorf("join function expected array of strings, got %T", val)
}
