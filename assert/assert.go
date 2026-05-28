package assert

import (
	"errors"
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v (%T), got %v (%T)", expected, expected, actual, actual)
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func ErrorIs(t *testing.T, err, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Fatalf("expected error %v, got %v", target, err)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func True(t *testing.T, val bool, msgAndArgs ...any) {
	t.Helper()
	if !val {
		if len(msgAndArgs) > 0 {
			t.Fatalf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		}
		t.Fatal("expected true, got false")
	}
}

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

func Len(t *testing.T, v any, expected int) {
	t.Helper()
	rv := reflect.ValueOf(v)
	if rv.Len() != expected {
		t.Fatalf("expected length %d, got %d", expected, rv.Len())
	}
}
