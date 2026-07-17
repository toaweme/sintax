package serialize

import (
	"encoding/json"
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// Test_JSON_Compact covers the default compact serialization via the typed
// clause. Map keys are asserted only on single-key inputs, since Go's
// encoding/json sorts keys and multi-key ordering is covered separately by
// round-trip parsing.
func Test_JSON_Compact(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"single key map", map[string]any{"name": "Alice"}, `{"name":"Alice"}`},
		{"string scalar", "hello", `"hello"`},
		{"int scalar", 42, "42"},
		{"float scalar", 3.14, "3.14"},
		{"bool true", true, "true"},
		{"bool false", false, "false"},
		{"nil is null", nil, "null"},
		{"int slice", []any{1, 2, 3}, "[1,2,3]"},
		{"string slice", []string{"a", "b"}, `["a","b"]`},
		{"empty slice", []any{}, "[]"},
		{"empty map", map[string]any{}, "{}"},
		{"empty string", "", `""`},
		{"nested", map[string]any{"user": map[string]any{"id": 1}}, `{"user":{"id":1}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := JSON(tt.value)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_JSON_KeysSorted proves multi-key objects come out with keys sorted
// alphabetically, which is what makes the output deterministic. The keys of the
// input map are deliberately not in alphabetical order.
func Test_JSON_KeysSorted(t *testing.T) {
	in := map[string]any{"region": "eu-west-1", "debug": false}
	out, err := JSON(in)
	assert.NoError(t, err)
	assert.Equal(t, `{"debug":false,"region":"eu-west-1"}`, out)
}

// Test_JSON_RoundTrip serializes a nested structure and parses it back, proving
// the output is valid JSON that preserves the data without asserting on the
// brittle key ordering of the raw string.
func Test_JSON_RoundTrip(t *testing.T) {
	in := map[string]any{
		"name": "Alice",
		"role": "admin",
		"tags": []any{"x", "y"},
		"meta": map[string]any{"active": true},
	}
	out, err := JSON(in)
	assert.NoError(t, err)

	var back map[string]any
	assert.NoError(t, json.Unmarshal([]byte(out), &back))
	assert.Equal(t, "Alice", back["name"])
	assert.Equal(t, "admin", back["role"])
	assert.Len(t, back["tags"].([]any), 2)
	assert.Equal(t, true, back["meta"].(map[string]any)["active"])
}

// Test_JSON_Pretty covers indented output selected by the 'pretty' mode. Keys
// still sort alphabetically, so debug precedes region in the output.
func Test_JSON_Pretty(t *testing.T) {
	in := map[string]any{"region": "eu-west-1", "debug": false}
	expected := "{\n  \"debug\": false,\n  \"region\": \"eu-west-1\"\n}"
	out, err := JSONMode(in, "pretty")
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

// Test_JSON_PrettySingleKey locks in the exact two-space indentation on a
// single-key object so the layout is unambiguous regardless of key sorting.
func Test_JSON_PrettySingleKey(t *testing.T) {
	in := map[string]any{"name": "Alice"}
	out, err := JSONMode(in, "pretty")
	assert.NoError(t, err)
	assert.Equal(t, "{\n  \"name\": \"Alice\"\n}", out)
}

// Test_JSON_ModeVariants proves only the exact 'pretty' mode triggers
// indentation; any other mode string falls back to compact output.
func Test_JSON_ModeVariants(t *testing.T) {
	in := map[string]any{"name": "Alice"}
	compact := `{"name":"Alice"}`
	tests := []struct {
		name     string
		mode     string
		expected string
	}{
		{"unknown mode stays compact", "nice", compact},
		{"empty mode stays compact", "", compact},
		{"pretty triggers indent", "pretty", "{\n  \"name\": \"Alice\"\n}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := JSONMode(in, tt.mode)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_JSON_Dispatch drives the registered Overload, proving no-param calls take
// the compact clause and a 'pretty' string param takes the indent clause.
func Test_JSON_Dispatch(t *testing.T) {
	jsonMod := jsonModifier
	in := map[string]any{"name": "Alice"}

	tests := []struct {
		name     string
		params   []any
		expected string
	}{
		{"no params compact", nil, `{"name":"Alice"}`},
		{"empty params compact", []any{}, `{"name":"Alice"}`},
		{"unknown param compact", []any{"nice"}, `{"name":"Alice"}`},
		{"pretty param indents", []any{"pretty"}, "{\n  \"name\": \"Alice\"\n}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := jsonMod(in, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_JSON_DispatchRejectsBadParams proves the typed clauses reject a non-string
// param and any call with more than one param, since neither the no-param clause
// nor the one-string-param clause accepts those shapes.
func Test_JSON_DispatchRejectsBadParams(t *testing.T) {
	jsonMod := jsonModifier
	in := map[string]any{"name": "Alice"}

	_, err := jsonMod(in, []any{true})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)

	_, err = jsonMod(in, []any{"x", "pretty"})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)
}

// Test_JSON_UnmarshalableErrors covers the error path. A channel cannot be
// serialized to JSON, so both clauses return a wrapped error and an empty string.
func Test_JSON_UnmarshalableErrors(t *testing.T) {
	out, err := JSON(make(chan int))
	assert.Error(t, err)
	assert.Equal(t, "", out)

	out, err = JSONMode(make(chan int), "pretty")
	assert.Error(t, err)
	assert.Equal(t, "", out)
}
