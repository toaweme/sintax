// Package query provides modifiers that read components of a filesystem path.
package query

import (
	"path/filepath"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameDirname is the template name for the Dirname modifier.
const ModifierNameDirname functions.ModifierName = "dirname"

// Dirname returns the directory portion of a file path, everything up to but not
// including the final path element. A trailing slash means the final element is
// empty, so the path is returned without it, and a bare name with no directory
// yields "." so callers can join it back without special-casing.
func Dirname(s string) (string, error) {
	return filepath.Dir(s), nil
}
