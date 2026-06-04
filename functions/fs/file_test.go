package fs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_resolveSafePaths(t *testing.T) {
	tests := []struct {
		name      string
		value     any
		safeDirs  []string
		wantRel   string
		wantPaths []string
		wantErr   bool
		wantErrIs error
	}{
		{
			name:      "non-string arg is rejected",
			value:     123,
			safeDirs:  []string{"tpl"},
			wantErr:   true,
			wantErrIs: functions.ErrInvalidValueType,
		},
		{
			name:     "empty safe dirs errors",
			value:    "greeting.txt",
			safeDirs: nil,
			wantRel:  "greeting.txt",
			wantErr:  true,
		},
		{
			name:      "single safe dir resolves",
			value:     "greeting.txt",
			safeDirs:  []string{"tpl"},
			wantRel:   "greeting.txt",
			wantPaths: []string{filepath.Join("tpl", "greeting.txt")},
		},
		{
			name:      "nested relative path resolves",
			value:     "a/b/doc.md",
			safeDirs:  []string{"tpl"},
			wantRel:   "a/b/doc.md",
			wantPaths: []string{filepath.Join("tpl", "a", "b", "doc.md")},
		},
		{
			name:      "candidate per safe dir",
			value:     "doc.md",
			safeDirs:  []string{"a", "b"},
			wantRel:   "doc.md",
			wantPaths: []string{filepath.Join("a", "doc.md"), filepath.Join("b", "doc.md")},
		},
		{
			name:      "parent traversal is dropped",
			value:     "../secret.txt",
			safeDirs:  []string{"safe"},
			wantRel:   "../secret.txt",
			wantErr:   true,
			wantErrIs: os.ErrNotExist,
		},
		{
			name:      "traversal that stays inside is kept",
			value:     "sub/../doc.md",
			safeDirs:  []string{"tpl"},
			wantRel:   "sub/../doc.md",
			wantPaths: []string{filepath.Join("tpl", "doc.md")},
		},
		{
			name:      "escape in one dir is dropped but kept in another",
			value:     "../b/doc.md",
			safeDirs:  []string{"a", filepath.Join("a", "b")},
			wantRel:   "../b/doc.md",
			wantPaths: []string{filepath.Join("a", "b", "doc.md")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rel, paths, err := resolveSafePaths(tt.value, tt.safeDirs)

			assert.Equal(t, tt.wantRel, rel)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.wantErrIs != nil {
					assert.ErrorIs(t, err, tt.wantErrIs)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantPaths, paths)
		})
	}
}

func Test_File_ReadsFromSafeDir(t *testing.T) {
	dir := t.TempDir()
	want := "hello from disk"
	if err := os.WriteFile(filepath.Join(dir, "greeting.txt"), []byte(want), 0o600); err != nil {
		t.Fatalf("failed to write fixture: %v", err)
	}

	out, err := File([]string{dir})("greeting.txt", nil)
	assert.NoError(t, err)
	assert.Equal(t, want, out)
}

func Test_File_ResolvesFromFirstMatchingDir(t *testing.T) {
	missing := t.TempDir()
	present := t.TempDir()
	want := "found in second dir"
	if err := os.WriteFile(filepath.Join(present, "doc.md"), []byte(want), 0o600); err != nil {
		t.Fatalf("failed to write fixture: %v", err)
	}

	out, err := File([]string{missing, present})("doc.md", nil)
	assert.NoError(t, err)
	assert.Equal(t, want, out)
}

func Test_File_MissingFileErrors(t *testing.T) {
	dir := t.TempDir()

	_, err := File([]string{dir})("nope.txt", nil)
	assert.ErrorIs(t, err, os.ErrNotExist)
}

func Test_File_TraversalEscapeIsRejected(t *testing.T) {
	parent := t.TempDir()
	if err := os.WriteFile(filepath.Join(parent, "secret.txt"), []byte("top secret"), 0o600); err != nil {
		t.Fatalf("failed to write fixture: %v", err)
	}
	safe := filepath.Join(parent, "safe")
	if err := os.Mkdir(safe, 0o700); err != nil {
		t.Fatalf("failed to create safe dir: %v", err)
	}

	// the file exists, but only reachable by escaping the safe dir
	_, err := File([]string{safe})("../secret.txt", nil)
	assert.ErrorIs(t, err, os.ErrNotExist)
}

func Test_File_NoSafeDirsErrors(t *testing.T) {
	_, err := File(nil)("greeting.txt", nil)
	assert.Error(t, err)
}
