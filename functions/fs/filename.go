package fs

import (
	"path/filepath"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFilename is the template name for the Filename modifier.
const ModifierNameFilename functions.ModifierName = "filename"

// ModifierNameFilenameTrimExt is the template name for the FilenameTrimExt modifier.
const ModifierNameFilenameTrimExt functions.ModifierName = "ext_trim"

// ModifierNameFilenameExt is the template name for the FilenameExt modifier.
const ModifierNameFilenameExt functions.ModifierName = "ext"

// ModifierNameFilenameExtDot is the template name for the FilenameExtDot modifier.
const ModifierNameFilenameExtDot functions.ModifierName = "ext_dot"

// ModifierNameFilenamePrependExt is the template name for the FilenamePrependExt modifier.
const ModifierNameFilenamePrependExt functions.ModifierName = "ext_prepend"

// Filename returns the base file name from a path, including the extension.
//
// value: string
// returns: string
//
// example: pull the file name out of a full path
// in:  file_path = "/var/log/app/server.log"
// tpl: {{ file_path | filename }}
// out: server.log
//
// example: shorten a download path to its name
// in:  source = "downloads/2024/Q3-report.pdf"
// tpl: {{ source | filename }}
// out: Q3-report.pdf
func Filename(value any, params []any) (any, error) {
	path, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}

	return filepath.Base(path), nil
}

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

// FilenameExt returns the file extension without the leading dot.
//
// value: string
// returns: string
//
// example: read the type of an upload
// in:  file_path = "/uploads/avatar.png"
// tpl: {{ file_path | ext }}
// out: png
//
// example: detect a document format
// in:  source = "Q3-report.pdf"
// tpl: {{ source | ext }}
// out: pdf
func FilenameExt(value any, params []any) (any, error) {
	path, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}
	oldExt := filepath.Ext(path)
	if oldExt == "" {
		return "", nil
	}
	return oldExt[1:], nil
}

// FilenameExtDot returns the file extension including the leading dot.
//
// value: string
// returns: string
//
// example: read an extension with the dot included
// in:  file_path = "/uploads/avatar.png"
// tpl: {{ file_path | ext_dot }}
// out: .png
//
// example: build a filename suffix
// in:  source = "Q3-report.pdf"
// tpl: {{ source | ext_dot }}
// out: .pdf
func FilenameExtDot(value any, params []any) (any, error) {
	path, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}
	oldExt := filepath.Ext(path)
	if oldExt == "" {
		return "", nil
	}
	return oldExt, nil
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
