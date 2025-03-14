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
			expectedFuncs:   []Func{{Name: "trim", Args: []any{}}},
		},
		{
			name:            "variable with shorten function",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: "content | shorten:5"},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "shorten", Args: []any{"5"}}},
		},
		{
			name:            "multiple functions",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: "content | trim | shorten:5 | length"},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "trim", Args: []any{}},
				{Name: "shorten", Args: []any{"5"}},
				{Name: "length", Args: []any{}},
			},
		},
		{
			name:            "args with single quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: "content | shorten:'5'"},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "shorten", Args: []any{"5"}}},
		},
		{
			name:            "args with double quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:"5"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "shorten", Args: []any{"5"}}},
		},
		{
			name:            "multiple args in multiple funcs with single quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:'5' | shorten:'10'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "shorten", Args: []any{"5"}},
				{Name: "shorten", Args: []any{"10"}},
			},
		},
		{
			name:            "multiple args in multiple funcs with double quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:"5" | shorten:"10"`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "shorten", Args: []any{"5"}},
				{Name: "shorten", Args: []any{"10"}},
			},
		},
		{
			name:            "multiple args in multiple funcs with mixed quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:"5" | shorten:'10'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "shorten", Args: []any{"5"}},
				{Name: "shorten", Args: []any{"10"}},
			},
		},
		{
			name:            "args with spaces",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | shorten:"5 10"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "shorten", Args: []any{"5 10"}}},
		},
		{
			name:            "escaped single quotes inside single quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | escape:'don\'t'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "escape", Args: []any{"don't"}}},
		},
		{
			name:            "escaped double quotes inside double quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | escape:"\"hello\""`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "escape", Args: []any{`"hello"`}}},
		},
		{
			name:            "mixed quote types with escapes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | mixQuotes:'"double" and \'single\''`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "mixQuotes", Args: []any{`"double" and 'single'`}}},
		},
		{
			name:            "multiple escaped args to single func",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | multiEscape:'don\'t', "\"think\"", 'so'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "multiEscape", Args: []any{"don't", `"think"`, "so"}}},
		},
		{
			name:            "single quotes with escaped single quotes and spaces",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | withSpace:'don\'t stop'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "withSpace", Args: []any{"don't stop"}}},
		},
		{
			name:            "double quotes with escaped double quotes and spaces",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | withSpace:"\"start over\""`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "withSpace", Args: []any{`"start over"`}}},
		},
		{
			name:            "mixed quotes with multiple escapes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | mixed:"escaped \'single\'", 'escaped "double"'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "mixed", Args: []any{"escaped \\'single\\'", `escaped "double"`}}},
		},
		{
			name:            "multiple functions with escaped args",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | escape1:'don\'t' | escape2:"\"think\"" | escape3:'so'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "escape1", Args: []any{"don't"}},
				{Name: "escape2", Args: []any{`"think"`}},
				{Name: "escape3", Args: []any{"so"}},
			},
		},
		{
			name:            "single func with multiple mixed escaped args",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | multiMixed:'single\'s', "\"doubles\"", "mix\'ed"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "multiMixed", Args: []any{"single's", `"doubles"`, `mix\'ed`}}},
		},
		{
			name:            "escaped args with spaces in multiple funcs",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | first:"first escape \'with spaces\'" | second:'second escape "also with spaces"'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "first", Args: []any{"first escape \\'with spaces\\'"}},
				{Name: "second", Args: []any{`second escape "also with spaces"`}},
			},
		},
		{
			name:            "complex mixed escapes and funcs",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | complex:'mix\'ed', "dou\"ble", 'sin\'gle' | another:"esc\'ape", 'more "complexity"'`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "complex", Args: []any{"mix'ed", `dou"ble`, "sin'gle"}},
				{Name: "another", Args: []any{`esc\'ape`, `more "complexity"`}},
			},
		},
		{
			name:            "escaped commas in args",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | withComma:'arg,ument', "ano,ther"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "withComma", Args: []any{"arg,ument", "ano,ther"}}},
		},
		{
			name:            "multiple funcs with mixed escaped quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | first:'"double"' | second:"'single'"`},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "first", Args: []any{`"double"`}},
				{Name: "second", Args: []any{"'single'"}},
			},
		},
		{
			name:            "nested quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | nested:'"start \'nested\' end"'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "nested", Args: []any{`"start 'nested' end"`}}},
		},
		{
			name:            "deeply nested quotes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | deeplyNested:"'start \"deep \'deeper\' deep\" end'"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "deeplyNested", Args: []any{"'start \"deep \\'deeper\\' deep\" end'"}}},
		},
		{
			name:            "single quotes with backslash",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | backslash:'\\'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "backslash", Args: []any{"\\\\"}}},
		},
		{
			name:            "double quotes with backslash",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | backslash:"\\"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "backslash", Args: []any{"\\\\"}}},
		},
		{
			name:            "multiple backslashes",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | multiBackslash:"\\\\"`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "multiBackslash", Args: []any{"\\\\\\\\"}}},
		},
		{
			name:            "backslashes and quotes mixed",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | mixedBackslashes:'\\"\\'\\\\"'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "mixedBackslashes", Args: []any{"\\\\\"\\'\\\\\\\\\""}}},
		},
		{
			name:            "complex scenario with all elements",
			token:           BaseToken{TokenType: FilteredVariableToken, RawValue: `content | complexAll:'don\'t "mix"', "seriously, \'don't", 'but why? \\\\'`},
			expectedVarName: "content",
			expectedFuncs:   []Func{{Name: "complexAll", Args: []any{"don't \"mix\"", "seriously, \\'don't", "but why? \\\\\\\\"}}},
		},
		{
			name: "double function complex scenario with all elements",
			token: BaseToken{
				TokenType: FilteredVariableToken,
				RawValue:  `content | complexAll:'don\'t "mix"', "seriously, \'don\'t", 'but why? \\\\' | complexAll2:'don\'t "mix"', "seriously, \'don\'t", 'but why? \\\\'`,
			},
			expectedVarName: "content",
			expectedFuncs: []Func{
				{Name: "complexAll", Args: []any{"don't \"mix\"", "seriously, \\'don\\'t", "but why? \\\\\\\\"}},
				{Name: "complexAll2", Args: []any{"don't \"mix\"", "seriously, \\'don\\'t", "but why? \\\\\\\\"}},
			},
		},
	}

	r := NewStringRenderer(map[string]GlobalModifier{
		"trim":    trim,
		"shorten": shorten,
		"length":  length,
	})

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
