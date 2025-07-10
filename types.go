package sintax

import "errors"

var (
	ErrInvalidTokenType    = errors.New("invalid token")
	ErrVariableNotFound    = errors.New("variable not found")
	ErrFunctionNotFound    = errors.New("function not found")
	ErrFunctionApplyFailed = errors.New("function failed to apply")
	ErrCircularDependency  = errors.New("circular dependency detected")
	ErrParseFailed         = errors.New("failed to parse template")
	ErrRenderFailed        = errors.New("failed to render template")
)

type Parser interface {
	Parse(template string) ([]Token, error)
}

type Renderer interface {
	Render(tokens []Token, vars map[string]any) (any, error)
}

type Syntax interface {
	ResolveVariables(vars map[string]any) (map[string]any, error)
	ResolveCondition(condition string, vars map[string]any) (bool, error)
	ExtractDependencies(vars map[string]any) ([]string, error)
	Render(input string, vars map[string]any) (any, error)
}
