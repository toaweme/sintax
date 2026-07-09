package edit

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameShorten is the template name for the Shorten modifier.
const ModifierNameShorten functions.ModifierName = "shorten"

// Shorten truncates a string to the given maximum character length.
//
// value: string
// param:0: int
// returns: string
//
// example: clip a description for a card preview
// in:  description = "Hand-picked single-origin coffee, slow-roasted in small batches."
// tpl: {{ description | shorten:30 }}
// out: Hand-picked single-origin cof
//
// example: limit a name to a column width
// in:  name = "Alexandra Christine Whitehead"
// tpl: {{ name | shorten:12 }}
// out: Alexandra Ch
var Shorten = func(value any, args []any) (any, error) {
	str, ok := value.(string)
	if !ok {
		return "", errors.New("shorten requires a text argument")
	}

	if len(args) != 1 {
		return "", errors.New("shorten requires 1 numeric argument")
	}

	length, err := strconv.Atoi(fmt.Sprint(args[0]))
	if err != nil {
		return "", errors.New("shorten requires a numeric argument")
	}

	if len(str) > length {
		return str[:length], nil
	}

	return str, nil
}
