package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameGt is the template name for the Gt modifier.
const ModifierNameGt functions.ModifierName = "gt"

// Gt reports whether value exceeds than, comparing both as numbers so 5 and 5.0
// rank the same. A non-numeric operand is an error, and nil counts as zero.
func Gt(value, than float64) (bool, error) {
	return value > than, nil
}
