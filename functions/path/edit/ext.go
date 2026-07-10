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
//
// value: string
// returns: string
//
// example: drop the extension from a file path
// in:  file_path = "/uploads/avatar.png"
// tpl: {{ file_path | ext_trim }}
// out: /uploads/avatar
//
// example: get the bare report name
// in:  source = "Q3-report.pdf"
// tpl: {{ source | ext_trim }}
// out: Q3-report
//
// example: only the last extension is removed
// in:  archive = "backups/archive.tar.gz"
// tpl: {{ archive | ext_trim }}
// out: backups/archive.tar
//
// example: a path with no extension is returned unchanged
// in:  name = "README"
// tpl: {{ name | ext_trim }}
// out: README
func FilenameTrimExt(value any, params []any) (any, error) {
	path, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}
	oldExt := filepath.Ext(path)
	if oldExt == "" {
		return path, nil
	}
	path = path[:len(path)-len(oldExt)]
	return path, nil
}

// FilenamePrependExt inserts an extra extension segment just before the existing
// file extension, so "styles.css" with param "min" becomes "styles.min.css". The
// inserted segment is taken literally and should not include a leading dot, since
// the dots are added for you. When the path has no extension, the segment is
// appended as a new extension ("noext" becomes "noext.min"). Only the final
// extension is considered, so "archive.tar.gz" becomes "archive.tar.min.gz".
//
// value: string
// param:0: string, the extension segment to insert before the current extension, written without a leading dot (e.g. "min", "backup")
// returns: string
//
// example: mark a stylesheet as minified
// in:  file_path = "assets/styles.css"
// tpl: {{ file_path | ext_prepend:'min' }}
// out: assets/styles.min.css
//
// example: tag a script as minified
// in:  source = "app.js"
// tpl: {{ source | ext_prepend:'min' }}
// out: app.min.js
//
// example: only the last extension is split on
// in:  archive = "archive.tar.gz"
// tpl: {{ archive | ext_prepend:'bak' }}
// out: archive.tar.bak.gz
//
// example: a path with no extension gets the segment as a new extension
// in:  name = "backups/data"
// tpl: {{ name | ext_prepend:'bak' }}
// out: backups/data.bak
func FilenamePrependExt(value any, params []any) (any, error) {
	path, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}
	// add another extension before the last dot
	newExt, err := functions.ParamString(params, 0)
	if err != nil {
		return nil, err
	}
	oldExt := filepath.Ext(path)
	if oldExt == "" {
		path += "." + newExt
	} else {
		path = path[:len(path)-len(oldExt)] + "." + newExt + oldExt
	}
	return path, nil
}
