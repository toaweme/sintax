// Package fs provides template modifiers for filesystem path manipulation.
package fs

import (
	"path/filepath"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameDirname is the template name for the Dirname modifier.
const ModifierNameDirname functions.ModifierName = "dirname"

// Dirname returns the directory portion of a file path.
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
func Dirname(value any, params []any) (any, error) {
	path, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}

	return filepath.Dir(path), nil
}
