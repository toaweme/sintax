package query_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/path/query"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(query.Modifiers())).Render(tpl, vars)
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

// ExampleDirname_bareName returns "." when the path is a bare file name with no
// directory part, so it can still be joined back safely.
func ExampleDirname_bareName() {
	fmt.Println(render(`{{ path | dirname }}`, map[string]any{
		"path": "invoice.txt",
	}))
	// Output: .
}

// ExampleDirname_trailingSlash drops the trailing slash and returns the parent
// directory of the final segment.
func ExampleDirname_trailingSlash() {
	fmt.Println(render(`{{ path | dirname }}`, map[string]any{
		"path": "docs/reports/",
	}))
	// Output: docs/reports
}

// ExampleDirname_absolute keeps the leading slash for an absolute path.
func ExampleDirname_absolute() {
	fmt.Println(render(`{{ path | dirname }}`, map[string]any{
		"path": "/var/log/syslog",
	}))
	// Output: /var/log
}

// ExampleFilename returns the base file name, the final path element including
// its extension.
func ExampleFilename() {
	fmt.Println(render(`{{ path | filename }}`, map[string]any{
		"path": "docs/reports/q1.pdf",
	}))
	// Output: q1.pdf
}

// ExampleFilename_trailingSlash ignores a trailing slash and returns the last
// real segment of the path.
func ExampleFilename_trailingSlash() {
	fmt.Println(render(`{{ path | filename }}`, map[string]any{
		"path": "docs/reports/",
	}))
	// Output: reports
}

// ExampleFilename_bareName leaves a bare file name unchanged when there is no
// directory to strip.
func ExampleFilename_bareName() {
	fmt.Println(render(`{{ path | filename }}`, map[string]any{
		"path": "config.yaml",
	}))
	// Output: config.yaml
}

// ExampleFilename_unicode returns the final segment of a multi-byte path
// untouched.
func ExampleFilename_unicode() {
	fmt.Println(render(`{{ path | filename }}`, map[string]any{
		"path": "файлы/отчёт.pdf",
	}))
	// Output: отчёт.pdf
}

// ExampleFilenameExt returns the file extension without the leading dot.
func ExampleFilenameExt() {
	fmt.Println(render(`{{ path | ext }}`, map[string]any{
		"path": "avatar.png",
	}))
	// Output: png
}

// ExampleFilenameExt_compound returns only the part after the final dot, so a
// compound name yields its last extension.
func ExampleFilenameExt_compound() {
	fmt.Println(render(`{{ path | ext }}`, map[string]any{
		"path": "archive.tar.gz",
	}))
	// Output: gz
}

// ExampleFilenameExt_fullPath reads the extension from the final element of a
// multi-segment path, ignoring the directories that precede it.
func ExampleFilenameExt_fullPath() {
	fmt.Println(render(`{{ path | ext }}`, map[string]any{
		"path": "logs/2026/app.log",
	}))
	// Output: log
}

// ExampleFilenameExtDot returns the file extension including the leading dot,
// the counterpart to ext for callers rebuilding a name.
func ExampleFilenameExtDot() {
	fmt.Println(render(`{{ path | ext_dot }}`, map[string]any{
		"path": "avatar.png",
	}))
	// Output: .png
}

// ExampleFilenameExtDot_compound keeps the separator on the final extension of a
// compound name.
func ExampleFilenameExtDot_compound() {
	fmt.Println(render(`{{ path | ext_dot }}`, map[string]any{
		"path": "archive.tar.gz",
	}))
	// Output: .gz
}

// ExampleFilenameExtDot_fullPath reads the dotted extension from the final
// element of a multi-segment path, ready to append when rebuilding a name.
func ExampleFilenameExtDot_fullPath() {
	fmt.Println(render(`{{ path | ext_dot }}`, map[string]any{
		"path": "assets/img/logo.svg",
	}))
	// Output: .svg
}
