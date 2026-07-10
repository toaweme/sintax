package render

import (
	"errors"
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// Test_Template exercises the modifier's own logic (value coercion and scope
// selection) in isolation, with a stub render callback standing in for the
// engine. Real nested rendering, recursion guarding, and file composition are
// covered by the E2E tests in the root package.
func Test_Template(t *testing.T) {
	parent := map[string]any{"name": "parent"}
	extra := map[string]any{"city": "Vilnius"}

	tests := []struct {
		title     string
		value     any
		params    []any
		wantSrc   string
		wantScope map[string]any
		wantErr   bool
	}{
		{
			title:     "inherits parent vars when no params",
			value:     "Hi {{ name }}",
			params:    nil,
			wantSrc:   "Hi {{ name }}",
			wantScope: parent,
		},
		{
			title:     "renders against only the extra map when given",
			value:     "{{ city }}",
			params:    []any{extra},
			wantSrc:   "{{ city }}",
			wantScope: extra,
		},
		{
			title:     "ignores a nil extra param and inherits",
			value:     "x",
			params:    []any{nil},
			wantSrc:   "x",
			wantScope: parent,
		},
		{
			title:   "errors when the extra param is not a map",
			value:   "x",
			params:  []any{"not-a-map"},
			wantErr: true,
		},
		{
			title:   "errors when the value is not a string",
			value:   123,
			params:  nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			rendered := false
			var gotSrc string
			var gotScope map[string]any
			render := func(src string, vars map[string]any) (any, error) {
				rendered = true
				gotSrc, gotScope = src, vars
				return "out:" + src, nil
			}

			out, err := Template(render, parent, tt.value, tt.params)
			if tt.wantErr {
				assert.Error(t, err)
				assert.True(t, !rendered, "render must not be called when the modifier rejects its input")
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, "out:"+tt.wantSrc, out)
			assert.Equal(t, tt.wantSrc, gotSrc)
			assert.Equal(t, tt.wantScope, gotScope)
		})
	}
}

// Test_Template_NonStringValue confirms the modifier rejects a non-string value
// with the shared ErrInvalidValueType sentinel before it ever calls render, so
// callers can branch on the well-known error rather than a message string.
func Test_Template_NonStringValue(t *testing.T) {
	render := func(string, map[string]any) (any, error) {
		t.Fatalf("render must not be called for a non-string value")
		return nil, nil
	}

	_, err := Template(render, map[string]any{}, 123, nil)
	assert.Error(t, err)
	assert.ErrorIs(t, err, functions.ErrInvalidValueType)
}

// Test_Template_NonMapExtra confirms that an isolated-scope param of the wrong
// type is rejected, since an isolated scope must be a variable map.
func Test_Template_NonMapExtra(t *testing.T) {
	render := func(string, map[string]any) (any, error) {
		t.Fatalf("render must not be called for a non-map extra param")
		return nil, nil
	}

	_, err := Template(render, map[string]any{}, "x", []any{"not-a-map"})
	assert.Error(t, err)
}

// Test_Template_PropagatesRenderError proves the modifier surfaces a failure from
// the nested render rather than swallowing it. The recursion guard reaches the
// modifier this way: at max depth the engine's render callback returns
// ErrMaxDepthExceeded, and the modifier must wrap and propagate it. That real
// sentinel lives in the root package (an import cycle keeps it out of this unit
// test), so the end-to-end guard is asserted in Test_E2E_TemplateModifier_RecursionGuard;
// here a stand-in error stands for any nested-render failure.
func Test_Template_PropagatesRenderError(t *testing.T) {
	depthExceeded := errors.New("max template nesting depth exceeded")
	render := func(string, map[string]any) (any, error) {
		return nil, depthExceeded
	}

	_, err := Template(render, map[string]any{}, "{{ self | template }}", nil)
	assert.Error(t, err)
	assert.ErrorIs(t, err, depthExceeded)
}

// Test_Template_ScopeIsReplacedNotMerged documents that isolated scope is total
// replacement, not a merge: a variable present only in the parent must not leak
// into the isolated render.
func Test_Template_ScopeIsReplacedNotMerged(t *testing.T) {
	parent := map[string]any{"name": "parent", "secret": "leak"}
	extra := map[string]any{"city": "Vilnius"}

	var gotScope map[string]any
	render := func(_ string, vars map[string]any) (any, error) {
		gotScope = vars
		return "ok", nil
	}

	_, err := Template(render, parent, "{{ city }}", []any{extra})
	assert.NoError(t, err)
	assert.Equal(t, extra, gotScope)
	if _, leaked := gotScope["secret"]; leaked {
		t.Fatalf("parent var leaked into isolated scope: %v", gotScope)
	}
}
