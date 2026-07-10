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

// Filename returns the base file name from a path, that is the final path
// element including its extension. A trailing slash is ignored, so the last
// real segment is returned. The input must be a string.
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
//
// example: a trailing slash is ignored and the last folder name is returned
// in:  dir = "/foo/bar/"
// tpl: {{ dir | filename }}
// out: bar
func Filename(value any, params []any) (any, error) {
	path, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}

	return filepath.Base(path), nil
}

// FilenameExt returns the file extension without the leading dot, for example
// "png" for "avatar.png". Use ext when you want the bare type, and ext_dot when
// you want the dot kept (".png"). The extension is the part after the final dot,
// so a name like "archive.tar.gz" yields only "gz". A path with no dot has no
// extension and yields an empty string. A hidden dotfile such as ".gitignore"
// is treated as all extension, since its only dot is the leading one, so ext
// returns "gitignore". The input must be a string.
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
//
// example: only the last extension of a double extension is returned
// in:  archive = "backup.tar.gz"
// tpl: {{ archive | ext }}
// out: gz
//
// example: a name with no dot has no extension
// in:  name = "README"
// tpl: {{ name | ext }}
// out:
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

// FilenameExtDot returns the file extension including the leading dot, for
// example ".png" for "avatar.png". It is the counterpart to ext, which drops
// the dot ("png"). Reach for ext_dot when you are rebuilding a file name and
// need the separator kept. As with ext, the extension is the part from the
// final dot onward, so "archive.tar.gz" yields ".gz", a name with no dot yields
// an empty string, and a hidden dotfile such as ".gitignore" is treated as all
// extension and yields ".gitignore". The input must be a string.
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
//
// example: only the last extension of a double extension is returned
// in:  archive = "backup.tar.gz"
// tpl: {{ archive | ext_dot }}
// out: .gz
//
// example: a name with no dot has no extension
// in:  name = "README"
// tpl: {{ name | ext_dot }}
// out:
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
