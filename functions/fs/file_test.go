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
		path      string
		safeDirs  []string
		wantPaths []string
		wantErr   bool
		wantErrIs error
	}{
		{
			name:     "empty safe dirs errors",
			path:     "greeting.txt",
			safeDirs: nil,
			wantErr:  true,
		},
		{
			name:      "single safe dir resolves",
			path:      "greeting.txt",
			safeDirs:  []string{"tpl"},
			wantPaths: []string{filepath.Join("tpl", "greeting.txt")},
		},
		{
			name:      "nested relative path resolves",
			path:      "a/b/doc.md",
			safeDirs:  []string{"tpl"},
			wantPaths: []string{filepath.Join("tpl", "a", "b", "doc.md")},
		},
		{
			name:      "candidate per safe dir",
			path:      "doc.md",
			safeDirs:  []string{"a", "b"},
			wantPaths: []string{filepath.Join("a", "doc.md"), filepath.Join("b", "doc.md")},
		},
		{
			name:      "parent traversal is dropped",
			path:      "../secret.txt",
			safeDirs:  []string{"safe"},
			wantErr:   true,
			wantErrIs: os.ErrNotExist,
		},
		{
			name:      "traversal that stays inside is kept",
			path:      "sub/../doc.md",
			safeDirs:  []string{"tpl"},
			wantPaths: []string{filepath.Join("tpl", "doc.md")},
		},
		{
			name:      "escape in one dir is dropped but kept in another",
			path:      "../b/doc.md",
			safeDirs:  []string{"a", filepath.Join("a", "b")},
			wantPaths: []string{filepath.Join("a", "b", "doc.md")},
		},
		{
			// an absolute-looking path is joined onto the safe dir, never the
			// real filesystem root, so it stays sandboxed
			name:      "absolute path is contained inside the safe dir",
			path:      "/etc/passwd",
			safeDirs:  []string{"tpl"},
			wantPaths: []string{filepath.Join("tpl", "etc", "passwd")},
		},
		{
			// a deep traversal that only partway escapes is still dropped
			name:      "deep parent traversal is dropped",
			path:      "../../../../etc/passwd",
			safeDirs:  []string{"tpl"},
			wantErr:   true,
			wantErrIs: os.ErrNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paths, err := resolveSafePaths(tt.path, tt.safeDirs)

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

	out, err := File([]string{dir})("greeting.txt")
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

	out, err := File([]string{missing, present})("doc.md")
	assert.NoError(t, err)
	assert.Equal(t, want, out)
}

func Test_File_ReadsNestedPath(t *testing.T) {
	dir := t.TempDir()
	nested := filepath.Join(dir, "emails")
	if err := os.Mkdir(nested, 0o700); err != nil {
		t.Fatalf("failed to create nested dir: %v", err)
	}
	want := "welcome aboard"
	if err := os.WriteFile(filepath.Join(nested, "welcome.txt"), []byte(want), 0o600); err != nil {
		t.Fatalf("failed to write fixture: %v", err)
	}

	out, err := File([]string{dir})("emails/welcome.txt")
	assert.NoError(t, err)
	assert.Equal(t, want, out)
}

func Test_File_AbsolutePathIsSandboxed(t *testing.T) {
	dir := t.TempDir()
	// "/etc/passwd" must resolve under the safe dir, not the real root, so
	// reading it fails with not-exist rather than leaking the host file
	_, err := File([]string{dir})("/etc/passwd")
	assert.ErrorIs(t, err, os.ErrNotExist)
}

func Test_File_MissingFileErrors(t *testing.T) {
	dir := t.TempDir()

	_, err := File([]string{dir})("nope.txt")
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
	_, err := File([]string{safe})("../secret.txt")
	assert.ErrorIs(t, err, os.ErrNotExist)
}

func Test_File_NoSafeDirsErrors(t *testing.T) {
	_, err := File(nil)("greeting.txt")
	assert.Error(t, err)
}

// Test_File_Modifier_ReadsThroughWrap exercises the registered modifier so the
// Wrap adapter's value coercion is covered alongside the typed body.
func Test_File_Modifier_ReadsThroughWrap(t *testing.T) {
	dir := t.TempDir()
	want := "hello from disk"
	if err := os.WriteFile(filepath.Join(dir, "greeting.txt"), []byte(want), 0o600); err != nil {
		t.Fatalf("failed to write fixture: %v", err)
	}

	file := Modifiers([]string{dir})[string(ModifierNameFile)]
	out, err := file("greeting.txt", nil)
	assert.NoError(t, err)
	assert.Equal(t, want, out)
}

func Test_File_Modifier_NonStringValueIsRejected(t *testing.T) {
	dir := t.TempDir()

	file := Modifiers([]string{dir})[string(ModifierNameFile)]
	_, err := file(42, nil)
	assert.ErrorIs(t, err, functions.ErrInvalidValueType)
}

func Test_File_Modifier_ParamsAreRejected(t *testing.T) {
	dir := t.TempDir()

	// the file modifier takes zero params, so Wrap rejects any argument
	file := Modifiers([]string{dir})[string(ModifierNameFile)]
	_, err := file("greeting.txt", []any{"extra"})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)
}
