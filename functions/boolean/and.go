package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameAnd is the template name for the And modifier.
const ModifierNameAnd functions.ModifierName = "and"

// And reports whether the piped value and its argument are both truthy, reading
// truthiness the way a template `if` does. It is what lets one condition be
// written as one expression, so `{{ enabled | and:has_rows }}` replaces a nested
// pair of if blocks.
//
// A miss arriving down the pipe is falsey and answered here rather than failing,
// so a flag that was never set makes the whole condition false. An argument
// naming a variable that does not exist is a different thing and still fails,
// since the template asked for something the data was never going to hold.
// Both operands are evaluated, because an argument is resolved before the
// modifier runs, so there is no short circuit to lean on.
func And(value any, other any) (bool, error) {
	return functions.ConditionIsTrue(value) && functions.ConditionIsTrue(other), nil
}
