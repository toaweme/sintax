package functions

import (
	"bytes"
	"fmt"
	"strings"
)

func Lines(value any, params []any) (any, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case string:
		return strings.Split(v, "\n"), nil
	case []byte:
		return bytes.Split(v, []byte("\n")), nil
	}

	return nil, fmt.Errorf("lines function expected string or bytes, got %T", value)
}
