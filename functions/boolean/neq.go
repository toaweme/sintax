package boolean

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameNeq is the template name for the Neq modifier.
const ModifierNameNeq functions.ModifierName = "neq"

// neq is an Overload mirroring eq clause for clause, and every clause answers by
// negating its eq counterpart rather than by comparing again, so the two can
// never drift into disagreeing about what equality means.

// NeqNumber reports numeric inequality across the int and float kinds, so 5 and
// 5.0 are equal and this returns false for them.
func NeqNumber(value, other float64) (bool, error) {
	equal, err := EqNumber(value, other)
	if err != nil {
		return false, fmt.Errorf("failed to compare two numbers: %w", err)
	}
	return !equal, nil
}

// NeqString reports verbatim string inequality, the neq clause reached when
// neither operand is numeric.
func NeqString(value, other string) (bool, error) {
	equal, err := EqString(value, other)
	if err != nil {
		return false, fmt.Errorf("failed to compare two strings: %w", err)
	}
	return !equal, nil
}

// NeqAny is neq's fallback for operands that are neither numeric nor both
// strings. A number and its string form reach here and are unequal, so neq
// reports true for them.
func NeqAny(value, other any) (bool, error) {
	equal, err := EqAny(value, other)
	if err != nil {
		return false, fmt.Errorf("failed to compare two values: %w", err)
	}
	return !equal, nil
}

// neqNilGuard runs ahead of the typed clauses, because NeqNumber would otherwise
// coerce nil to zero and report nil and 0 as equal.
//
// A nil piped value is refused for the same reason eq refuses one. Absent data
// compares unequal to everything, which is a verdict a misspelled path should not
// be able to produce. A nil argument came from data the template did name, and a
// non-nil value is unequal to it.
//
// When neither operand is nil the guard declines with ErrInvalidValueType so
// Overload falls through.
func neqNilGuard(value any, params []any) (any, error) {
	other, err := functions.ParamAny(params, 0)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, functions.RefuseNil(ModifierNameNeq)
	}
	if other == nil {
		return true, nil
	}
	return nil, functions.ErrInvalidValueType
}
