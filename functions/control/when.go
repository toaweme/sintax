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
// A nil value is refused rather than taken as the falsy branch. Absent data is
// not one of the two cases the template named, and picking a branch for it hides
// a misspelled path behind an answer that looks deliberate. Put a default ahead
// of the when to say which branch absence belongs to.
func When(value any, truthy any, falsy any) (any, error) {
	if value == nil {
		return nil, functions.RefuseNil(ModifierNameWhen)
	}
	if functions.ConditionIsTrue(value) {
		return truthy, nil
	}
	return falsy, nil
}
