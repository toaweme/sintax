package sintax

import (
	"errors"
	"fmt"
	"strings"
)

// Sentinel errors returned by the parser and renderer.
var (
	ErrInvalidTokenType    = errors.New("invalid token")
	ErrVariableNotFound    = errors.New("variable not found")
	ErrFunctionNotFound    = errors.New("function not found")
	ErrFunctionApplyFailed = errors.New("function failed to apply")
	ErrMaxDepthExceeded    = errors.New("max template nesting depth exceeded")
	ErrUnterminatedIf      = errors.New("missing endif")
	ErrUnterminatedFor     = errors.New("missing endfor")
	ErrUnexpectedToken     = errors.New("unexpected control token")
	ErrInvalidForExpr      = errors.New("invalid for expression")
	ErrNotIterable         = errors.New("value is not iterable")
	ErrNotAnExpression     = errors.New("not a variable expression")
)

// Position is where a token was written in the template it was parsed from.
// Line and Column count from 1, and are both 0 where the position is unknown,
// which is a token the engine built rather than parsed.
type Position struct {
	Line   int
	Column int
	// Text is the whole line the token sits on, without its line ending, which
	// is what an editor underlines Column into.
	Text string
}

// Known reports whether the position was recorded.
func (p Position) Known() bool { return p.Line > 0 }

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
	// Source is the template the failing modifier was written in, which is the
	// nested template rather than the outer one where a modifier re-entered the
	// engine. It is empty where the position is unknown.
	Source string
	// Position is where in Source the modifier's variable was written.
	Position Position
}

var _ error = (*ModifierError)(nil)

// Error names the line and column only where the template has more than one
// line. On a single-line template there is one position and naming it says
// nothing, so the fields carry it and the message does not.
func (e *ModifierError) Error() string {
	if !e.Position.Known() || !strings.Contains(e.Source, "\n") {
		return fmt.Sprintf("modifier %q: %v", e.Modifier, e.Err)
	}

	return fmt.Sprintf("modifier %q at %d:%d: %v", e.Modifier, e.Position.Line, e.Position.Column, e.Err)
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
