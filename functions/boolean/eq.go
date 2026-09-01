// Package boolean provides template modifiers for boolean comparisons.
package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameEq is the template name for the Eq modifier.
const ModifierNameEq functions.ModifierName = "eq"

// eq is an Overload of clauses reflecting how equality is decided: numbers
// compare numerically across the int and float kinds (5 equals 5.0), strings
// compare verbatim, and any other pair falls back to Go equality, so a number
// and its string form are never equal. nil is equal only to nil, which the guard
// clause enforces ahead of the typed clauses.

// EqNumber reports numeric equality across the int and float kinds, so 5 equals
// 5.0. Numbers and their string form are never equal, so 5 does not equal "5".
func EqNumber(value, other float64) (bool, error) {
	return value == other, nil
}

// EqString reports verbatim string equality, the eq clause reached when neither
// operand is numeric.
func EqString(value, other string) (bool, error) {
	return value == other, nil
}

// EqAny is eq's fallback for operands that are neither numeric nor both strings
// (for example two bools), comparing with Go equality. A number and its string
// form reach here and are unequal because the numeric and string clauses each
// reject the mixed pair first.
func EqAny(value, other any) (bool, error) {
	return value == other, nil
}

// eqNilGuard runs ahead of the typed clauses, because EqNumber would otherwise
// coerce nil to zero and compare it numerically.
//
// A nil piped value is refused. `{{ row.status | eq:'published' }}` over a column
// nobody spelled right answers "not published" with the same confidence as a real
// row, and the answer is usually inverted straight afterwards, so the refusal is
// what keeps a typo from reading as a verdict. A nil argument is a different
// thing, since it came from data the template did name, and a non-nil value is
// simply unequal to it.
//
// When neither operand is nil the guard declines with ErrInvalidValueType so
// Overload falls through.
func eqNilGuard(value any, params []any) (any, error) {
	other, err := functions.ParamAny(params, 0)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, functions.RefuseNil(ModifierNameEq)
	}
	if other == nil {
		return false, nil
	}
	return nil, functions.ErrInvalidValueType
}
