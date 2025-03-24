package functions

import "fmt"

func Default(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("default requires at least one parameter")
	}
	if value == nil {
		return params[0], nil
	}
	if value == "" {
		return params[0], nil
	}

	return value, nil
}
