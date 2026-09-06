package math

import "github.com/toaweme/sintax/functions"

// ModifierNameSubtract is the template name for the Subtract modifier.
const ModifierNameSubtract functions.ModifierName = "subtract"

// Subtract returns the value less the given number, the way a remaining
// quantity is worked out from a total and what has gone, as in
// `{{ seats | subtract:booked }}`. It reads left to right like the sentence it
// stands for, which `add` with a negated argument cannot do when the amount is
// a field.
func Subtract(value any, amount any) (float64, error) {
	left, right, err := operands(ModifierNameSubtract, value, amount)
	if err != nil {
		return 0, err
	}
	return left - right, nil
}
