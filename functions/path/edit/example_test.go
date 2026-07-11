package edit_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/path/edit"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(edit.Modifiers()).Render(tpl, vars)
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

// ExampleFilenamePrependExt inserts an extra extension segment just before the
// existing extension, so a minified name keeps its original suffix.
func ExampleFilenamePrependExt() {
	fmt.Println(render(`{{ path | ext_prepend:'min' }}`, map[string]any{
		"path": "styles.css",
	}))
	// Output: styles.min.css
}
