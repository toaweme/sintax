package sintax

import (
	"github.com/toaweme/sintax/functions/boolean"
	"github.com/toaweme/sintax/functions/collections/access"
	collquery "github.com/toaweme/sintax/functions/collections/query"
	"github.com/toaweme/sintax/functions/collections/transform"
	"github.com/toaweme/sintax/functions/control"
	"github.com/toaweme/sintax/functions/convert/parse"
	"github.com/toaweme/sintax/functions/convert/serialize"
	"github.com/toaweme/sintax/functions/escape"
	"github.com/toaweme/sintax/functions/format"
	"github.com/toaweme/sintax/functions/fs"
	pathedit "github.com/toaweme/sintax/functions/path/edit"
	pathquery "github.com/toaweme/sintax/functions/path/query"
	"github.com/toaweme/sintax/functions/render"
	casing "github.com/toaweme/sintax/functions/text/case"
	textedit "github.com/toaweme/sintax/functions/text/edit"
	"github.com/toaweme/sintax/functions/text/splitjoin"
	"github.com/toaweme/sintax/functions/text/trim"
)

// GlobalModifier is a stateless modifier that transforms a piped value given
// its call params, independent of any render context.
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
		string(render.ModifierNameTemplate): render.Template,
	}
}

// BuiltinFunctions builds the map of global modifiers, applying overrides and
// wiring safeDirs into any modifiers that need to constrain filesystem access.
var BuiltinFunctions = func(overrides map[string]GlobalModifier, safeDirs []string) map[string]GlobalModifier {
	funcs := map[string]GlobalModifier{
		// convert/serialize
		string(serialize.ModifierNameJSON): serialize.JSON,
		// the following built-in functions deliberately aren't implemented and return an error because
		// they depend on 3rd party libraries, and we don't want to bloat this package
		string(serialize.ModifierNameYAML):     serialize.YAML,
		string(serialize.ModifierNameMarkdown): serialize.Markdown,

		// convert/parse
		string(parse.ModifierNameFrom): parse.From,

		// escape
		string(escape.ModifierNameHTML): escape.HTML,
		string(escape.ModifierNameURL):  escape.URL,
		string(escape.ModifierNameJS):   escape.JS,

		// control
		string(control.ModifierNameDefault): control.Default,

		// format
		string(format.ModifierNameFormat):      format.Format,
		string(format.ModifierNameLength):      format.Length,
		string(format.ModifierNameLineNumbers): format.LineNumbers,
		string(format.ModifierNameDecimal):     format.Decimal,
		string(format.ModifierNameCurrency):    format.Currency,

		// path/query
		string(pathquery.ModifierNameDirname):        pathquery.Dirname,
		string(pathquery.ModifierNameFilename):       pathquery.Filename,
		string(pathquery.ModifierNameFilenameExt):    pathquery.FilenameExt,
		string(pathquery.ModifierNameFilenameExtDot): pathquery.FilenameExtDot,

		// path/edit
		string(pathedit.ModifierNameFilenamePrependExt): pathedit.FilenamePrependExt,
		string(pathedit.ModifierNameFilenameTrimExt):    pathedit.FilenameTrimExt,

		// fs
		string(fs.ModifierNameFile): fs.File(safeDirs),

		// text/case
		string(casing.ModifierNameToLower):    casing.ToLower,
		string(casing.ModifierNameToUpper):    casing.ToUpper,
		string(casing.ModifierNameSlug):       casing.Slug,
		string(casing.ModifierNameTitle):      casing.Title,
		string(casing.ModifierNameModelTitle): casing.ModelTitle,

		// text/trim
		string(trim.ModifierNameTrim):       trim.Trim,
		string(trim.ModifierNameTrimPrefix): trim.TrimPrefix,
		string(trim.ModifierNameTrimSuffix): trim.TrimSuffix,

		// text/edit
		string(textedit.ModifierNameShorten):        textedit.Shorten,
		string(textedit.ModifierNameConcat):         textedit.Concat,
		string(textedit.ModifierNameReplace):        textedit.Replace,
		string(textedit.ModifierNameReplacePattern): textedit.ReplacePattern,
		string(textedit.ModifierNameReverse):        textedit.Reverse,
		string(textedit.ModifierNameWrap):           textedit.Wrap,

		// text/splitjoin
		string(splitjoin.ModifierNameLines): splitjoin.Lines,
		string(splitjoin.ModifierNameJoin):  splitjoin.Join,
		string(splitjoin.ModifierNameSplit): splitjoin.Split,

		// collections/access
		string(access.ModifierNameFirst): access.First,
		string(access.ModifierNameLast):  access.Last,
		string(access.ModifierNameKey):   access.Key,
		string(access.ModifierNamePluck): access.Pluck,
		string(access.ModifierNameFind):  access.Find,

		// collections/query
		string(collquery.ModifierNameFilter): collquery.Filter,
		string(collquery.ModifierNameHas):    collquery.Has,
		string(collquery.ModifierNameIs):     collquery.Is,

		// collections/transform
		string(transform.ModifierNameMap):     transform.Map,
		string(transform.ModifierNameSort):    transform.Sort,
		string(transform.ModifierNameMerge):   transform.Merge,
		string(transform.ModifierNameSum):     transform.Sum,
		string(transform.ModifierNameFlatten): transform.Flatten,

		// boolean
		string(boolean.ModifierNameNot): boolean.Not,
		string(boolean.ModifierNameGt):  boolean.Gt,
		string(boolean.ModifierNameGte): boolean.Gte,
		string(boolean.ModifierNameEq):  boolean.Eq,
	}

	for name, fn := range overrides {
		funcs[name] = fn
	}

	return funcs
}
