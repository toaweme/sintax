package sintax

import (
	"testing"

	"github.com/toaweme/sintax/assert"
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
			input: "{{ if condition }} Hello, World! {{ endif }}",
			expected: []Token{
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: TextToken, RawValue: " Hello, World! "},
				BaseToken{TokenType: IfEndToken, RawValue: ""},
			},
		},
		{
			name:  "basic conditional without spaces",
			input: "{{if condition}} Hello, World! {{endif}}",
			expected: []Token{
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: TextToken, RawValue: " Hello, World! "},
				BaseToken{TokenType: IfEndToken, RawValue: ""},
			},
		},
		{
			name:  "conditional with variable",
			input: "{{ if condition }}{{ content }}{{ endif }}",
			expected: []Token{
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: VariableToken, RawValue: "content", Var: "content"},
				BaseToken{TokenType: IfEndToken, RawValue: ""},
			},
		},
		{
			name:  "conditional with variable with filters",
			input: "{{ if condition }}{{ content | xss | summary:255,300 }}{{ endif }}",
			expected: []Token{
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: FilteredVariableToken, RawValue: "content | xss | summary:255,300", Var: "content", parsedVar: "content", parsedFuncs: []Func{{Name: "xss", Args: []Arg{}}, {Name: "summary", Args: []Arg{{Value: 255}, {Value: 300}}}}},
				BaseToken{TokenType: IfEndToken, RawValue: ""},
			},
		},
		{
			name:  "wrapping text with conditional with variable with filters",
			input: "something cool {{ if condition }} beep {{ content | xss | summary:255,300 }}{{ endif }} cool ending ",
			expected: []Token{
				BaseToken{TokenType: TextToken, RawValue: "something cool "},
				BaseToken{TokenType: IfToken, RawValue: "condition"},
				BaseToken{TokenType: TextToken, RawValue: " beep "},
				BaseToken{TokenType: FilteredVariableToken, RawValue: "content | xss | summary:255,300", Var: "content", parsedVar: "content", parsedFuncs: []Func{{Name: "xss", Args: []Arg{}}, {Name: "summary", Args: []Arg{{Value: 255}, {Value: 300}}}}},
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
				BaseToken{TokenType: FilteredVariableToken, RawValue: "vars.content | xss | summary:255,300", Var: "vars.content", parsedVar: "vars.content", parsedFuncs: []Func{{Name: "xss", Args: []Arg{}}, {Name: "summary", Args: []Arg{{Value: 255}, {Value: 300}}}}},
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
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, tokens)
		})
	}
}

func Benchmark_Parser_Parse(b *testing.B) {
	const tmpl = "something cool {{ if condition }} beep {{ content | xss | summary:255,300 }}{{ endif }} cool ending "
	b.SetBytes(int64(len(tmpl)))
	for range b.N {
		p := NewStringParser()
		p.Parse(tmpl)
	}
}

// Control tags are recognized on a word boundary, so this pins both the tags
// themselves and the ordinary names that merely start with one. A name like
// `elsewhere` must stay a variable, since a token classified as a block tag
// never resolves and leaves nothing in the output to explain the gap.
func Test_DetectTokenType_Classification(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected TokenType
	}

	testCases := []testCase{
		{name: "if tag", input: "if active", expected: IfToken},
		{name: "if tag with a dotted condition", input: "if user.active", expected: IfToken},
		{name: "bare if", input: "if", expected: IfToken},
		{name: "endif tag", input: "endif", expected: IfEndToken},
		{name: "else tag", input: "else", expected: ElseToken},
		{name: "for tag", input: "for item in items", expected: ForToken},
		{name: "endfor tag", input: "endfor", expected: ForEndToken},

		{name: "variable starting with else", input: "elsewhere", expected: VariableToken},
		{name: "variable starting with else and underscore", input: "else_branch", expected: VariableToken},
		{name: "variable elsewise", input: "elsewise", expected: VariableToken},
		{name: "variable starting with endif", input: "endifx", expected: VariableToken},
		{name: "variable starting with endfor", input: "endforall", expected: VariableToken},
		{name: "variable starting with for", input: "format", expected: VariableToken},
		{name: "variable forward", input: "forward", expected: VariableToken},
		{name: "variable starting with if", input: "iffy", expected: VariableToken},
		{name: "bare for", input: "for", expected: VariableToken},
		{name: "bare else nested in a name", input: "config.elsewhere", expected: VariableToken},

		{name: "filtered else-prefixed variable", input: "elsewhere | upper", expected: FilteredVariableToken},
		{name: "filtered variable", input: "name | upper", expected: FilteredVariableToken},
		// a quoted argument holding " ? " and " : " is an ordinary pipeline, the
		// engine has no ternary form that could claim it.
		{name: "filtered variable with ternary-looking literal", input: `msg | default:'yes ? no : maybe'`, expected: FilteredVariableToken},

		{name: "unrecognized expression", input: "1 + 2", expected: UndefinedToken},
		{name: "else with a trailing condition", input: "else if active", expected: UndefinedToken},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := detectTokenType(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}
