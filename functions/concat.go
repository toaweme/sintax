package functions

import (
	"fmt"
	"strings"
)

func Concat(value any, params []any) (any, error) {
	joined := []string{}
	switch v := value.(type) {
	case string:
		joined = append(joined, v)
		for _, p := range params {
			switch p := p.(type) {
			case string:
				joined = append(joined, p)
			default:
				return nil, fmt.Errorf("concat function expected string params, got %T", p)
			}
		}
		return strings.Join(joined, ""), nil
	}

	return nil, fmt.Errorf("concat function expected string, got %T", value)
}
