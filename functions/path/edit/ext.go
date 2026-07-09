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

// FilenameTrimExt returns the file path without its extension.
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

// FilenamePrependExt inserts an additional extension before the existing file extension.
//
// value: string
// param:0: string
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
