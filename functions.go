package sintax

import (
	"github.com/toaweme/sintax/functions/boolean"
	"github.com/toaweme/sintax/functions/collections"
	"github.com/toaweme/sintax/functions/convert"
	"github.com/toaweme/sintax/functions/fs"
	"github.com/toaweme/sintax/functions/money"
	"github.com/toaweme/sintax/functions/text"
	"github.com/toaweme/sintax/functions/utils"
)

type GlobalModifier func(value any, params []any) (any, error)

// ContextualModifier is a modifier that needs live render state - the current
// variables and a re-entrant renderer - rather than only its piped value. The
// render callback runs a string template through the same engine (same
// modifiers, same recursion guard). The callback type is written inline so
// modifiers in functions/ subpackages stay structurally compatible without
// importing this package.
//
// The vars map must be treated as read-only and borrowed: read it during the
// call, but do not mutate it or retain a reference past return. Inside a `for`
// body the engine reuses a single scope map across iterations, so a retained
// reference would observe later iterations' values rather than a stable
// snapshot. Copy out anything that must outlive the call.
type ContextualModifier func(render func(template string, vars map[string]any) (any, error), vars map[string]any, value any, params []any) (any, error)

// builtinContextualModifiers returns the contextual modifiers wired into every
// renderer. Unlike GlobalModifiers these need live render state, so they are
// registered here rather than exposed through BuiltinFunctions.
func builtinContextualModifiers() map[string]ContextualModifier {
	return map[string]ContextualModifier{
		string(text.ModifierNameTemplate): text.Template,
	}
}

var BuiltinFunctions = func(overrides map[string]GlobalModifier, safeDirs []string) map[string]GlobalModifier {
	funcs := map[string]GlobalModifier{
		// convert
		string(convert.ModifierNameJSON): convert.JSON,
		string(convert.ModifierNameFrom): convert.From,
		// the following built-in functions deliberately aren't implemented and return an error because
		// they depend on 3rd party libraries, and we don't want to bloat this package
		string(convert.ModifierNameYAML): convert.YAML,
		// string(convert.ModifierNameMarkdown): convert.Markdown,

		// utils
		string(utils.ModifierNameDefault):     utils.Default,
		string(utils.ModifierNameFormat):      utils.Format,
		string(utils.ModifierNameLength):      utils.Length,
		string(utils.ModifierNameLineNumbers): utils.LineNumbers,
		string(utils.ModifierNameDecimal):     utils.Decimal,

		// fs
		string(fs.ModifierNameDirname):            fs.Dirname,
		string(fs.ModifierNameFilename):           fs.Filename,
		string(fs.ModifierNameFilenameExt):        fs.FilenameExt,
		string(fs.ModifierNameFilenameExtDot):     fs.FilenameExtDot,
		string(fs.ModifierNameFilenamePrependExt): fs.FilenamePrependExt,
		string(fs.ModifierNameFilenameTrimExt):    fs.FilenameTrimExt,
		string(fs.ModifierNameFile):               fs.File(safeDirs),

		// text
		// string(text.ModifierNameSexy): text.Sexy,
		string(text.ModifierNameLines):          text.Lines,
		string(text.ModifierNameJoin):           text.Join,
		string(text.ModifierNameSplit):          text.Split,
		string(text.ModifierNameTrim):           text.Trim,
		string(text.ModifierNameTrimPrefix):     text.TrimPrefix,
		string(text.ModifierNameTrimSuffix):     text.TrimSuffix,
		string(text.ModifierNameShorten):        text.Shorten,
		string(text.ModifierNameConcat):         text.Concat,
		string(text.ModifierNameSlug):           text.Slug,
		string(text.ModifierNameTitle):          text.Title,
		string(text.ModifierNameModelTitle):     text.ModelTitle,
		string(text.ModifierNameToLower):        text.ToLower,
		string(text.ModifierNameToUpper):        text.ToUpper,
		string(text.ModifierNameReplace):        text.Replace,
		string(text.ModifierNameReplacePattern): text.ReplacePattern,
		string(text.ModifierNameReverse):        text.Reverse,

		// collections
		string(collections.ModifierNameFirst):   collections.First,
		string(collections.ModifierNameLast):    collections.Last,
		string(collections.ModifierNameFind):    collections.Find,
		string(collections.ModifierNameFilter):  collections.Filter,
		string(collections.ModifierNameHas):     collections.Has,
		string(collections.ModifierNameIs):      collections.Is,
		string(collections.ModifierNameKey):     collections.Key,
		string(collections.ModifierNameMap):     collections.Map,
		string(collections.ModifierNameWrap):    collections.Wrap,
		string(collections.ModifierNameSort):    collections.Sort,
		string(collections.ModifierNameMerge):   collections.Merge,
		string(collections.ModifierNameSum):     collections.Sum,
		string(collections.ModifierNamePluck):   collections.Pluck,
		string(collections.ModifierNameFlatten): collections.Flatten,

		// boolean
		string(boolean.ModifierNameNot): boolean.Not,
		string(boolean.ModifierNameGt):  boolean.Gt,
		string(boolean.ModifierNameGte): boolean.Gte,
		string(boolean.ModifierNameEq):  boolean.Eq,

		// money
		string(money.ModifierNameCurrency): money.Currency,
	}

	for name, fn := range overrides {
		funcs[name] = fn
	}

	return funcs
}
