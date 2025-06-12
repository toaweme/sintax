package boolean

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

var Eq = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("eq function requires at least one parameter")
	}

	other := params[0]

	// Handle nil comparison
	if value == nil || other == nil {
		return value == other, nil
	}

	// Try numeric comparison
	valNum, errVal := functions.ValueNumber(value)
	otherNum, errOther := functions.ValueNumber(other)
	if errVal == nil && errOther == nil {
		return valNum == otherNum, nil
	}

	// Try string comparison
	valStr, okVal := value.(string)
	otherStr, okOther := other.(string)
	if okVal && okOther {
		return valStr == otherStr, nil
	}

	// Fallback to direct equality
	return value == other, nil
}
