package sintax

import "errors"

// Sentinel errors returned by the parser and renderer.
var (
	ErrInvalidTokenType    = errors.New("invalid token")
	ErrVariableNotFound    = errors.New("variable not found")
	ErrFunctionNotFound    = errors.New("function not found")
	ErrFunctionApplyFailed = errors.New("function failed to apply")
	ErrMaxDepthExceeded    = errors.New("max template nesting depth exceeded")
)

// Sintax renders a template string against a variable set.
type Sintax interface {
	Render(template string, vars map[string]any) (any, error)
}

// Parser tokenizes a template string.
type Parser interface {
	Parse(template string) ([]Token, error)
}

// Renderer renders a token stream against a variable set.
type Renderer interface {
	Render(tokens []Token, vars map[string]any) (any, error)
}
