// Package defaults provides the batteries-included set of sintax modifiers.
// Pass All() to sintax.New for the whole battery. Importing this package links
// every built-in modifier; consumers that want a smaller binary should compose
// only the modifier groups they use instead (each group under functions/*
// exposes a Modifiers() constructor).
package defaults

import (
	"maps"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions"
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

// New returns every built-in global modifier keyed by its template name. Pass
// one or more safeDirs to enable the `file` modifier against that allowlist;
// with no dirs, file reads stay disabled.
func New(safeDirs ...string) map[string]functions.GlobalModifier {
	groups := []map[string]functions.GlobalModifier{
		casing.Modifiers(),
		trim.Modifiers(),
		textedit.Modifiers(),
		splitjoin.Modifiers(),
		access.Modifiers(),
		collquery.Modifiers(),
		transform.Modifiers(),
		serialize.Modifiers(),
		parse.Modifiers(),
		format.Modifiers(),
		boolean.Modifiers(),
		escape.Modifiers(),
		pathquery.Modifiers(),
		pathedit.Modifiers(),
		control.Modifiers(),
		fs.Modifiers(safeDirs),
	}

	all := make(map[string]functions.GlobalModifier)
	for _, g := range groups {
		maps.Copy(all, g)
	}
	return all
}

// Contextual returns the built-in contextual modifiers, which need live render
// state rather than only their piped value. They are wired separately from the
// global set because sintax.New takes the two through different options.
func Contextual() map[string]functions.ContextualModifier {
	return render.ContextualModifiers()
}

// All bundles every built-in modifier, global and contextual alike, into a
// single option for sintax.New. Pass one or more safeDirs to enable the `file`
// modifier against that allowlist; with none, file reads stay disabled.
//
// Options merge in order, so layering your own on top replaces a built-in of
// the same name:
//
//	s := sintax.New(defaults.All(), sintax.WithModifiers(overrides))
func All(safeDirs ...string) sintax.Option {
	return sintax.WithOptions(
		sintax.WithModifiers(New(safeDirs...)),
		sintax.WithContextualModifiers(Contextual()),
	)
}
