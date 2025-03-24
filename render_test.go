package sintax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringRenderer_RenderVariable(t *testing.T) {
	type testCase struct {
		name        string
		token       Token
		vars        map[string]any
		expected    string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:     "basic variable",
			token:    BaseToken{TokenType: VariableToken, RawValue: "content", Var: "content"},
			vars:     map[string]any{"content": "Hello, World!"},
			expected: "Hello, World!",
		},
		{
			name:     "variable with trim function",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | trim", Var: "content"},
			vars:     map[string]any{"content": " Hello, World! "},
			expected: "Hello, World!",
		},
		{
			name:     "variable with shorten function",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | shorten:5", Var: "content"},
			vars:     map[string]any{"content": "Hello, World!"},
			expected: "Hello",
		},
		{
			name:     "variable with length function",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | length", Var: "content"},
			vars:     map[string]any{"content": "Hello, World!"},
			expected: "13",
		},
		{
			name:        "variable not found",
			token:       BaseToken{TokenType: VariableToken, RawValue: "anything_wrong", Var: "anything_wrong"},
			vars:        map[string]any{"content": "Hello, World!"},
			expected:    "",
			expectedErr: ErrVariableNotFound,
		},
		{
			name:        "invalid token type",
			token:       BaseToken{TokenType: IfToken, RawValue: "condition"},
			vars:        map[string]any{"content": "Hello, World!"},
			expected:    "",
			expectedErr: ErrInvalidTokenType,
		},
		{
			name:        "invalid function",
			token:       BaseToken{TokenType: FilteredVariableToken, RawValue: "content | invalid", Var: "content"},
			vars:        map[string]any{"content": "Hello, World!"},
			expected:    "",
			expectedErr: ErrFunctionNotFound,
		},
		{
			// trim: "Hello, World!"
			// shorten: "Hello"
			// length: "5"
			name:     "multiple functions",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | trim | shorten:5 | length", Var: "content"},
			vars:     map[string]any{"content": " Hello, World! "},
			expected: "5",
		},
		{
			name:     "multiple functions",
			token:    BaseToken{TokenType: FilteredVariableToken, RawValue: "content | trim | shorten:5", Var: "content"},
			vars:     map[string]any{"content": "Hello"},
			expected: "Hello",
		},
		{
			name: "pick key function",
			token: BaseToken{
				TokenType: FilteredVariableToken,
				RawValue:  "mapping | key:'content'",
				Var:       "content",
			},
			vars: map[string]any{
				"mapping": map[string]any{
					"content": "Hello",
				},
			},
			expected: "Hello",
		},
	}

	r := NewStringRenderer(BuiltinFunctions)

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := r.renderVariable(tt.token, tt.vars)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_StringRenderer_getVarAndFunctions(t *testing.T) {
	type testCase struct {
		name            string
		token           Token
		expectedVarName string
		expectedFuncs   []Func
	}

	testCases := []testCase{
		{
			name:            "basic variable",
			token:           BaseToken{TokenType: VariableToken, RawValue: "content"},
			expectedVarName: "content",
			expectedFuncs:   []Func{},
		},
		{
			name:            "variable with trim function",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: "content | trim"},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "trim", Args: []Arg{}}},
		},
		{
			name:            "variable with shorten function",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: "content | shorten:5"},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "shorten", Args: []Arg{{Value: 5}}}},
		},
		{
			name:            "multiple functions",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: "content | trim | shorten:5 | length"},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "trim", Args: []Arg{}},
				{Name: "shorten", Args: []Arg{{Value: 5}}},
				{Name: "length", Args: []Arg{}},
			},
		},
		{
			name:            "args with single quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: "content | shorten:'5'"},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "shorten", Args: []Arg{{Value: "5"}}}},
		},
		{
			name:            "args with double quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:"5"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "shorten", Args: []Arg{{Value: "5"}}}},
		},
		{
			name:            "multiple args in multiple funcs with single quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:'5' | shorten:'10'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "shorten", Args: []Arg{{Value: "5"}}},
				{Name: "shorten", Args: []Arg{{Value: "10"}}},
			},
		},
		{
			name:            "multiple args in multiple funcs with double quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:"5" | shorten:"10"`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "shorten", Args: []Arg{{Value: "5"}}},
				{Name: "shorten", Args: []Arg{{Value: "10"}}},
			},
		},
		{
			name:            "multiple args in multiple funcs with mixed quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:"5" | shorten:'10'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "shorten", Args: []Arg{{Value: "5"}}},
				{Name: "shorten", Args: []Arg{{Value: "10"}}},
			},
		},
		{
			name:            "args with spaces",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:"5 10"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "shorten", Args: []Arg{{Value: "5 10"}}}},
		},
		{
			name:            "escaped single quotes inside single quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | escape:'don\'t'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "escape", Args: []Arg{{Value: "don't"}}}},
		},
		{
			name:            "escaped double quotes inside double quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | escape:"\"hello\""`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "escape", Args: []Arg{{Value: `"hello"`}}}},
		},
		{
			name:            "mixed quote types with escapes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | mixQuotes:'"double" and \'single\''`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "mixQuotes", Args: []Arg{{Value: `"double" and 'single'`}}}},
		},
		{
			name:            "multiple escaped args to single func",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | multiEscape:'don\'t', "\"think\"", 'so'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "multiEscape", Args: []Arg{{Value: "don't"}, {Value: `"think"`}, {Value: "so"}}}},
		},
		{
			name:            "single quotes with escaped single quotes and spaces",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | withSpace:'don\'t stop'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "withSpace", Args: []Arg{{Value: "don't stop"}}}},
		},
		{
			name:            "double quotes with escaped double quotes and spaces",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | withSpace:"\"start over\""`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "withSpace", Args: []Arg{{Value: `"start over"`}}}},
		},
		{
			name:            "mixed quotes with multiple escapes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | mixed:"escaped \'single\'", 'escaped "double"'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "mixed", Args: []Arg{{Value: "escaped \\'single\\'"}, {Value: `escaped "double"`}}}},
		},
		{
			name:            "multiple functions with escaped args",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | escape1:'don\'t' | escape2:"\"think\"" | escape3:'so'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "escape1", Args: []Arg{{Value: "don't"}}},
				{Name: "escape2", Args: []Arg{{Value: `"think"`}}},
				{Name: "escape3", Args: []Arg{{Value: "so"}}},
			},
		},
		{
			name:            "single func with multiple mixed escaped args",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | multiMixed:'single\'s', "\"doubles\"", "mix\'ed"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "multiMixed", Args: []Arg{{Value: "single's"}, {Value: `"doubles"`}, {Value: `mix\'ed`}}}},
		},
		{
			name:            "escaped args with spaces in multiple funcs",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | first:"first escape \'with spaces\'" | second:'second escape "also with spaces"'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "first", Args: []Arg{{Value: "first escape \\'with spaces\\'"}}},
				{Name: "second", Args: []Arg{{Value: `second escape "also with spaces"`}}},
			},
		},
		{
			name:            "complex mixed escapes and funcs",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | complex:'mix\'ed', "dou\"ble", 'sin\'gle' | another:"esc\'ape", 'more "complexity"'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "complex", Args: []Arg{{Value: "mix'ed"}, {Value: `dou"ble`}, {Value: "sin'gle"}}},
				{Name: "another", Args: []Arg{{Value: `esc\'ape`}, {Value: `more "complexity"`}}},
			},
		},
		{
			name:            "escaped commas in args",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | withComma:'arg,ument', "ano,ther"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "withComma", Args: []Arg{{Value: "arg,ument"}, {Value: "ano,ther"}}}},
		},
		{
			name:            "multiple funcs with mixed escaped quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | first:'"double"' | second:"'single'"`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "first", Args: []Arg{{Value: `"double"`}}},
				{Name: "second", Args: []Arg{{Value: "'single'"}}},
			},
		},
		{
			name:            "nested quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | nested:'"start \'nested\' end"'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "nested", Args: []Arg{{Value: `"start 'nested' end"`}}}},
		},
		{
			name:            "deeply nested quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | deeplyNested:"'start \"deep \'deeper\' deep\" end'"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "deeplyNested", Args: []Arg{{Value: "'start \"deep \\'deeper\\' deep\" end'"}}}},
		},
		{
			name:            "single quotes with backslash",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | backslash:'\\'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "backslash", Args: []Arg{{Value: "\\\\"}}}},
		},
		{
			name:            "double quotes with backslash",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | backslash:"\\"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "backslash", Args: []Arg{{Value: "\\\\"}}}},
		},
		{
			name:            "multiple backslashes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | multiBackslash:"\\\\"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "multiBackslash", Args: []Arg{{Value: "\\\\\\\\"}}}},
		},
		{
			name:            "backslashes and quotes mixed",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | mixedBackslashes:'\\"\\'\\\\"'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "mixedBackslashes", Args: []Arg{{Value: "\\\\\"\\'\\\\\\\\\""}}}},
		},
		{
			name:            "complex scenario with all elements",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | complexAll:'don\'t "mix"', "seriously, \'don't", 'but why? \\\\'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "complexAll", Args: []Arg{{Value: "don't \"mix\""}, {Value: "seriously, \\'don't"}, {Value: "but why? \\\\\\\\"}}}},
		},
		{
			name: "double function complex scenario with all elements",
			token: BaseToken{
				TokenType: FilteredVariableToken,
				RawValue:  `content | complexAll:'don\'t "mix"', "seriously, \'don\'t", 'but why? \\\\' | complexAll2:'don\'t "mix"', "seriously, \'don\'t", 'but why? \\\\'`,
			},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "complexAll", Args: []Arg{{Value: "don't \"mix\""}, {Value: "seriously, \\'don\\'t"}, {Value: "but why? \\\\\\\\"}}},
				{Name: "complexAll2", Args: []Arg{{Value: "don't \"mix\""}, {Value: "seriously, \\'don\\'t"}, {Value: "but why? \\\\\\\\"}}},
			},
		},
	}

	r := NewStringRenderer(BuiltinFunctions)

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			varName, funcs := r.getVarAndFunctions(tt.token)
			assert.Equal(t, tt.expectedVarName, varName)
			assert.Equal(t, tt.expectedFuncs, funcs)
		})
	}
}

