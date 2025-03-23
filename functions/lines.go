package functions

import (
	"bytes"
	"fmt"
	"strings"
)

func Lines(val any, params []any) (any, error) {
	if val == nil {
		return nil, nil
	}

	switch v := val.(type) {
	case string:
		return strings.Split(v, "\n"), nil
	case []byte:
		return bytes.Split(v, []byte("\n")), nil
	}

	return nil, fmt.Errorf("lines function expected string or bytes, got %T", val)
}
