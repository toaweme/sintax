// Package query provides modifiers that read components of a filesystem path.
package query

import (
	"path/filepath"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameDirname is the template name for the Dirname modifier.
const ModifierNameDirname functions.ModifierName = "dirname"

// Dirname returns the directory portion of a file path, that is everything up
// to but not including the final path element. A path that already ends in a
// slash has an empty final element, so dirname returns the path itself without
// the trailing slash. A path with no directory part yields "." (the current
// directory), and the input must be a string.
//
// value: string
// returns: string
//
// example: get the folder of a file path
// in:  file_path = "/var/log/app/server.log"
// tpl: {{ file_path | dirname }}
// out: /var/log/app
//
// example: strip the file from a relative path
// in:  source_file = "src/handlers/users.go"
// tpl: {{ source_file | dirname }}
// out: src/handlers
//
// example: a bare file name has no directory, so the result is "."
// in:  name = "file.txt"
// tpl: {{ name | dirname }}
// out: .
//
// example: a trailing slash means the path is already a directory
// in:  dir = "/foo/bar/"
// tpl: {{ dir | dirname }}
// out: /foo/bar
func Dirname(value any, params []any) (any, error) {
	path, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}

	return filepath.Dir(path), nil
}