func Test_splitRespectingQuotes_simple(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		sep      string
		expected []string
	}

	testCases := []testCase{
		{
			name:     "basic split",
			input:    "Hello, World",
			sep:      ",",
			expected: []string{"Hello", "World"},
		},
		{
			name:     "quoted section",
			input:    `"Hello, World", Test`,
			sep:      ",",
			expected: []string{`"Hello, World"`, "Test"},
		},
		{
			name:     "escaped quotes",
			input:    `"Hello, \"World\"", Test`,
			sep:      ",",
			expected: []string{`"Hello, \"World\""`, "Test"},
		},
		{
			name:     "multiple quoted sections",
			input:    `"Hello, World", "Test, Case"`,
			sep:      ",",
			expected: []string{`"Hello, World"`, `"Test, Case"`},
		},
		{
			name:     "mixed quotes",
			input:    `"Hello, 'World'", Test`,
			sep:      ",",
			expected: []string{`"Hello, 'World'"`, "Test"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := splitRespectingQuotes(tt.input, tt.sep)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_splitRespectingQuotes_pipes(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		sep      string
		expected []string
	}

	testCases := []testCase{
		{
			name:     "basic split",
			input:    "Hello | World",
			sep:      "|",
			expected: []string{"Hello", "World"},
		},
		{
			name:  "quoted section",
			input: `content | complexAll:'don\'t "mix"', "seriously, \'don\'t", 'but why? \\\\' | complexAll2:'don\'t "mix"', "seriously, \'don\'t", 'but why? \\\\'`,
			sep:   "|",
			expected: []string{
				"content",
				`complexAll:'don\'t "mix"', "seriously, \'don\'t", 'but why? \\\\'`,
				`complexAll2:'don\'t "mix"', "seriously, \'don\'t", 'but why? \\\\'`,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := splitRespectingQuotes(tt.input, tt.sep)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
