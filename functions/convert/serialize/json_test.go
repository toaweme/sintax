package serialize

import (
	"encoding/json"
	"testing"

	"github.com/toaweme/sintax/assert"
)

// mustJSON runs JSON and fails the test on error or a non-string result.
func mustJSON(t *testing.T, value any, params []any) string {
	t.Helper()
	out, err := JSON(value, params)
	assert.NoError(t, err)
	s, ok := out.(string)
	if !ok {
		t.Fatalf("JSON returned %T, want string", out)
	}
	return s
}

// Test_JSON_Compact covers the default compact serialization. Map keys are
// asserted only on single-key inputs, since Go's encoding/json sorts keys and
// multi-key ordering is covered separately by round-trip parsing.
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
			assert.Equal(t, tt.expected, mustJSON(t, tt.value, nil))
		})
	}
}

// Test_JSON_KeysSorted proves multi-key objects come out with keys sorted
// alphabetically, which is what makes the output deterministic. The keys of the
// input map are deliberately not in alphabetical order.
func Test_JSON_KeysSorted(t *testing.T) {
	in := map[string]any{"region": "eu-west-1", "debug": false}
	assert.Equal(t, `{"debug":false,"region":"eu-west-1"}`, mustJSON(t, in, nil))
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
	out := mustJSON(t, in, nil)

	var back map[string]any
	assert.NoError(t, json.Unmarshal([]byte(out), &back))
	assert.Equal(t, "Alice", back["name"])
	assert.Equal(t, "admin", back["role"])
	assert.Len(t, back["tags"].([]any), 2)
	assert.Equal(t, true, back["meta"].(map[string]any)["active"])
}

// Test_JSON_Pretty covers indented output selected by the 'pretty' param. Keys
// still sort alphabetically, so debug precedes region in the output.
func Test_JSON_Pretty(t *testing.T) {
	in := map[string]any{"region": "eu-west-1", "debug": false}
	expected := "{\n  \"debug\": false,\n  \"region\": \"eu-west-1\"\n}"
	assert.Equal(t, expected, mustJSON(t, in, []any{"pretty"}))
}

// Test_JSON_PrettySingleKey locks in the exact two-space indentation on a
// single-key object so the layout is unambiguous regardless of key sorting.
func Test_JSON_PrettySingleKey(t *testing.T) {
	in := map[string]any{"name": "Alice"}
	assert.Equal(t, "{\n  \"name\": \"Alice\"\n}", mustJSON(t, in, []any{"pretty"}))
}

// Test_JSON_ParamVariants proves only the exact 'pretty' literal triggers
// indentation. Any other param, or none, produces compact output.
func Test_JSON_ParamVariants(t *testing.T) {
	in := map[string]any{"name": "Alice"}
	compact := `{"name":"Alice"}`
	tests := []struct {
		name     string
		params   []any
		expected string
	}{
		{"no params", nil, compact},
		{"empty params", []any{}, compact},
		{"unknown param stays compact", []any{"nice"}, compact},
		{"wrong type param stays compact", []any{true}, compact},
		{"pretty in later slot ignored", []any{"x", "pretty"}, compact},
		{"pretty triggers indent", []any{"pretty"}, "{\n  \"name\": \"Alice\"\n}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, mustJSON(t, in, tt.params))
		})
	}
}

// Test_JSON_UnmarshalableErrors covers the error path. A channel cannot be
// serialized to JSON, so the modifier returns a wrapped error and an empty
// string.
func Test_JSON_UnmarshalableErrors(t *testing.T) {
	out, err := JSON(make(chan int), nil)
	assert.Error(t, err)
	assert.Equal(t, "", out)

	out, err = JSON(make(chan int), []any{"pretty"})
	assert.Error(t, err)
	assert.Equal(t, "", out)
}
