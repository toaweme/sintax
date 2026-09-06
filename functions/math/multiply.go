package math

import "github.com/toaweme/sintax/functions"

// ModifierNameMultiply is the template name for the Multiply modifier.
const ModifierNameMultiply functions.ModifierName = "multiply"

// Multiply returns the value scaled by the given number, which is a line total
// from a price and a quantity, a tax-inclusive amount from a rate, or a count
// converted into a period.
func Multiply(value any, factor any) (float64, error) {
	left, right, err := operands(ModifierNameMultiply, value, factor)
	if err != nil {
		return 0, err
	}
	return left * right, nil
}
