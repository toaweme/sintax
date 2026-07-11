package query_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/path/query"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(query.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleDirname returns the directory portion of a path, everything up to but
// not including the final element.
func ExampleDirname() {
	fmt.Println(render(`{{ path | dirname }}`, map[string]any{
		"path": "docs/reports/q1.pdf",
	}))
	// Output: docs/reports
}

// ExampleFilename returns the base file name, the final path element including
// its extension.
func ExampleFilename() {
	fmt.Println(render(`{{ path | filename }}`, map[string]any{
		"path": "docs/reports/q1.pdf",
	}))
	// Output: q1.pdf
}

// ExampleFilenameExt returns the file extension without the leading dot.
func ExampleFilenameExt() {
	fmt.Println(render(`{{ path | ext }}`, map[string]any{
		"path": "avatar.png",
	}))
	// Output: png
}

// ExampleFilenameExtDot returns the file extension including the leading dot,
// the counterpart to ext for callers rebuilding a name.
func ExampleFilenameExtDot() {
	fmt.Println(render(`{{ path | ext_dot }}`, map[string]any{
		"path": "avatar.png",
	}))
	// Output: .png
}
