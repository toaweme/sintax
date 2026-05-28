package text

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameLines is the template name for the Lines modifier.
const ModifierNameLines functions.ModifierName = "lines"

// Lines splits a string or byte slice into an array of lines.
//
// value: string, bytes
// returns: array
//
// example: split a multi-line note
// in:  note = "Buy milk\nWalk the dog\nPay rent"
// tpl: {{ note | lines }}
// out: ["Buy milk", "Walk the dog", "Pay rent"]
//
// example: take only the first line of a message
// in:  message = "Subject: Welcome\nThanks for signing up!"
// tpl: {{ message | lines | first }}
// out: Subject: Welcome
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
