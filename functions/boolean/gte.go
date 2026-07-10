package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameGte is the template name for the Gte modifier.
const ModifierNameGte functions.ModifierName = "gte"

// Gte reports whether value is greater than or equal to than. WrapOne coerces
// both operands to float64 across the int and float kinds before the call, with
// nil coercing to zero, so the body is a plain numeric comparison.
func Gte(value, than float64) (bool, error) {
	return value >= than, nil
}
