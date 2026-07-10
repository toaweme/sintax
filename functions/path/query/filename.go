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

// Filename returns the base file name from a path, the final path element
// including its extension. A trailing slash is ignored, so the last real segment
// is returned.
func Filename(s string) (string, error) {
	return filepath.Base(s), nil
}

// FilenameExt returns the file extension without the leading dot, for example
// "png" for "avatar.png". Only the part after the final dot is returned, so
// "archive.tar.gz" yields "gz" and a name with no dot yields "".
func FilenameExt(s string) (string, error) {
	ext := filepath.Ext(s)
	if ext == "" {
		return "", nil
	}
	return ext[1:], nil
}

// FilenameExtDot returns the file extension including the leading dot, for
// example ".png" for "avatar.png". It is the counterpart to FilenameExt for
// callers rebuilding a name that need the separator kept, so a name with no dot
// yields "".
func FilenameExtDot(s string) (string, error) {
	return filepath.Ext(s), nil
}
