package sintax

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/toaweme/log"
)

func Test_Sintax_ResolveVariables(t *testing.T) {
	type testCase struct {
		name        string
		vars        map[string]any
		expected    map[string]any
		expectedErr error
	}

	now := time.Now()
	formatted := now.Format("2006-01-02-15:04:05")
	formattedDate := now.Format("2006-01-02")

	testCases := []testCase{
		{
			name: "var not found",
			vars: map[string]any{
				"base": "root",
				"varD": `{{ varC | default:"1" }}-D`,
			},
			expected: map[string]any{
				"base": "root",
				"varD": "root-A-B-C-D",
			},
			expectedErr: ErrVariableNotFound,
		},
		{
			name: "interpolated variables",
			vars: map[string]any{
				"now":                 now,
				"ext":                 "md",
				"base_name":           `{{ now | format:"Y-m-d-H:i:s" }}`,
				"output_file":         "data/daily/test-{{ base_name }}.{{ ext }}",
				"output_file_content": `date: {{ now | format:"Y-m-d" }}`,
				"path":                "{{ output_file }}",
				"content":             "```shell\n{{ path }}\n{{ output_file_content }}```",
			},
			expected: map[string]any{
				"now":                 now,
				"ext":                 "md",
				"base_name":           formatted,
				"output_file":         "data/daily/test-" + formatted + ".md",
				"output_file_content": "date: " + formattedDate,
				"path":                "data/daily/test-" + formatted + ".md",
				"content":             "```shell\ndata/daily/test-" + formatted + ".md\ndate: " + formattedDate + "```",
			},
		},
		{
			name: "complex dependency tree with filters",
			vars: map[string]any{
				"base": "root",
				"varA": "{{ base }}-A",
				"varB": "{{ varA }}-B",
				"varC": "{{ varB }}-C",
				"varD": `{{ varC | default:"1" }}-D`,
			},
			expected: map[string]any{
				"base": "root",
				"varA": "root-A",
				"varB": "root-A-B",
				"varC": "root-A-B-C",
				"varD": "root-A-B-C-D",
			},
		},
		{
			name: "no interpolated variables",
			vars: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			expected: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "inter-dependant interpolated variables",
			vars: map[string]any{
				"global1": "globalvalue1",
				"global2": "globalvalue2",
				"var1":    "{{ global1 }}",
				"var2":    "v2-{{ var1 }}",
				"var3":    50,
				"var4":    "{{ var2}}:{{ var3 }}",
			},
			expected: map[string]any{
				"global1": "globalvalue1",
				"global2": "globalvalue2",
				"var1":    "globalvalue1",
				"var2":    "v2-globalvalue1",
				"var3":    50,
				"var4":    "v2-globalvalue1:50",
			},
		},
		{
			name: "nested interpolation",
			vars: map[string]any{
				"host":   "localhost",
				"url":    "http://{{ host }}:8080",
				"apiUrl": "{{ url }}/api",
			},
			expected: map[string]any{
				"host":   "localhost",
				"url":    "http://localhost:8080",
				"apiUrl": "http://localhost:8080/api",
			},
		},
		{
			name: "variable referencing itself (error case)",
			vars: map[string]any{
				"self": "{{ self }}",
			},
			expectedErr: ErrCircularDependency,
		},
		{
			name: "complex dependency tree",
			vars: map[string]any{
				"base": "root",
				"varD": "{{ varC }}-D",
				"varA": "{{ base }}-A",
				"varB": "{{ varA }}-B",
				"varC": "{{ varB }}-C",
			},
			expected: map[string]any{
				"base": "root",
				"varA": "root-A",
				"varB": "root-A-B",
				"varC": "root-A-B-C",
				"varD": "root-A-B-C-D",
			},
		},
		{
			name: "integer and boolean values",
			vars: map[string]any{
				"intVar":       123,
				"boolVar":      true,
				"boolFalseVar": false,
				"stringVar":    "{{ intVar }}-{{ boolVar }}",
				"stringVar2":   "{{ intVar }}-{{ boolFalseVar }}",
			},
			expected: map[string]any{
				"intVar":       123,
				"boolVar":      true,
				"boolFalseVar": false,
				"stringVar":    "123-true",
				"stringVar2":   "123-false",
			},
		},
		{
			name: "vars override is not allowed",
			vars: map[string]any{
				"override": "{{ override }}-modified",
			},
			expectedErr: ErrCircularDependency,
		},
		{
			name: "multi-level nested map",
			vars: map[string]any{
				"level1": map[string]any{
					"level2": map[string]any{
						"var": "value",
					},
				},
			},
			expected: map[string]any{
				"level1": map[string]any{
					"level2": map[string]any{
						"var": "value",
					},
				},
			},
		},
		{
			name: "empty variable values",
			vars: map[string]any{
				"emptyString": "",
				"nilValue":    nil,
			},
			expected: map[string]any{
				"emptyString": "",
				"nilValue":    nil,
			},
		},
		{
			name: "independent vars that falsely look like vars",
			vars: map[string]any{
				"a": "{ something | boop }}",
				"b": "plain text",
			},
			expected: map[string]any{
				"a": "{ something | boop }}",
				"b": "plain text",
			},
		},
		{
			name: "independent vars that falsely look like vars",
			vars: map[string]any{
				"a": "{ something | boop }}}}",
				"b": "plain text",
			},
			expected: map[string]any{
				"a": "{ something | boop }}}}",
				"b": "plain text",
			},
		},
		{
			name: "keeps JSON strings as is",
			vars: map[string]any{
				"a": `{"beep":"}} some {{ text"}`,
				"b": "plain text",
			},
			expected: map[string]any{
				"a": `{"beep":"}} some {{ text"}`,
				"b": "plain text",
			},
		},
		{
			name: "keeps JSON strings as is",
			vars: map[string]any{
				"a": `{"beep":"}} some {{ text"}`,
				"b": "plain text",
			},
			expected: map[string]any{
				"a": `{"beep":"}} some {{ text"}`,
				"b": "plain text",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(BuiltinFunctions)
			actual, err := s.ResolveVariables(tc.vars)
			if tc.expectedErr != nil {
				log.Debug("expected", "error", err)
				assert.ErrorIs(t, err, tc.expectedErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
func Test_Sintax_ResolveVariablesFunc(t *testing.T) {
	type testCase struct {
		name        string
		vars        map[string]any
		expected    map[string]any
		expectedErr error
	}

	now := time.Now()
	formatted := now.Format("2006-01-02-15:04:05")

	testCases := []testCase{
		{
			name: "interpolated variables",
			vars: map[string]any{
				"now":       now,
				"base_name": `{{ now | format:"Y-m-d-H:i:s" }}`,
			},
			expected: map[string]any{
				"now":       now,
				"base_name": formatted,
			},
		},
		{
			name: "complex dependency tree with filters",
			vars: map[string]any{
				"base": "root",
				"varA": "{{ base }}-A",
				"varB": "{{ varA }}-B",
				"varC": "{{ varB }}-C",
				"varD": `{{ varC | default:"1" }}-D`,
			},
			expected: map[string]any{
				"base": "root",
				"varA": "root-A",
				"varB": "root-A-B",
				"varC": "root-A-B-C",
				"varD": "root-A-B-C-D",
			},
		},
		// {
		// 	name: "nested dependency tree with filters",
		// 	vars: map[string]any{
		// 		"base_url": "http://localhost:8080",
		// 		"url":      "{{ base_url }}/endpoint",
		// 		"method":   "POST",
		// 		"body":     "{{ body_map | json }}",
		// 		"obj": map[string]any{
		// 			"id":       "xkcd",
		// 			"title":    "some title",
		// 			"currency": "USD",
		// 		},
		// 		"body_map": map[string]any{
		// 			"title":       "{{ obj | key:title }} {{ obj | key:currency | default:USD }}",
		// 			"ai_model":    "{{ obj | key:id }}",
		// 			"cost_input":  11,
		// 			"cost_output": 22,
		// 			"cost_cached": false,
		// 		},
		// 	},
		// 	expected: map[string]any{
		// 		"base_url": "http://localhost:8080",
		// 		"url":      "http://localhost:8080/endpoint",
		// 		"method":   "POST",
		// 		"body":     `{"title":"some title USD","ai_model":"xkcd","cost_input":11,"cost_output":22,"cost_cached":false}`,
		// 		"obj": map[string]any{
		// 			"id":       "xkcd",
		// 			"title":    "some title",
		// 			"currency": "USD",
		// 		},
		// 		"body_map": map[string]any{
		// 			"title":       "some title USD",
		// 			"ai_model":    "xkcd",
		// 			"cost_input":  11,
		// 			"cost_output": 22,
		// 			"cost_cached": false,
		// 		},
		// 	},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(BuiltinFunctions)
			actual, err := s.ResolveVariables(tc.vars)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_Sintax_Render(t *testing.T) {
	type testCase struct {
		name        string
		input       string
		vars        map[string]any
		expected    string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:  "key func render",
			input: `{{ mapping | key:content }}`,
			vars: map[string]any{
				"mapping": map[string]any{
					"content": "Hello",
				},
			},
			expected: "Hello",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(BuiltinFunctions)
			actual, err := s.Render(tc.input, tc.vars)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
