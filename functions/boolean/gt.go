package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameGt is the template name for the Gt modifier.
const ModifierNameGt functions.ModifierName = "gt"

// Gt reports whether value is greater than than. WrapOne coerces both operands
// to float64 across the int and float kinds before the call, with nil coercing
// to zero, so the body is a plain numeric comparison and a non-numeric operand
// is rejected by the wrapper as ErrInvalidValueType / ErrInvalidParamType.
func Gt(value, than float64) (bool, error) {
	return value > than, nil
}
