package query

import (
	"path/filepath"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFilename is the template name for the Filename modifier.
const ModifierNameFilename functions.ModifierName = "filename"

// ModifierNameFilenameExt is the template name for the FilenameExt modifier.
const ModifierNameFilenameExt functions.ModifierName = "ext"

// ModifierNameFilenameExtDot is the template name for the FilenameExtDot modifier.
const ModifierNameFilenameExtDot functions.ModifierName = "ext_dot"

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
