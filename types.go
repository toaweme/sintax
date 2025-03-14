package sintax

import "errors"

var ErrInvalidTokenType = errors.New("invalid token")
var ErrVariableNotFound = errors.New("variable not found")
var ErrFunctionNotFound = errors.New("function not found")
var ErrFunctionApplyFailed = errors.New("function failed to apply")
var ErrCircularDependency = errors.New("circular dependency detected at 'self'")

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
