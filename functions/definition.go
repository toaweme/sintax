// Package functions defines the modifier registry and shared types used by
// the built-in template modifiers under functions/*.
package functions

import (
	"errors"
	"fmt"
)

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

// Miss reports that the data a modifier went looking for was not there, which
// is the condition the default modifier exists to catch. A lookup that matched
// no element, a collection with no element to take, and a file that is not in
// the allowlist are all misses. The counterpart is a type or arity rejection,
// which says the template itself is wrong and must stay terminal no matter what
// the pipeline holds downstream.
//
// The returned error reads as its own message while still matching
// errors.Is(err, ErrAllowsDefaultFunc). Wrapping the sentinel with %w directly
// would splice "non-fatal error" into text a template author reads, and it says
// nothing they can act on.
//
// The format takes %w, so a miss can carry a cause without giving up the marker:
// Miss("failed to read file %q: %w", path, os.ErrNotExist) matches errors.Is
// against both ErrAllowsDefaultFunc and os.ErrNotExist.
func Miss(format string, args ...any) error {
	return &missError{inner: fmt.Errorf(format, args...)}
}

// missError keeps ErrAllowsDefaultFunc reachable through Unwrap rather than
// through the message text, alongside whatever the caller wrapped with %w.
// Returning both from Unwrap means errors.Is finds the marker and the cause,
// which a single-error Unwrap could not do without one shadowing the other.
type missError struct{ inner error }

var _ error = (*missError)(nil)

func (e *missError) Error() string { return e.inner.Error() }

func (e *missError) Unwrap() []error { return []error{ErrAllowsDefaultFunc, e.inner} }

var (
	// ErrAllowsDefaultFunc marks non-fatal errors that the default modifier can
	// catch. Report one with Miss rather than wrapping this directly, so the
	// marker stays out of the message.
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

// IsParamError reports whether err is a modifier rejecting its params rather
// than its value. The distinction matters while a miss is traveling down a
// pipeline. A modifier handed the nil that stands in for absent data may
// reasonably reject it, which says nothing about the template, but a param is
// written in the template and does not depend on the data, so `upper:'x'` is
// wrong whether or not the value arrived. Without this, a mistake like that
// would hide behind a default for exactly as long as the data stayed absent.
//
// It deliberately asks only about params. A value rejection has many shapes,
// some of them bare errors carrying no sentinel at all, so treating "not a param
// error" as "declining the value" keeps an unrecognized rejection catchable
// rather than promoting it to a hard failure on a guess.
func IsParamError(err error) bool {
	return errors.Is(err, ErrInvalidParamType) ||
		errors.Is(err, ErrMissingParam) ||
		errors.Is(err, ErrInvalidParamValue)
}
