package functions

import (
	"errors"
	"testing"
)

func Test_AsText(t *testing.T) {
	// a string-only modifier: upper-cases, rejects a non-string with the sentinel.
	upper := Wrap(func(s string) (string, error) {
		out := make([]byte, len(s))
		for i := range s {
			c := s[i]
			if c >= 'a' && c <= 'z' {
				c -= 32
			}
			out[i] = c
		}
		return string(out), nil
	})
	text := AsText(upper)

	tests := []struct {
		name    string
		value   any
		want    any
		wantErr error
	}{
		{"string passes through", "hi", "HI", nil},
		{"int becomes its digits", 42, "42", nil},
		{"float stringifies", 3.5, "3.5", nil},
		{"bool stringifies", true, "TRUE", nil},
		{"nil is left untouched and rejected", nil, nil, ErrInvalidValueType},
		{"slice is left untouched and rejected", []any{1, 2}, nil, ErrInvalidValueType},
		{"map is left untouched and rejected", map[string]any{"a": 1}, nil, ErrInvalidValueType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := text(tt.value, nil)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected %v, got %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if out != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, out)
			}
		})
	}
}

// Test_AsText_LeavesParamsAndArity confirms AsText only rewrites the value: it
// does not touch params, so the wrapped modifier's arity checks still fire.
func Test_AsText_LeavesParamsAndArity(t *testing.T) {
	mod := AsText(Wrap(func(s string) (string, error) { return s, nil }))
	if _, err := mod(42, []any{"extra"}); !errors.Is(err, ErrInvalidParamType) {
		t.Fatalf("expected arity rejection, got %v", err)
	}
}
