package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameNot is the template name for the Not modifier.
const ModifierNameNot functions.ModifierName = "not"

// Not inverts the truthiness of the value, following the same rules as a
// template `if`: booleans by their value, numbers true when greater than zero,
// strings true when non-empty (except the literal "false"), and collections
// true when non-empty. nil and any unrecognised type are falsey, so Not returns
// true for them.
func Not(value any) (bool, error) {
	return !functions.ConditionIsTrue(value), nil
}
