// Package boolean provides template modifiers for boolean comparisons.
package boolean

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameEq is the template name for the Eq modifier.
const ModifierNameEq functions.ModifierName = "eq"

// Eq returns true if the value equals the given parameter. Numbers compare
// numerically across the int and float kinds, so the integer 5 and the float
// 5.0 are equal. Strings compare verbatim, and any other pair falls back to
// Go's direct equality. Note that a number and its string form are never equal,
// so 5 does not equal "5". nil is only equal to nil, so comparing nil against 0
// yields false.
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
//
// example: a mismatched value is not equal
// in:  status = "active"
// tpl: {{ status | eq:'archived' }}
// out: false
var Eq = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errors.New("eq function requires at least one parameter")
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
