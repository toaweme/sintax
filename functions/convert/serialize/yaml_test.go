package serialize

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

// Test_YAML_StubReturnsError proves the shipped stub returns an error until a
// consumer injects a real YAML codec. The value is irrelevant to the stub, so a
// range of inputs all take the same path.
func Test_YAML_StubReturnsError(t *testing.T) {
	tests := []struct {
		name  string
		value any
	}{
		{"config map", map[string]any{"region": "eu-west-1", "debug": false}},
		{"string", "region: eu-west-1"},
		{"empty string", ""},
		{"nil", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := YAML(tt.value)
			assert.Error(t, err)
			assert.Equal(t, "", out)
			assert.Equal(t, "yaml function needs to be injected", err.Error())
		})
	}
}

// Test_YAML_Dispatch proves the registered modifier surfaces the stub error too.
func Test_YAML_Dispatch(t *testing.T) {
	yamlMod := yamlModifier
	out, err := yamlMod("region: eu-west-1", nil)
	assert.Error(t, err)
	assert.Equal(t, "", out)
	assert.Equal(t, "yaml function needs to be injected", err.Error())
}
