// Package math provides the four arithmetic template modifiers, add, subtract,
// multiply and divide. Each takes the number to work with as its argument, so
// a pipeline reads as the operation applied to the value it carries.
//
// A numeric string is accepted, since a number read out of a store or a form
// often arrives as text, and anything that is not a number is refused rather
// than counted as zero. nil counts as zero, the same reading sum and decimal
// give it.
package math

import (
	"github.com/toaweme/sintax/functions"
)

// ModifierNameAdd is the template name for the Add modifier.
const ModifierNameAdd functions.ModifierName = "add"

// Add returns the value plus the given number, so `{{ subtotal | add:shipping }}`
// is the two of them together. A negative argument subtracts, which makes
// subtract a convenience rather than the only way round.
func Add(value any, addend any) (float64, error) {
	left, right, err := operands(ModifierNameAdd, value, addend)
	if err != nil {
		return 0, err
	}
	return left + right, nil
}
