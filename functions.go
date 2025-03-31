package sintax

import (
	"github.com/toaweme/sintax/functions/boolean"
	"github.com/toaweme/sintax/functions/collections"
	"github.com/toaweme/sintax/functions/convert"
	"github.com/toaweme/sintax/functions/money"
	"github.com/toaweme/sintax/functions/text"
	"github.com/toaweme/sintax/functions/utils"
)

type GlobalModifier func(value any, params []any) (any, error)

var BuiltinFunctions = map[string]GlobalModifier{
	// convert:
	"json": convert.JSON,
	"yaml": convert.YAML,
	"from": convert.From,

	// utils:
	"default": utils.Default,
	"format":  utils.Format,
	"length":  utils.Length,

	// text
	"sexy":    text.Sexy,
	"lines":   text.Lines,
	"join":    text.Join,
	"split":   text.Split,
	"trim":    text.Trim,
	"shorten": text.Shorten,
	"concat":  text.Concat,
	"slug":    text.Slug,
	"title":   text.Title,
	// "lower":   utils.Lower,
	// "upper":   utils.Upper,
	// "replace": utils.Replace,
	// "reverse": utils.Reverse,

	// objects:
	"first": collections.First,
	"find":  collections.Find,
	"key":   collections.Key,
	"map":   collections.Map,
	"wrap":  collections.Wrap,

	// boolean:
	"not": boolean.Not,
	"gt":  boolean.Gt,
	"gte": boolean.Gte,

	// money:
	"currency": money.Currency,
}
