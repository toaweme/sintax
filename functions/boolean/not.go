package boolean

import (
	"github.com/toaweme/sintax/functions"
)

var NotDefinition = functions.ModifierDefinition{}

var Not = func(value any, _ []any) (any, error) {
	valueIsTrue := functions.ConditionIsTrue(value)
	return !valueIsTrue, nil
}
