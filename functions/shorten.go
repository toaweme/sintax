package functions

import (
	"fmt"
	"strconv"
)

var Shorten = func(s any, args []any) (any, error) {
	str, ok := s.(string)
	if !ok {
		return "", fmt.Errorf("shorten requires a text argument")
	}

	if len(args) != 1 {
		return "", fmt.Errorf("shorten requires 1 numeric argument")
	}

	length, err := strconv.Atoi(fmt.Sprint(args[0]))
	if err != nil {
		return "", fmt.Errorf("shorten requires a numeric argument")
	}

	if len(str) > length {
		return str[:length], nil
	}

	return str, nil
}
