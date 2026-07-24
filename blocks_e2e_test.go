package sintax

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

// A variable whose name merely starts with a tag keyword must render its own
// value. Classifying it as a block tag is the worst kind of failure: the value
// is present in the data, the template looks right, and the output is simply
// missing the field.
func Test_Render_TagPrefixedVariables(t *testing.T) {
	type testCase struct {
		name     string
		template string
		vars     map[string]any
		expected string
	}

	testCases := []testCase{
		{
			name:     "elsewhere",
			template: "{{ elsewhere }}",
			vars:     map[string]any{"elsewhere": "over there"},
			expected: "over there",
		},
		{
			name:     "else_branch",
			template: "value: {{ else_branch }}",
			vars:     map[string]any{"else_branch": "left"},
			expected: "value: left",
		},
		{
			name:     "elsewise",
			template: "{{ elsewise }}",
			vars:     map[string]any{"elsewise": "otherwise"},
			expected: "otherwise",
		},
		{
			name:     "else-prefixed name inside an if body",
			template: "{{ if on }}{{ elsewhere }}{{ endif }}",
			vars:     map[string]any{"on": true, "elsewhere": "inner"},
			expected: "inner",
		},
		{
			name:     "else-prefixed name with a modifier",
			template: "{{ elsewhere | upper }}",
			vars:     map[string]any{"elsewhere": "over there"},
			expected: "OVER THERE",
		},
		{
			name:     "for-prefixed name",
			template: "{{ format }}",
			vars:     map[string]any{"format": "csv"},
			expected: "csv",
		},
		{
			name:     "endif-prefixed name",
			template: "{{ endifx }}",
			vars:     map[string]any{"endifx": "ok"},
			expected: "ok",
		},
		{
			name:     "if-prefixed name",
			template: "{{ iffy }}",
			vars:     map[string]any{"iffy": "maybe"},
			expected: "maybe",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(builtins()).RenderString(tt.template, tt.vars)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

// The genuine else tag keeps working, in both branches and when nested.
func Test_Render_ElseTag(t *testing.T) {
	type testCase struct {
		name     string
		template string
		vars     map[string]any
		expected string
	}

	testCases := []testCase{
		{
			name:     "condition true takes the if branch",
			template: "{{ if active }}yes{{ else }}no{{ endif }}",
			vars:     map[string]any{"active": true},
			expected: "yes",
		},
		{
			name:     "condition false takes the else branch",
			template: "{{ if active }}yes{{ else }}no{{ endif }}",
			vars:     map[string]any{"active": false},
			expected: "no",
		},
		{
			name:     "else branch renders variables",
			template: "{{ if active }}{{ on_label }}{{ else }}{{ off_label }}{{ endif }}",
			vars:     map[string]any{"active": false, "on_label": "ON", "off_label": "OFF"},
			expected: "OFF",
		},
		{
			name:     "nested if/else",
			template: "{{ if a }}{{ if b }}ab{{ else }}a{{ endif }}{{ else }}none{{ endif }}",
			vars:     map[string]any{"a": true, "b": false},
			expected: "a",
		},
		{
			name:     "else inside a loop body",
			template: "{{ for n in nums }}{{ if n }}t{{ else }}f{{ endif }}{{ endfor }}",
			vars:     map[string]any{"nums": []any{true, false, true}},
			expected: "tft",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(builtins()).RenderString(tt.template, tt.vars)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

// There is no ternary form in the template language, so a span that merely looks
// like one is classified by its actual shape. The pipeline case is the one that
// bites in practice, since a quoted modifier argument is free to contain any
// punctuation at all.
func Test_Render_TernaryLikeSpans(t *testing.T) {
	t.Run("quoted argument containing ? and :", func(t *testing.T) {
		got, err := New(builtins()).RenderString(`{{ msg | default:'yes ? no : maybe' }}`, map[string]any{"msg": "present"})
		assert.NoError(t, err)
		assert.Equal(t, "present", got)
	})

	t.Run("quoted argument containing ? and : with a missing variable", func(t *testing.T) {
		got, err := New(builtins()).RenderString(`{{ msg | default:'yes ? no : maybe' }}`, map[string]any{})
		assert.NoError(t, err)
		assert.Equal(t, "yes ? no : maybe", got)
	})

	t.Run("bare ternary is not a token", func(t *testing.T) {
		tokens, err := NewStringParser().Parse(`{{ a ? b : c }}`)
		assert.NoError(t, err)
		assert.Len(t, tokens, 1)
		assert.Equal(t, TextToken, tokens[0].Type())
	})
}

// Unterminated blocks and stray closers are reported with sentinels, so a caller
// can tell a malformed template apart from missing data without reading messages.
func Test_Render_BlockErrors_AreMatchable(t *testing.T) {
	type testCase struct {
		name     string
		template string
		vars     map[string]any
		target   error
	}

	testCases := []testCase{
		{
			name:     "missing endif",
			template: "{{ if active }}yes",
			vars:     map[string]any{"active": true},
			target:   ErrUnterminatedIf,
		},
		{
			name:     "missing endif with an else",
			template: "{{ if active }}yes{{ else }}no",
			vars:     map[string]any{"active": false},
			target:   ErrUnterminatedIf,
		},
		{
			name:     "missing endfor",
			template: "{{ for n in nums }}{{ n }}",
			vars:     map[string]any{"nums": []any{1, 2}},
			target:   ErrUnterminatedFor,
		},
		{
			name:     "stray endif",
			template: "text {{ endif }}",
			vars:     map[string]any{},
			target:   ErrUnexpectedToken,
		},
		{
			name:     "stray else",
			template: "text {{ else }} more",
			vars:     map[string]any{},
			target:   ErrUnexpectedToken,
		},
		{
			name:     "for without an iterable expression",
			template: "{{ for n }}{{ n }}{{ endfor }}",
			vars:     map[string]any{},
			target:   ErrInvalidForExpr,
		},
		{
			name:     "for over a scalar",
			template: "{{ for n in count }}{{ n }}{{ endfor }}",
			vars:     map[string]any{"count": 3},
			target:   ErrNotIterable,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(builtins()).RenderString(tt.template, tt.vars)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tt.target)
		})
	}
}
