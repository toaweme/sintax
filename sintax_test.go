package sintax

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/toaweme/log"
)

func Test_Sintax_ResolveVariables(t *testing.T) {
	type testCase struct {
		name        string
		systemVars  map[string]any
		configVars  map[string]any
		actionVars  map[string]any
		outputVars  map[string]any
		expected    map[string]any
		expectedErr error
	}
	
	testCases := []testCase{
		{},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(BuiltinFunctions)
			actual, err := s.ResolveVariables(tc.systemVars, tc.configVars, tc.actionVars, tc.outputVars)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
				return
			}
			assert.NoError(t, err)
			log.Debug("actual", "vars", actual)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_Sintax_resolveVariables(t *testing.T) {
	type testCase struct {
		name        string
		systemVars  map[string]any
		vars        map[string]any
		expected    map[string]any
		expectedErr error
	}
	
	testCases := []testCase{
		{
			name: "no interpolated variables",
			systemVars: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			vars: map[string]any{},
			expected: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "inter-dependant interpolated variables",
			systemVars: map[string]any{
				"global1": "globalvalue1",
				"global2": "globalvalue2",
			},
			vars: map[string]any{
				"var1": "{{ global1 }}",
				"var2": "v2-{{ var1 }}",
				"var3": 50,
			},
			expected: map[string]any{
				"global1": "globalvalue1",
				"global2": "globalvalue2",
				"var1":    "globalvalue1",
				"var2":    "v2-globalvalue1",
				"var3":    50,
			},
		},
		{
			name: "nested interpolation",
			systemVars: map[string]any{
				"host": "localhost",
			},
			vars: map[string]any{
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
			name:       "variable referencing itself (error case)",
			systemVars: map[string]any{},
			vars: map[string]any{
				"self": "{{ self }}",
			},
			expectedErr: ErrCircularDependency,
		},
		{
			name: "complex dependency tree",
			systemVars: map[string]any{
				"base": "root",
			},
			vars: map[string]any{
				"varA": "{{ base }}-A",
				"varB": "{{ varA }}-B",
				"varC": "{{ varB }}-C",
				"varD": "{{ varC }}-D",
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
			name:       "integer and boolean values",
			systemVars: map[string]any{},
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
			systemVars: map[string]any{
				"override": "original",
			},
			vars: map[string]any{
				"override": "{{ override }}-modified",
			},
			expectedErr: ErrCircularDependency,
		},
		{
			name:       "multi-level nested map",
			systemVars: map[string]any{},
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
			name:       "empty variable values",
			systemVars: map[string]any{},
			vars: map[string]any{
				"emptyString": "",
				"nilValue":    nil,
			},
			expected: map[string]any{
				"emptyString": "",
				"nilValue":    nil,
			},
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(BuiltinFunctions)
			actual, err := s.resolveVariables(tc.systemVars, tc.vars)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
				return
			}
			assert.NoError(t, err)
			log.Debug("actual", "vars", actual)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
