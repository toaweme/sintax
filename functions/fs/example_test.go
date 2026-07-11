package fs_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/fs"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(fs.Modifiers([]string{"testdata"})).Render(tpl, vars)
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
