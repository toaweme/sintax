package sintax

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/toaweme/log"
)

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
			name: "interpolated system and config variables",
			systemVars: map[string]any{
				"global1": "globalvalue1",
				"global2": "globalvalue2",
			},
			vars: map[string]any{
				"var1": "{{ global1 }}",
				"var2": "v2-{{ var1 }}",
			},
			expected: map[string]any{
				"global1": "globalvalue1",
				"global2": "globalvalue2",
				"var1":    "globalvalue1",
				"var2":    "v2-globalvalue1",
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
