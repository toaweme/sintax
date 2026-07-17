package edit_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/path/edit"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(edit.Modifiers())).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleFilenameTrimExt returns the path without its trailing extension,
// removing only the final segment and preserving the directory portion.
func ExampleFilenameTrimExt() {
	fmt.Println(render(`{{ path | ext_trim }}`, map[string]any{
		"path": "docs/report.pdf",
	}))
	// Output: docs/report
}

// ExampleFilenameTrimExt_multiExtension removes only the final extension, so a
// double-suffixed archive keeps its inner segment.
func ExampleFilenameTrimExt_multiExtension() {
	fmt.Println(render(`{{ path | ext_trim }}`, map[string]any{
		"path": "backup/archive.tar.gz",
	}))
	// Output: backup/archive.tar
}

// ExampleFilenameTrimExt_noExtension leaves a path that has no extension exactly
// as it was.
func ExampleFilenameTrimExt_noExtension() {
	fmt.Println(render(`{{ path | ext_trim }}`, map[string]any{
		"path": "bin/tool",
	}))
	// Output: bin/tool
}

// ExampleFilenameTrimExt_hiddenDirectory trims the extension from a file that
// lives inside a leading-dot directory, keeping the hidden directory intact.
func ExampleFilenameTrimExt_hiddenDirectory() {
	fmt.Println(render(`{{ path | ext_trim }}`, map[string]any{
		"path": ".config/settings.json",
	}))
	// Output: .config/settings
}

// ExampleFilenamePrependExt inserts an extra extension segment just before the
// existing extension, so a minified name keeps its original suffix.
func ExampleFilenamePrependExt() {
	fmt.Println(render(`{{ path | ext_prepend:'min' }}`, map[string]any{
		"path": "styles.css",
	}))
	// Output: styles.min.css
}

// ExampleFilenamePrependExt_noExtension appends the segment as a new extension
// when the path has none to begin with.
func ExampleFilenamePrependExt_noExtension() {
	fmt.Println(render(`{{ path | ext_prepend:'tmp' }}`, map[string]any{
		"path": "backup",
	}))
	// Output: backup.tmp
}

// ExampleFilenamePrependExt_multiExtension inserts the segment before only the
// final extension, leaving earlier suffixes in place.
func ExampleFilenamePrependExt_multiExtension() {
	fmt.Println(render(`{{ path | ext_prepend:'bak' }}`, map[string]any{
		"path": "archive.tar.gz",
	}))
	// Output: archive.tar.bak.gz
}

// ExampleFilenamePrependExt_directoryPath keeps the directory portion intact
// while rewriting only the final filename's extension.
func ExampleFilenamePrependExt_directoryPath() {
	fmt.Println(render(`{{ path | ext_prepend:'en' }}`, map[string]any{
		"path": "public/site/index.html",
	}))
	// Output: public/site/index.en.html
}
