package edit

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Replace(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected string
	}{
		{"swap a word", "Hello world", []any{"world", "everyone"}, "Hello everyone"},
		{"redact a phrase", "The password is hunter2", []any{"hunter2", "******"}, "The password is ******"},
		{"replaces all occurrences", "a-b-c", []any{"-", "_"}, "a_b_c"},
		{"no match unchanged", "hello", []any{"z", "Z"}, "hello"},
		{"delete substring", "foobar", []any{"o", ""}, "fbar"},
		{"empty value", "", []any{"a", "b"}, ""},
		{"unicode substitution", "café time", []any{"é", "e"}, "cafe time"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Replace(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Replace_Errors(t *testing.T) {
	t.Run("too few params", func(t *testing.T) {
		_, err := Replace("x", []any{"only-one"})
		assert.Error(t, err)
	})
	t.Run("no params", func(t *testing.T) {
		_, err := Replace("x", nil)
		assert.Error(t, err)
	})
	t.Run("non-string value", func(t *testing.T) {
		_, err := Replace(42, []any{"a", "b"})
		assert.Error(t, err)
	})
	t.Run("non-string old param", func(t *testing.T) {
		_, err := Replace("x", []any{1, "b"})
		assert.Error(t, err)
	})
	t.Run("non-string new param", func(t *testing.T) {
		_, err := Replace("x", []any{"a", 2})
		assert.Error(t, err)
	})
}

func Test_ReplacePattern(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected string
	}{
		{"collapse whitespace", "hello    world", []any{`\s+`, " "}, "hello world"},
		{"delete non-slug chars", "hello-world!@#", []any{`[^a-z0-9\-]`, ""}, "hello-world"},
		{"reorder capture groups", "Doe, Jane", []any{`(\w+), (\w+)`, "$2 $1"}, "Jane Doe"},
		{"strip digits", "a1b2c3", []any{`\d`, ""}, "abc"},
		{"anchored match", "  trim me", []any{`^\s+`, ""}, "trim me"},
		{"no match unchanged", "plain", []any{`\d+`, "#"}, "plain"},
		{"empty value", "", []any{`\d`, "x"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := ReplacePattern(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_ReplacePattern_Errors(t *testing.T) {
	t.Run("too few params", func(t *testing.T) {
		_, err := ReplacePattern("x", []any{`\d`})
		assert.Error(t, err)
	})
	t.Run("non-string value", func(t *testing.T) {
		_, err := ReplacePattern(42, []any{`\d`, ""})
		assert.Error(t, err)
	})
	t.Run("non-string pattern param", func(t *testing.T) {
		_, err := ReplacePattern("x", []any{1, ""})
		assert.Error(t, err)
	})
	t.Run("non-string replacement param", func(t *testing.T) {
		_, err := ReplacePattern("x", []any{`\d`, 2})
		assert.Error(t, err)
	})
	t.Run("invalid regex", func(t *testing.T) {
		_, err := ReplacePattern("x", []any{`[unterminated`, ""})
		assert.Error(t, err)
	})
}

func Test_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"name", "Alice", "ecilA"},
		{"alphanumeric", "ABC123", "321CBA"},
		{"empty", "", ""},
		{"single char", "x", "x"},
		{"palindrome", "level", "level"},
		{"multi-byte stays intact", "café", "éfac"},
		{"cjk by rune", "日本語", "語本日"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Reverse(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Reverse_Errors(t *testing.T) {
	t.Run("non-string value", func(t *testing.T) {
		_, err := Reverse(42, nil)
		assert.Error(t, err)
	})
	t.Run("nil value", func(t *testing.T) {
		_, err := Reverse(nil, nil)
		assert.Error(t, err)
	})
}
