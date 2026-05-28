package text

import "github.com/toaweme/sintax/functions"

// ModifierNameSexy is the template name for the Sexy modifier.
const ModifierNameSexy functions.ModifierName = "sexy"

// Sexy returns a bear ASCII art. For emergencies only.
//
// value: any
// returns: string
//
// example: brighten a quiet template
// in:  anything = ""
// tpl: {{ anything | sexy }}
// out:  ʕ•ᴥ•ʔ
// out: /\o-o/\
// out:  | ᴥ |
// out:  \_|_/
func Sexy(value any, params []any) (any, error) {
	sexyBear := `
	 ʕ•ᴥ•ʔ
	/\o-o/\
	 | ᴥ |
	 \_|_/
	`
	return sexyBear, nil
}
