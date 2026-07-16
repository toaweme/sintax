package fs_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/fs"
	sintaxrender "github.com/toaweme/sintax/functions/render"
)

// the contextual modifiers come along so ExampleFile_intoTemplate can pipe a
// file's contents into `template`.
func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(
		sintax.WithModifiers(fs.Modifiers([]string{"testdata"})),
		sintax.WithContextualModifiers(sintaxrender.ContextualModifiers()),
	).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleFile reads a file's contents from the caller-supplied allowlist of safe
// directories, here the local testdata dir.
func ExampleFile() {
	fmt.Println(render(`{{ p | file }}`, map[string]any{
		"p": "greeting.txt",
	}))
	// Output: hello from sintax
}

// ExampleFile_missing reports a not-found error when the path names no file
// inside any allowed directory.
func ExampleFile_missing() {
	fmt.Println(render(`{{ p | file }}`, map[string]any{
		"p": "does-not-exist.txt",
	}))
	// Output: error: failed to render template: failed to render variable token 'p': function failed to apply: failed to read file "does-not-exist.txt": file does not exist
}

// ExampleFile_traversal keeps a path from escaping the allowlist, so "../"
// segments cannot reach a parent directory and the read is refused.
func ExampleFile_traversal() {
	fmt.Println(render(`{{ p | file }}`, map[string]any{
		"p": "../secret.txt",
	}))
	// Output: error: failed to render template: failed to render variable token 'p': function failed to apply: failed to read file "../secret.txt": file does not exist
}

// ExampleFile_intoTemplate reads a file and renders its contents as a template,
// the intended way to embed an on-disk partial.
func ExampleFile_intoTemplate() {
	fmt.Println(render(`{{ p | file | template }}`, map[string]any{
		"p": "greeting.txt",
	}))
	// Output: hello from sintax
}
