package sintax

import (
	"maps"

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
	casing "github.com/toaweme/sintax/functions/text/case"
	textedit "github.com/toaweme/sintax/functions/text/edit"
	"github.com/toaweme/sintax/functions/text/splitjoin"
	"github.com/toaweme/sintax/functions/text/trim"
)

// builtins assembles the full built-in modifier set for the engine's own tests.
// It mirrors defaults.All, which the internal test package cannot import (that
// would be an import cycle, since defaults imports sintax).
func builtins(safeDirs ...string) map[string]GlobalModifier {
	groups := []map[string]GlobalModifier{
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

	all := make(map[string]GlobalModifier)
	for _, g := range groups {
		maps.Copy(all, g)
	}
	return all
}
