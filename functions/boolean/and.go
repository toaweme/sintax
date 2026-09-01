package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameAnd is the template name for the And modifier.
const ModifierNameAnd functions.ModifierName = "and"

// And reports whether the piped value and its argument are both truthy, reading
// truthiness the way a template `if` does. It is what lets one condition be
// written as one expression, so `{{ enabled | and:has_rows }}` replaces a nested
// pair of if blocks.
//
// A nil value is refused rather than read as false, since a flag that was never
// set and a flag whose path was misspelled look the same by the time they get
// here. Put a default ahead of it to say what absence means. An argument naming a
// variable that does not exist fails for its own reason, since the template asked
// for something the data was never going to hold. Both operands are evaluated,
// because an argument is resolved before the modifier runs, so there is no short
// circuit to lean on.
func And(value any, other any) (bool, error) {
	if value == nil {
		return false, functions.RefuseNil(ModifierNameAnd)
	}
	return functions.ConditionIsTrue(value) && functions.ConditionIsTrue(other), nil
}
