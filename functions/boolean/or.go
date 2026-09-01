package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameOr is the template name for the Or modifier.
const ModifierNameOr functions.ModifierName = "or"

// Or reports whether either the piped value or its argument is truthy, reading
// truthiness the way a template `if` does. It answers "is any of these set"
// without a chain of if blocks, and it composes with not, as in
// `{{ draft | not | or:published }}`.
//
// A nil value is refused rather than read as false, because a missing key and an
// explicit null are indistinguishable from a path nobody ever spelled right, and
// answering one of those turns the argument into the whole condition by accident.
// Put a default ahead of it to say what absence means. Or returns a bool rather
// than the truthy operand, which makes it a condition rather than a fallback. Use
// default when what you want is a value.
func Or(value any, other any) (bool, error) {
	if value == nil {
		return false, functions.RefuseNil(ModifierNameOr)
	}
	return functions.ConditionIsTrue(value) || functions.ConditionIsTrue(other), nil
}
