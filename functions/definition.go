package functions

import "errors"

// ModifierName is the template-syntax name used to invoke a modifier (e.g. "lower", "trim-prefix").
type ModifierName string

// Type represents the kind of value a modifier accepts or returns.
type Type string

const (
	TypeAny    Type = "any"
	TypeString Type = "string"
	TypeBytes  Type = "bytes"
	TypeInt    Type = "int"
	TypeFloat  Type = "float"
	TypeBool   Type = "bool"
	TypeArray  Type = "array"
	TypeMap    Type = "map"
)

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
