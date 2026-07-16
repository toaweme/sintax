// Package functions defines the modifier registry and shared types used by
// the built-in template modifiers under functions/*.
package functions

import "errors"

// ModifierName is the template-syntax name used to invoke a modifier (e.g. "lower", "trim_prefix").
type ModifierName string

// GlobalModifier is a stateless modifier that transforms a piped value given
// its call params, independent of any render context. It lives here, in the
// shared base package, so every modifier subpackage and the batteries-included
// defaults package can name it without importing the engine.
type GlobalModifier func(value any, params []any) (any, error)

// ContextualModifier is a modifier that needs live render state - the current
// variables and a re-entrant renderer - rather than only its piped value. The
// render callback runs a string template through the same engine (same
// modifiers, same recursion guard).
//
// The vars map must be treated as read-only and borrowed: read it during the
// call, but do not mutate it or retain a reference past return. Inside a `for`
// body the engine reuses a single scope map across iterations, so a retained
// reference would observe later iterations' values rather than a stable
// snapshot. Copy out anything that must outlive the call.
type ContextualModifier func(render func(template string, vars map[string]any) (any, error), vars map[string]any, value any, params []any) (any, error)

var (
	// ErrAllowsDefaultFunc marks non-fatal errors that the default modifier can catch.
	ErrAllowsDefaultFunc = errors.New("non-fatal error")

	// ErrInvalidValueType is returned when a modifier receives an unsupported value type.
	ErrInvalidValueType = errors.New("invalid value type")

	// ErrInvalidParamType is returned when a modifier parameter has the wrong type.
	ErrInvalidParamType = errors.New("invalid param type")

	// ErrMissingParam is returned when a required modifier parameter is absent.
	ErrMissingParam = errors.New("missing param")

	// ErrInvalidParamValue is returned when a parameter value does not meet constraints.
	ErrInvalidParamValue = errors.New("invalid param value")
)
