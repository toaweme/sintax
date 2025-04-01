package sintax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parser_Parse(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected []Token
		err      error
	}

	testCases := []testCase{
		{
			name:  "basic conditional",
			input: "{{ if condition }} Hello, World! {{ /if }}",
			expected: []Token{
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: TextToken, RawValue: " Hello, World! "},
				BaseToken{TokenType: IfEndToken, RawValue: ""},
			},
		},
		{
			name:  "basic conditional without spaces",
			input: "{{if condition}} Hello, World! {{/if}}",
			expected: []Token{
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: TextToken, RawValue: " Hello, World! "},
				BaseToken{TokenType: IfEndToken, RawValue: ""},
			},
		},
		{
			name:  "conditional with variable",
			input: "{{ if condition }}{{ content }}{{ /if }}",
			expected: []Token{
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: VariableToken, RawValue: "content", Var: "content"},
				BaseToken{TokenType: IfEndToken, RawValue: ""},
			},
		},
		{
			name:  "conditional with variable with filters",
			input: "{{ if condition }}{{ content | xss | summary:255,300 }}{{ /if }}",
			expected: []Token{
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: FilteredVariableToken, RawValue: "content | xss | summary:255,300", Var: "content"},
				BaseToken{TokenType: IfEndToken, RawValue: ""},
			},
		},
		{
			name:  "wrapping text with conditional with variable with filters",
			input: "something cool {{ if condition }} beep {{ content | xss | summary:255,300 }}{{ /if }} cool ending ",
			expected: []Token{
				BaseToken{TokenType: TextToken, RawValue: "something cool "},
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: TextToken, RawValue: " beep "},
				BaseToken{TokenType: FilteredVariableToken, RawValue: "content | xss | summary:255,300", Var: "content"},
				BaseToken{TokenType: IfEndToken, RawValue: ""},
				BaseToken{TokenType: TextToken, RawValue: " cool ending "},
			},
		},
	}

	p := NewStringParser()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := p.Parse(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Parser_ParseVariable(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected []Token
		err      error
	}

	testCases := []testCase{
		{
			name:  "basic variable",
			input: "{{ content }}",
			expected: []Token{
				BaseToken{TokenType: VariableToken, RawValue: "content", Var: "content"},
			},
		},
		{
			name:  "variable with filters",
			input: "{{ vars.content | xss | summary:255,300 }}",
			expected: []Token{
				BaseToken{TokenType: FilteredVariableToken, RawValue: "vars.content | xss | summary:255,300", Var: "vars.content"},
			},
		},
		{
			name:  "incorrect syntax variable without double curly braces",
			input: "{ vars.content | xss | summary:255,300 }",
			expected: []Token{
				BaseToken{TokenType: TextToken, RawValue: "{ vars.content | xss | summary:255,300 }", Var: ""},
			},
		},
	}

	p := NewStringParser()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := p.Parse(tt.input)
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
				return
			}
			assert.NoError(t, tt.err, err)
			assert.Equal(t, tt.expected, tokens)
		})
	}
}

func Benchmark_Parser_Parse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewStringParser()
		p.Parse("something cool {{ if condition }} beep {{ content | xss | summary:255,300 }}{{ /if }} cool ending ")
	}
}
