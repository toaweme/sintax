package sintax

import (
	"github.com/toaweme/sintax/functions"
)

type GlobalModifier func(value any, params []any) (any, error)

var BuiltinFunctions = map[string]GlobalModifier{
	"format":  functions.Format,
	"default": functions.Default,
	"json":    functions.JSON,
	"yaml":    functions.YAML,
	"from":    functions.From,
	"sexy":    functions.Sexy,
	"lines":   functions.Lines,
	"join":    functions.Join,
	"trim":    functions.Trim,
	"shorten": functions.Shorten,
	"length":  functions.Length,
	"first":   functions.First,
	"key":     functions.Key,
	"not":     functions.Not,
	"concat":  functions.Concat,
	"wrap":    functions.Wrap,
}
