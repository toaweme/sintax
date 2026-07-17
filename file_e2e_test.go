package sintax

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_E2E_FileModifier(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "greeting.txt"), []byte("Hello there"), 0o600); err != nil {
		t.Fatalf("failed to write fixture: %v", err)
	}

	s := New(builtins(dir))

	out, err := s.Render(`{{ "greeting.txt" | file }}`, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Hello there", out)
}

func Test_E2E_FileModifier_VariablePath(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "doc.md"), []byte("# Title"), 0o600); err != nil {
		t.Fatalf("failed to write fixture: %v", err)
	}

	s := New(builtins(dir))

	out, err := s.Render(`{{ path | file }}`, map[string]any{"path": "doc.md"})
	assert.NoError(t, err)
	assert.Equal(t, "# Title", out)
}

func Test_E2E_LiteralHead(t *testing.T) {
	s := New(builtins())

	out, err := s.Render(`{{ "hello" | upper }}`, nil)
	assert.NoError(t, err)
	assert.Equal(t, "HELLO", out)
}
