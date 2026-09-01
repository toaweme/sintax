// Package control provides template modifiers that steer value resolution
// rather than transform data, such as falling back when a value is missing.
package control

import "github.com/toaweme/sintax/functions"

// ModifierNameDefault is the template name for the Default modifier.
const ModifierNameDefault functions.ModifierName = "default"

// Default supplies a fallback so a template never renders a missing value. It
// swaps in the fallback in exactly two situations and passes the real value
// through untouched otherwise.
//
// It applies the fallback when the piped value is nil (an absent or null
// variable) or an empty string. An empty list, an empty object, the number zero,
// and the boolean false are all real values and are kept as-is. `default` guards
// against "nothing there", not against "a value that happens to be empty".
//
// The second situation is a miss earlier in the same pipe, such as find matching
// no row or key not holding the requested field. A miss travels down the
// pipeline as nil, and default answers it simply by accepting that nil and
// returning the fallback. The engine gives this modifier no special standing.
// Any modifier that makes sense of nil answers a miss the same way.
//
// Where it does earn its place is ahead of a modifier that turns a value into a
// decision. `not`, `when`, `and`, `or`, `eq`, `neq`, `has` and `is` refuse nil
// rather than answer on it, so a default in front of one is how a template says
// what absence means, as in `{{ flag | default:false | not }}`. Written after the
// decision it is too late, because the decision already failed.
func Default(value any, fallback any) (any, error) {
	if value == nil || value == "" {
		return fallback, nil
	}
	return value, nil
}
