package boolean

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameEq is the template name for the Eq modifier.
const ModifierNameEq functions.ModifierName = "eq"

// Eq returns true if the value equals the given parameter.
// Numbers compare numerically across int/float types, strings compare verbatim,
// and other types fall back to direct equality.
//
// value: any
// param:0: any
// returns: bool
//
// example: compare a status string
// in:  status = "active"
// tpl: {{ status | eq:'active' }}
// out: true
//
// example: check for an empty inbox
// in:  unread = 0
// tpl: {{ unread | eq:0 }}
// out: true
//
// example: check a feature flag
// in:  newsletter = true
// tpl: {{ newsletter | eq:true }}
// out: true
var Eq = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("eq function requires at least one parameter")
	}

	other := params[0]

	if value == nil || other == nil {
		return value == other, nil
	}

	valNum, errVal := functions.ValueNumber(value)
	otherNum, errOther := functions.ValueNumber(other)
	if errVal == nil && errOther == nil {
		return valNum == otherNum, nil
	}

	valStr, okVal := value.(string)
	otherStr, okOther := other.(string)
	if okVal && okOther {
		return valStr == otherStr, nil
	}

	return value == other, nil
}
