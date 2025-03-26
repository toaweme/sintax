package collections

import (
	"fmt"
)

var Wrap = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("wrap requires at least one parameter")
	}

	key := ""
	switch v := params[0].(type) {
	case string:
		key = v
	}
	return map[string]any{key: value}, nil
}
