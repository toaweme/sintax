package functions

import (
	"strings"
	"testing"
)

// benchUpper is a plain string-only modifier standing in for the text-first
// modifiers AsText decorates (upper, lower, trim and friends). It is defined
// here rather than imported because the real ones live in functions/text/*,
// which import this package.
var benchUpper = Wrap(func(s string) (string, error) {
	return strings.ToUpper(s), nil
})

// benchUpperAsText is the same modifier as it is actually registered, wrapped so
// a scalar reaches it as text.
var benchUpperAsText = AsText(benchUpper)

// The AsText benchmarks below pair a bare modifier against its decorated form on
// the same value, so the delta is the decorator alone. Strings are the case that
// matters most. Every text modifier in the battery carries this wrapper, and a
// string value gains nothing from it, so whatever it costs there is a pure tax
// paid on the common path.
func Benchmark_AsText_String_Undecorated(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		if _, err := benchUpper("hello world", nil); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_AsText_String_Decorated(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		if _, err := benchUpperAsText("hello world", nil); err != nil {
			b.Fatal(err)
		}
	}
}

// An int has no undecorated sibling, since the bare modifier rejects it, which
// is the whole reason AsText exists. This measures what the conversion costs on
// the path it enables, where fmt.Sprint is expected to dominate.
func Benchmark_AsText_Int_Decorated(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		if _, err := benchUpperAsText(42, nil); err != nil {
			b.Fatal(err)
		}
	}
}

// A slice falls through Stringish untouched and the wrapped modifier rejects it.
// This is the cost of the decorator's type switch on a value it declines to
// convert, which is what every composite value pays.
func Benchmark_AsText_Slice_Decorated(b *testing.B) {
	value := []any{1, 2, 3}
	b.ReportAllocs()
	for range b.N {
		if _, err := benchUpperAsText(value, nil); err == nil {
			b.Fatal("expected the wrapped modifier to reject a slice")
		}
	}
}

// Stringish backs AsText and is exported for modifiers that stringify their
// params, so its per-kind cost is worth reading on its own.
func Benchmark_Stringish_String(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		if _, ok := Stringish("hello world"); !ok {
			b.Fatal("expected a string to be stringish")
		}
	}
}

func Benchmark_Stringish_Int(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		if _, ok := Stringish(42); !ok {
			b.Fatal("expected an int to be stringish")
		}
	}
}

func Benchmark_Stringish_Slice(b *testing.B) {
	value := []any{1, 2, 3}
	b.ReportAllocs()
	for range b.N {
		if _, ok := Stringish(value); ok {
			b.Fatal("expected a slice not to be stringish")
		}
	}
}
