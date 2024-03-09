package sintax

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringRenderer_RenderVariable(t *testing.T) {
	type testCase struct {
		name     string
		token    Token
		vars     map[string]any
		expected string
		err      error
	}

	testCases := []testCase{
		{
			name:     "basic variable",
			token:    BaseToken{TokenType: VariableToken, RawValue: "content"},
			vars:     map[string]any{"content": "Hello, World!"},
			expected: "Hello, World!",
		},
		{
			name:     "variable with trim function",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | trim"},
			vars:     map[string]any{"content": " Hello, World! "},
			expected: "Hello, World!",
		},
		{
			name:     "variable with shorten function",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | shorten:5"},
			vars:     map[string]any{"content": "Hello, World!"},
			expected: "Hello",
		},
		{
			name:     "variable with length function",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | length"},
			vars:     map[string]any{"content": "Hello, World!"},
			expected: "13",
		},
		{
			name:     "variable not found",
			token:    BaseToken{TokenType: VariableToken, RawValue: "missing"},
			vars:     map[string]any{"content": "Hello, World!"},
			expected: "",
			err:      errors.New("variable 'missing' not found"),
		},
		{
			name:     "invalid token type",
			token:    BaseToken{TokenType: IfToken, RawValue: "condition"},
			vars:     map[string]any{"content": "Hello, World!"},
			expected: "",
			err:      errors.New("invalid token type: 3: condition"),
		},
		{
			name:     "invalid function",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | invalid"},
			vars:     map[string]any{"content": "Hello, World!"},
			expected: "",
			err:      errors.New("function 'invalid' not found"),
		},
		{
			// trim: "Hello, World!"
			// shorten: "Hello"
			// length: "5"
			name:     "multiple functions",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | trim | shorten:5 | length"},
			vars:     map[string]any{"content": " Hello, World! "},
			expected: "5",
		},
	}

	r := NewStringRenderer(map[string]GlobalModifier{
		"trim":    trim,
		"shorten": shorten,
		"length":  length,
	})

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := r.renderVariable(tt.token, tt.vars)
			if tt.err != nil {
				assert.Equal(t, err.Error(), tt.err.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

var trim = func(s any, _ []any) (any, error) {
	switch v := s.(type) {
	case string:
		return strings.TrimSpace(v), nil
		// default:
		// 	return "", fmt.Errorf("trim requires a text argument")
	}
	return s, nil
}

var shorten = func(s any, args []any) (any, error) {
	str, ok := s.(string)
	if !ok {
		return "", fmt.Errorf("shorten requires a text argument")
	}

	if len(args) != 1 {
		return "", fmt.Errorf("shorten requires 1 argument")
	}
	length, err := strconv.Atoi(fmt.Sprint(args[0]))
	if err != nil {
		return "", fmt.Errorf("shorten requires a numeric argument")
	}

	if len(str) > length {
		return str[:length], nil
	}

	return str, nil
}

var length = func(s any, _ []any) (any, error) {
	str, ok := s.(string)
	if !ok {
		return "", fmt.Errorf("length requires a text argument")
	}

	return strconv.Itoa(len(str)), nil
}
