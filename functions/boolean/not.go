package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameNot is the template name for the Not modifier.
const ModifierNameNot functions.ModifierName = "not"

// Not inverts the truthiness of the value, following the same rules as a
// template `if`. Booleans go by their value, numbers are true when greater than
// zero, strings are true when non-empty (except the literal "false"), and
// collections are true when non-empty. An unrecognized type is falsey, so Not
// returns true for it.
//
// nil is the one value Not will not invert. A missing key and a key holding an
// explicit null both arrive here as nil, and inverting either one turns a
// misspelled path into a confident true. Say what absence means with a default
// ahead of it, as in `{{ flag | default:false | not }}`.
func Not(value any) (bool, error) {
	if value == nil {
		return false, functions.RefuseNil(ModifierNameNot)
	}
	return !functions.ConditionIsTrue(value), nil
}
