package edit

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_FilenameTrimExt(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected any
	}{
		{"simple extension", "avatar.png", "avatar"},
		{"path with directories", "/uploads/avatar.png", "/uploads/avatar"},
		{"report name", "Q3-report.pdf", "Q3-report"},
		{"only last extension removed", "archive.tar.gz", "archive.tar"},
		{"many dots", "a.b.c.d", "a.b.c"},
		{"uppercase extension", "photo.JPG", "photo"},
		{"no extension unchanged", "README", "README"},
		{"no extension with dirs", "/etc/hosts", "/etc/hosts"},
		{"dot only in directory segment", "config.d/app", "config.d/app"},
		{"trailing slash directory", "/path/to/dir/", "/path/to/dir/"},
		{"trailing dot loses the dot", "file.", "file"},
		{"dotfile is all extension", ".gitignore", ""},
		{"empty input", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := FilenameTrimExt(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_FilenameTrimExt_NonStringValue(t *testing.T) {
	for _, v := range []any{42, 3.14, true, nil, []int{1}} {
		_, err := FilenameTrimExt(v, nil)
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	}
}

func Test_FilenamePrependExt(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		param    string
		expected any
	}{
		{"stylesheet minified", "assets/styles.css", "min", "assets/styles.min.css"},
		{"script minified", "app.js", "min", "app.min.js"},
		{"image with dirs", "/uploads/avatar.png", "thumb", "/uploads/avatar.thumb.png"},
		{"only last extension split", "archive.tar.gz", "bak", "archive.tar.bak.gz"},
		{"many dots", "a.b.c.d", "min", "a.b.c.min.d"},
		{"uppercase extension preserved", "photo.JPG", "min", "photo.min.JPG"},
		{"no extension gets new one", "backups/data", "bak", "backups/data.bak"},
		{"bare name no extension", "noext", "min", "noext.min"},
		{"dot only in directory segment", "config.d/app", "min", "config.d/app.min"},
		{"trailing slash directory", "/path/to/dir/", "min", "/path/to/dir/.min"},
		{"trailing dot", "file.", "min", "file.min."},
		{"dotfile is all extension", ".gitignore", "min", ".min.gitignore"},
		{"empty input", "", "min", ".min"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := FilenamePrependExt(tt.value, []any{tt.param})
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_FilenamePrependExt_MissingParam(t *testing.T) {
	_, err := FilenamePrependExt("app.js", nil)
	assert.ErrorIs(t, err, functions.ErrMissingParam)

	_, err = FilenamePrependExt("app.js", []any{})
	assert.ErrorIs(t, err, functions.ErrMissingParam)
}

func Test_FilenamePrependExt_NonStringParam(t *testing.T) {
	_, err := FilenamePrependExt("app.js", []any{42})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)
}

func Test_FilenamePrependExt_NonStringValue(t *testing.T) {
	for _, v := range []any{42, 3.14, true, nil, []int{1}} {
		_, err := FilenamePrependExt(v, []any{"min"})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	}
}
