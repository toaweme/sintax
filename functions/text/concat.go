package text

import (
	"fmt"
	"strings"
)

func Concat(value any, params []any) (any, error) {
	joined := make([]string, 0)

	switch v := value.(type) {
	case string:
		joined = append(joined, v)
		for _, p := range params {
			switch pv := p.(type) {
			case string:
				joined = append(joined, pv)
			default:
				return nil, fmt.Errorf("concat function expected string params, got %T", p)
			}
		}
		return strings.Join(joined, ""), nil
	}

	return nil, fmt.Errorf("concat function expected string, got %T", value)
}
