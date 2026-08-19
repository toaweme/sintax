package boolean

import "github.com/toaweme/sintax/functions"

// ModifierNameOr is the template name for the Or modifier.
const ModifierNameOr functions.ModifierName = "or"

// Or reports whether either the piped value or its argument is truthy, reading
// truthiness the way a template `if` does. It answers "is any of these set"
// without a chain of if blocks, and it composes with not, as in
// `{{ draft | not | or:published }}`.
//
// A miss arriving down the pipe is falsey and answered here rather than failing,
// so the argument decides the result on its own. Or returns a bool rather than
// the truthy operand, which makes it a condition rather than a fallback. Use
// default when what you want is a value.
func Or(value any, other any) (bool, error) {
	return functions.ConditionIsTrue(value) || functions.ConditionIsTrue(other), nil
}
