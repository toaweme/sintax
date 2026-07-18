// Package edit provides modifiers that rewrite a path's extension.
package edit

import (
	"path/filepath"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFilenameTrimExt is the template name for the FilenameTrimExt modifier.
const ModifierNameFilenameTrimExt functions.ModifierName = "ext_trim"

// ModifierNameFilenamePrependExt is the template name for the FilenamePrependExt modifier.
const ModifierNameFilenamePrependExt functions.ModifierName = "ext_prepend"

// FilenameTrimExt returns the file path without its trailing extension. Only the
// final extension is removed, so a path like "archive.tar.gz" keeps its "tar"
// segment and loses just ".gz". Directory segments are preserved untouched. A
// path with no extension, or one whose last element is a directory (trailing
// slash), comes back unchanged. A dotfile with nothing before the dot, such as
// ".gitignore", is treated by the standard library as being all extension, so
// trimming it yields an empty string.
func FilenameTrimExt(s string) (string, error) {
	oldExt := filepath.Ext(s)
	if oldExt == "" {
		return s, nil
	}
	return s[:len(s)-len(oldExt)], nil
}

// FilenamePrependExt inserts an extra extension segment just before the existing
// file extension, so "styles.css" with "min" becomes "styles.min.css". The
// inserted segment is taken literally and should not include a leading dot, since
// the dots are added for you. When the path has no extension, the segment is
// appended as a new extension ("noext" becomes "noext.min"). Only the final
// extension is considered, so "archive.tar.gz" becomes "archive.tar.min.gz".
func FilenamePrependExt(s, ext string) (string, error) {
	oldExt := filepath.Ext(s)
	if oldExt == "" {
		return s + "." + ext, nil
	}
	return s[:len(s)-len(oldExt)] + "." + ext + oldExt, nil
}
