package control

import "github.com/toaweme/sintax/functions"

// ModifierNameWhen is the template name for the When modifier.
const ModifierNameWhen functions.ModifierName = "when"

// When picks between two values by the truthiness of the piped one, reading
// truthiness the way a template `if` does. It is what lets a flag render as the
// word it means, so `{{ published | when:'live','draft' }}` says in one
// expression what a pair of string replacements said by accident.
//
// Both arguments are required, and either may be a literal or a variable. A
// missing second argument is a template mistake rather than a shorthand for an
// empty string, since a choice with one branch is a choice somebody forgot to
// finish.
//
// A miss arriving down the pipe is falsey and When answers it, because absent
// data is one of the two cases the template already named. The falsy branch is
// the result and the miss stops traveling, which means a default written after a
// when never fires. Put the default first when the fallback belongs to the value
// rather than to the choice.
func When(value any, truthy any, falsy any) (any, error) {
	if functions.ConditionIsTrue(value) {
		return truthy, nil
	}
	return falsy, nil
}
