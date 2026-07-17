package query

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Filename(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"absolute path", "/var/log/app/server.log", "server.log"},
		{"relative path", "downloads/2024/Q3-report.pdf", "Q3-report.pdf"},
		{"bare name unchanged", "file.txt", "file.txt"},
		{"trailing slash returns last folder", "/foo/bar/", "bar"},
		{"hidden dotfile", "/a/.env", ".env"},
		{"no extension", "/usr/local/bin/go", "go"},
		{"root", "/", "/"},
		{"empty input", "", "."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Filename(tt.value)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_FilenameExt(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"simple extension", "/uploads/avatar.png", "png"},
		{"document format", "Q3-report.pdf", "pdf"},
		{"only last of double extension", "backup.tar.gz", "gz"},
		{"uppercase preserved", "photo.JPG", "JPG"},
		{"no dot has no extension", "README", ""},
		{"trailing slash has no extension", "/foo/bar/", ""},
		{"hidden dotfile is all extension", ".gitignore", "gitignore"},
		{"empty input", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := FilenameExt(tt.value)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_FilenameExtDot(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"simple extension", "/uploads/avatar.png", ".png"},
		{"document format", "Q3-report.pdf", ".pdf"},
		{"only last of double extension", "backup.tar.gz", ".gz"},
		{"uppercase preserved", "photo.JPG", ".JPG"},
		{"no dot has no extension", "README", ""},
		{"trailing slash has no extension", "/foo/bar/", ""},
		{"hidden dotfile is all extension", ".gitignore", ".gitignore"},
		{"empty input", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := FilenameExtDot(tt.value)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Filename_RejectsNonString(t *testing.T) {
	fns := map[string]functions.GlobalModifier{
		string(ModifierNameFilename):       filenameModifier,
		string(ModifierNameFilenameExt):    extModifier,
		string(ModifierNameFilenameExtDot): extDotModifier,
	}
	for name, fn := range fns {
		t.Run(name, func(t *testing.T) {
			for _, v := range []any{nil, 42, 3.14, true, []string{"a"}} {
				_, err := fn(v, nil)
				assert.ErrorIs(t, err, functions.ErrInvalidValueType)
			}
		})
	}
}
