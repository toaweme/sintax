package math

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// operands reads the piped value and the argument as numbers, keeping the two
// rejections apart. A value that is not a number says the data is not what the
// template expected, while an argument that is not a number says the template
// itself is wrong, and only the second is worth failing on no matter what the
// pipeline holds downstream.
func operands(name functions.ModifierName, value any, param any) (float64, float64, error) {
	left, err := functions.ParseNumber(value)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read the value %s was given: %w", name, err)
	}
	right, err := functions.ParseNumber(param)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read %v as the %s argument: %w", param, name, functions.ErrInvalidParamType)
	}
	return left, right, nil
}
