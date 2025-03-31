package boolean

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

var GtDefinition = functions.ModifierDefinition{}

var Gt = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("gt function requires at least one parameter")
	}

	val, err := functions.ValueNumber(value)
	if err != nil {
		return nil, fmt.Errorf("gt function expected a number, got %T", value)
	}

	than, err := functions.ValueNumber(params[0])
	if err != nil {
		return nil, fmt.Errorf("gt function expected a number, got %T", params[0])
	}

	return val > than, nil
}
