package query

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Has_Slice(t *testing.T) {
	tags := []any{"featured", "sale", "new"}
	nums := []any{1, 2, 3}

	tests := []struct {
		name     string
		value    any
		params   []any
		expected bool
	}{
		{"element present", tags, []any{"featured"}, true},
		{"element absent", tags, []any{"clearance"}, false},
		{"any of several present", tags, []any{"clearance", "sale"}, true},
		{"number present", nums, []any{2}, true},
		{"number absent", nums, []any{9}, false},
		{"type mismatch string vs int", nums, []any{"2"}, false},
		{"empty slice is never a match", []any{}, []any{"featured"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Has(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Has_SliceOfMaps(t *testing.T) {
	items := []any{
		map[string]any{"name": "Coffee", "status": "sold-out"},
		map[string]any{"name": "Tea", "status": "active"},
	}

	tests := []struct {
		name     string
		value    any
		params   []any
		expected bool
	}{
		{"field value present", items, []any{"status", "active"}, true},
		{"field value absent", items, []any{"status", "pending"}, false},
		{"any of several field values present", items, []any{"status", "active", "pending"}, true},
		{"none of several field values present", items, []any{"status", "pending", "draft"}, false},
		{"unknown field", items, []any{"missing", "active"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Has(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Has_Map(t *testing.T) {
	config := map[string]any{"debug_mode": false, "region": "eu-west-1"}

	tests := []struct {
		name     string
		value    any
		params   []any
		expected bool
	}{
		{"key exists even when value is false", config, []any{"debug_mode"}, true},
		{"key exists", config, []any{"region"}, true},
		{"key missing", config, []any{"timeout"}, false},
		{"key and matching value", config, []any{"region", "eu-west-1"}, true},
		{"key and non-matching value", config, []any{"region", "us-east-1"}, false},
		{"key and any of several values", config, []any{"region", "us-east-1", "eu-west-1"}, true},
		{"missing key with value", config, []any{"timeout", 30}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Has(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Has_Errors covers the input-validation error paths: no parameters, a
// value that is neither a collection, and a slice of maps addressed with a
// single parameter (which needs both a key and at least one value).
func Test_Has_Errors(t *testing.T) {
	t.Run("no parameters", func(t *testing.T) {
		out, err := Has([]any{"a"}, nil)
		assert.Error(t, err)
		assert.Equal(t, false, out)
	})

	t.Run("value is a scalar", func(t *testing.T) {
		_, err := Has("hello", []any{"h"})
		assert.Error(t, err)
	})

	t.Run("slice of maps needs key and value", func(t *testing.T) {
		items := []any{map[string]any{"status": "active"}}
		_, err := Has(items, []any{"status"})
		assert.Error(t, err)
	})
}

func Test_Is(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected bool
	}{
		{"single match", "active", []any{"active"}, true},
		{"single no match", "archived", []any{"active"}, false},
		{"match any of several", "admin", []any{"admin", "superuser"}, true},
		{"none of several", "archived", []any{"active", "pending"}, false},
		{"number match", 5, []any{5}, true},
		{"number no match", 5, []any{6}, false},
		{"type mismatch int vs string", 5, []any{"5"}, false},
		{"bool match", true, []any{true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Is(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Is_NoParams covers the sole error path: Is with no candidate values.
func Test_Is_NoParams(t *testing.T) {
	out, err := Is("active", nil)
	assert.Error(t, err)
	assert.Equal(t, false, out)
}
