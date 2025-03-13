package sintax

import "errors"

var ErrInvalidToken = errors.New("invalid token")
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
	// systemVars are variables that are always available to the pipeline e.g. env, now, etc.
	// configVars are variables defined in the pipeline config
	// actionVars are variables defined in the action
	// previousOutputVars are variables that were output from the previous action
	ResolveVariables(systemVars map[string]any, configVars map[string]any, actionVars map[string]any, previousOutputVars map[string]any) (map[string]any, error)
	Render(input string, vars map[string]any) (string, error)
}
