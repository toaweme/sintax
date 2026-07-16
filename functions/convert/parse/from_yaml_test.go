package parse

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

// Test_FromYAML_StubReturnsError proves the shipped stub returns an error until
// a consumer injects a real YAML codec. The document is irrelevant to the stub,
// so a range of inputs all take the same path.
func Test_FromYAML_StubReturnsError(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"mapping", "region: eu-west-1"},
		{"empty string", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := FromYAML(tt.value)
			assert.Error(t, err)
			assert.Equal(t, map[string]any(nil), out)
			assert.Equal(t, "from_yaml function needs to be injected", err.Error())
		})
	}
}

// Test_FromYAML_Dispatch proves the registered modifier surfaces the stub error
// too, so a template naming from_yaml without an injected codec fails loudly
// rather than rendering nothing.
func Test_FromYAML_Dispatch(t *testing.T) {
	_, err := fromYAMLModifier("region: eu-west-1", nil)
	assert.Error(t, err)
}
