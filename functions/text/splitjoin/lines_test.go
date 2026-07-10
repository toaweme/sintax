package splitjoin

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Lines(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected []string
	}{
		{"three lines", "Buy milk\nWalk the dog\nPay rent", []string{"Buy milk", "Walk the dog", "Pay rent"}},
		{"single line", "hello", []string{"hello"}},
		{"empty string yields one empty element", "", []string{""}},
		{"trailing newline yields empty tail", "a\nb\n", []string{"a", "b", ""}},
		{"leading newline yields empty head", "\na", []string{"", "a"}},
		{"crlf leaves trailing carriage return", "a\r\nb", []string{"a\r", "b"}},
		{"blank line between", "a\n\nb", []string{"a", "", "b"}},
		{"unicode content", "café\ntea", []string{"café", "tea"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Lines(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Lines_Bytes(t *testing.T) {
	out, err := Lines([]byte("a\nb\nc"), nil)
	assert.NoError(t, err)
	assert.Equal(t, [][]byte{[]byte("a"), []byte("b"), []byte("c")}, out)
}

func Test_Lines_Nil(t *testing.T) {
	out, err := Lines(nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, nil, out)
}

func Test_Lines_Errors(t *testing.T) {
	tests := []struct {
		name  string
		value any
	}{
		{"integer value", 42},
		{"string slice value", []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Lines(tt.value, nil)
			assert.Error(t, err)
		})
	}
}
