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
	ParseVariable(val string) ([]Token, error)
}

type Renderer interface {
	RenderString(tokens []Token, vars map[string]any) (string, error)
	Render(tokens []Token, vars map[string]any) (any, error)
}

type Syntax interface {
	// ResolveVariables resolves all variables in the given system, config, and action variables.
	ResolveVariables(vars map[string]any) (map[string]any, error)
	Render(input string, vars map[string]any) (string, error)
}
