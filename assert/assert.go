// Package assert provides minimal test assertion helpers built on stdlib
// testing, avoiding a third-party test framework dependency.
package assert

import (
	"errors"
	"reflect"
	"testing"
)

// Equal fails the test if expected and actual are not deeply equal.
func Equal(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v (%T), got %v (%T)", expected, expected, actual, actual)
	}
}

// NoError fails the test if err is non-nil.
func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// ErrorIs fails the test unless errors.Is(err, target) holds.
func ErrorIs(t *testing.T, err, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Fatalf("expected error %v, got %v", target, err)
	}
}

// Error fails the test if err is nil.
func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// True fails the test if val is false, using msgAndArgs as an optional
// Fatalf-style message.
func True(t *testing.T, val bool, msgAndArgs ...any) {
	t.Helper()
	if !val {
		if len(msgAndArgs) > 0 {
			t.Fatalf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		}
		t.Fatal("expected true, got false")
	}
}

// Empty fails the test if v is a non-empty slice.
func Empty(t *testing.T, v any) {
	t.Helper()
	if v == nil {
		return
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Slice && rv.Len() != 0 {
		t.Fatalf("expected empty, got length %d", rv.Len())
	}
}

// Len fails the test if v does not have the expected length.
func Len(t *testing.T, v any, expected int) {
	t.Helper()
	rv := reflect.ValueOf(v)
	if rv.Len() != expected {
		t.Fatalf("expected length %d, got %d", expected, rv.Len())
	}
}
