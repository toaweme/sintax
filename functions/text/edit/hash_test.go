package edit

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Hash(t *testing.T) {
	hash := hashModifier
	tests := []struct {
		name     string
		value    any
		params   []any
		expected string
	}{
		{"sha256", "hello", []any{"sha256"}, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"sha512", "hello", []any{"sha512"}, "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"},
		{"sha1", "hello", []any{"sha1"}, "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{"md5", "hello", []any{"md5"}, "5d41402abc4b2a76b9719d911017c592"},
		{"empty value", "", []any{"sha256"}, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"number value", 42, []any{"sha256"}, "73475cb40a568e8da8a045ced110137e159f890ac4da883b6b17dc651b3a8049"},
		{"bool value", true, []any{"sha1"}, "5ffe533b830f08a0326348a9160afafc8ada44db"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := hash(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Hash_NoNormalization locks in that the bytes are hashed as they arrive.
// " Hello " keeps its surrounding spaces and its capital H, so it must not
// produce the digest of "hello".
func Test_Hash_NoNormalization(t *testing.T) {
	hash := hashModifier
	out, err := hash(" Hello ", []any{"sha256"})
	assert.NoError(t, err)
	assert.Equal(t, "4647f3f336c405d1d64df86a17a78929fa4f53678a518876e2215278b5fa79c4", out)
}

func Test_Hash_Errors(t *testing.T) {
	hash := hashModifier
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{"unknown algorithm", "hello", []any{"sha3"}},
		{"uppercase algorithm", "hello", []any{"SHA256"}},
		{"empty algorithm", "hello", []any{""}},
		{"no params", "hello", nil},
		{"too many params", "hello", []any{"sha256", "sha512"}},
		{"non-string param", "hello", []any{256}},
		{"composite value", []any{1, 2}, []any{"sha256"}},
		{"map value", map[string]any{"a": 1}, []any{"sha256"}},
		{"nil value", nil, []any{"sha256"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := hash(tt.value, tt.params)
			assert.Error(t, err)
		})
	}
}
