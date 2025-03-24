package functions

import (
	"fmt"
	"strings"
)

func Join(value any, params []any) (any, error) {
	switch v := value.(type) {
	case []string:
		return strings.Join(v, "\n"), nil
	}

	return nil, fmt.Errorf("join function expected array of strings, got %T", value)
}
