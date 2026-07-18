package sintax

import (
	"errors"
	"fmt"
)

// Sentinel errors returned by the parser and renderer.
var (
	ErrInvalidTokenType    = errors.New("invalid token")
	ErrVariableNotFound    = errors.New("variable not found")
	ErrFunctionNotFound    = errors.New("function not found")
	ErrFunctionApplyFailed = errors.New("function failed to apply")
	ErrMaxDepthExceeded    = errors.New("max template nesting depth exceeded")
)

// ModifierError reports a modifier that failed while rendering a variable's
// pipeline. A chain such as `{{ text | trim | upper:'z' | lower }}` has several
// places to fail, and the message alone cannot say which one did, so the failing
// modifier's name is carried as a field. Callers that need to act on it (an
// editor underlining the offending link, a client reporting the fault upstream)
// should reach it with errors.As rather than by parsing the message, which is
// not part of the contract and will change.
//
// Err wraps ErrFunctionApplyFailed over the modifier's own failure, so errors.Is
// still finds that sentinel and any sentinel beneath it.
type ModifierError struct {
	// Modifier is the template name of the modifier that failed, such as "upper".
	Modifier string
	// Variable is the name of the variable whose pipeline the modifier ran in.
	Variable string
	// Err is the failure the modifier reported.
	Err error
}

var _ error = (*ModifierError)(nil)

func (e *ModifierError) Error() string {
	return fmt.Sprintf("modifier %q: %v", e.Modifier, e.Err)
}

// Unwrap exposes the underlying failure to errors.Is and errors.As.
func (e *ModifierError) Unwrap() error { return e.Err }

// Sintax renders a template string against a variable set.
type Sintax interface {
	Render(template string, vars map[string]any) (any, error)
	RenderString(template string, vars map[string]any) (string, error)
}

// Parser tokenizes a template string.
type Parser interface {
	Parse(template string) ([]Token, error)
}

// Renderer renders a token stream against a variable set.
type Renderer interface {
	Render(tokens []Token, vars map[string]any) (any, error)
}
