package trim

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Trim(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"surrounding whitespace", "  Alice  ", nil, "Alice"},
		{"tabs and newlines", "\t\nAlice\n\t", nil, "Alice"},
		{"inner whitespace kept", "  a b  ", nil, "a b"},
		{"no whitespace", "Alice", nil, "Alice"},
		{"empty", "", nil, ""},
		{"whitespace only", "   ", nil, ""},
		{"cutset removes wrapping commas", ",apples,bananas,", []any{","}, "apples,bananas"},
		{"cutset is a set not a fixed string", "xy-hello-yx", []any{"xy"}, "-hello-"},
		{"cutset strips repeated char", "//path//", []any{"/"}, "path"},
		{"cutset absent is no-op", "hello", []any{"/"}, "hello"},
		{"cutset does not trim whitespace", "  hi  ", []any{"/"}, "  hi  "},
		{"empty cutset is no-op", "  hi  ", []any{""}, "  hi  "},
		{"bytes whitespace", []byte("  Alice  "), nil, []byte("Alice")},
		{"bytes cutset", []byte(",a,b,"), []any{","}, []byte("a,b")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Trim(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Trim_Errors(t *testing.T) {
	t.Run("wrong value type", func(t *testing.T) {
		_, err := Trim(42, nil)
		assert.Error(t, err)
	})
	t.Run("nil value", func(t *testing.T) {
		_, err := Trim(nil, nil)
		assert.Error(t, err)
	})
	t.Run("non-string param", func(t *testing.T) {
		_, err := Trim("x", []any{42})
		assert.Error(t, err)
	})
	t.Run("non-string param on bytes", func(t *testing.T) {
		_, err := Trim([]byte("x"), []any{42})
		assert.Error(t, err)
	})
}

func Test_TrimPrefix(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"leading whitespace", "   Welcome aboard.", nil, "Welcome aboard."},
		{"leading tabs and newlines", "\n\t hi", nil, "hi"},
		{"trailing whitespace kept", "  hi  ", nil, "hi  "},
		{"empty", "", nil, ""},
		{"remove prefix", "/api/v1/users", []any{"/"}, "api/v1/users"},
		{"prefix matched once", "//api", []any{"/"}, "/api"},
		{"multichar prefix", "prefix-body", []any{"prefix-"}, "body"},
		{"prefix absent is no-op", "api/v1/users", []any{"/"}, "api/v1/users"},
		{"prefix only in middle is no-op", "a/b/c", []any{"/"}, "a/b/c"},
		{"empty prefix is no-op", "hello", []any{""}, "hello"},
		{"bytes whitespace", []byte("  hi"), nil, []byte("hi")},
		{"bytes prefix", []byte("/api"), []any{"/"}, []byte("api")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := TrimPrefix(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_TrimPrefix_Errors(t *testing.T) {
	t.Run("wrong value type", func(t *testing.T) {
		_, err := TrimPrefix(3.14, nil)
		assert.Error(t, err)
	})
	t.Run("nil value", func(t *testing.T) {
		_, err := TrimPrefix(nil, nil)
		assert.Error(t, err)
	})
	t.Run("non-string param", func(t *testing.T) {
		_, err := TrimPrefix("x", []any{true})
		assert.Error(t, err)
	})
}

func Test_TrimSuffix(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"trailing whitespace", "Welcome aboard.   ", nil, "Welcome aboard."},
		{"trailing tabs and newlines", "hi \t\n", nil, "hi"},
		{"leading whitespace kept", "  hi  ", nil, "  hi"},
		{"empty", "", nil, ""},
		{"remove suffix", "https://example.com/users/", []any{"/"}, "https://example.com/users"},
		{"suffix matched once", "path//", []any{"/"}, "path/"},
		{"drop extension", "report.txt", []any{".txt"}, "report"},
		{"suffix absent is no-op", "report.md", []any{".txt"}, "report.md"},
		{"empty suffix is no-op", "hello", []any{""}, "hello"},
		{"bytes whitespace", []byte("hi  "), nil, []byte("hi")},
		{"bytes suffix", []byte("api/"), []any{"/"}, []byte("api")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := TrimSuffix(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_TrimSuffix_Errors(t *testing.T) {
	t.Run("wrong value type", func(t *testing.T) {
		_, err := TrimSuffix([]int{1}, nil)
		assert.Error(t, err)
	})
	t.Run("nil value", func(t *testing.T) {
		_, err := TrimSuffix(nil, nil)
		assert.Error(t, err)
	})
	t.Run("non-string param", func(t *testing.T) {
		_, err := TrimSuffix("x", []any{1.5})
		assert.Error(t, err)
	})
}
