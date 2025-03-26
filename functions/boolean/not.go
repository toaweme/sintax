package boolean

import (
	"github.com/toaweme/sintax/functions"
)

var NotDefinition = functions.ModifierDefinition{}

var Not = func(value any, _ []any) (any, error) {
	return functions.ConditionIsTrue(value) == false, nil
}
