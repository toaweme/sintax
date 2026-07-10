package query

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Dirname(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"absolute path", "/var/log/app/server.log", "/var/log/app"},
		{"relative path", "src/handlers/users.go", "src/handlers"},
		{"bare file name has no directory", "file.txt", "."},
		{"trailing slash means path is the directory", "/foo/bar/", "/foo/bar"},
		{"relative trailing slash", "a/b/c/", "a/b/c"},
		{"root stays root", "/", "/"},
		{"empty input", "", "."},
		{"dot-slash prefix", "./x", "."},
		{"hidden dotfile with dir", "/a/.env", "/a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Dirname(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Dirname_RejectsNonString(t *testing.T) {
	for _, v := range []any{nil, 42, 3.14, true, []string{"a"}} {
		_, err := Dirname(v, nil)
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	}
}
