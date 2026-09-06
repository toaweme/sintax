package math

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameDivide is the template name for the Divide modifier.
const ModifierNameDivide functions.ModifierName = "divide"

// Divide returns the value split by the given number, which is an average from
// a total and a count, or a share from an amount and a number of people.
//
// Dividing by zero is an error rather than an infinity. A count that came out
// zero is an empty collection, and a template that printed "+Inf" over it would
// be reporting the absence of data as a result.
func Divide(value any, divisor any) (float64, error) {
	left, right, err := operands(ModifierNameDivide, value, divisor)
	if err != nil {
		return 0, err
	}
	if right == 0 {
		return 0, fmt.Errorf("failed to divide %v by zero: %w", left, functions.ErrInvalidParamValue)
	}
	return left / right, nil
}
