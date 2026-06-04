package sintax

import "errors"

var (
	ErrInvalidTokenType    = errors.New("invalid token")
	ErrVariableNotFound    = errors.New("variable not found")
	ErrFunctionNotFound    = errors.New("function not found")
	ErrFunctionApplyFailed = errors.New("function failed to apply")
	ErrMaxDepthExceeded    = errors.New("max template nesting depth exceeded")
)

type Sintax interface {
	Render(template string, vars map[string]any) (any, error)
}

type Parser interface {
	Parse(template string) ([]Token, error)
}

type Renderer interface {
	Render(tokens []Token, vars map[string]any) (any, error)
}
