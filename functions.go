package sintax

import (
	"github.com/toaweme/sintax/functions"
)

type GlobalModifier func(value any, params []any) (any, error)

// TODO: modifiers accept a value and a list of parameters, all being any
// this means that every modifier has to validate the input types
// not only is this ugly, but it's also error-prone
// we should use a custom type system to enforce the types of the parameters

type ModifierValue = any
type ModifierParam = any

// TODO: use this to return the error to the user
type ModifierError struct {
	Template       string
	Function       string
	Value          any
	Params         []any
	Error          error
	ErrorLocalized string
}

var BuiltinFunctions = map[string]GlobalModifier{
	// convert:
	"json": functions.JSON,
	"yaml": functions.YAML,
	"from": functions.From,

	// format:
	"default": functions.Default,
	"format":  functions.Format,
	"sexy":    functions.Sexy,
	"lines":   functions.Lines,
	"join":    functions.Join,
	"split":   functions.Split,
	"trim":    functions.Trim,
	"shorten": functions.Shorten,
	"length":  functions.Length,
	"first":   functions.First,
	"concat":  functions.Concat,

	// objects:
	"key":  functions.Key,
	"wrap": functions.Wrap,

	// boolean:
	"not": functions.Not,
}
