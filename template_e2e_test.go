package sintax

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_E2E_TemplateModifier_InheritScope(t *testing.T) {
	s := New(builtins())

	out, err := s.Render(`{{ tpl | template }}`, map[string]any{
		"tpl":  "Hi {{ name }}",
		"name": "Bob",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Hi Bob", out)
}

func Test_E2E_TemplateModifier_FileComposed(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "p.tpl"), []byte("Hello {{ who }}"), 0o600); err != nil {
		t.Fatalf("failed to write fixture: %v", err)
	}

	s := New(builtins(dir))

	out, err := s.Render(`{{ "p.tpl" | file | template }}`, map[string]any{"who": "World"})
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", out)
}

func Test_E2E_TemplateModifier_IsolatedScope(t *testing.T) {
	s := New(builtins())

	out, err := s.Render(`{{ tpl | template:extra }}`, map[string]any{
		"tpl":   "{{ name | default:'?' }}/{{ city }}",
		"name":  "parent",
		"extra": map[string]any{"city": "Vilnius"},
	})
	assert.NoError(t, err)
	// parent's "name" is NOT visible in the isolated scope, "city" from extra is
	assert.Equal(t, "?/Vilnius", out)
}

func Test_E2E_TemplateModifier_RecursionGuard(t *testing.T) {
	s := New(builtins())

	_, err := s.Render(`{{ self | template }}`, map[string]any{
		"self": "{{ self | template }}",
	})
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrMaxDepthExceeded)
}

func Test_E2E_TemplateModifier_NonMapExtra(t *testing.T) {
	s := New(builtins())

	_, err := s.Render(`{{ tpl | template:extra }}`, map[string]any{
		"tpl":   "x",
		"extra": "not-a-map",
	})
	assert.Error(t, err)
}

func Test_E2E_TemplateModifier_MultiLevelNesting(t *testing.T) {
	s := New(builtins())

	out, err := s.Render(`{{ outer | template }}`, map[string]any{
		"outer": "[{{ inner | template }}]",
		"inner": "{{ leaf }}",
		"leaf":  "deep",
	})
	assert.NoError(t, err)
	assert.Equal(t, "[deep]", out)
}

func Test_E2E_TemplateModifier_NestedRenderError(t *testing.T) {
	s := New(builtins())

	// the nested template references an undefined variable, so the failure must
	// propagate out of the modifier rather than be swallowed.
	_, err := s.Render(`{{ tpl | template }}`, map[string]any{
		"tpl": "{{ missing }}",
	})
	assert.Error(t, err)
}
