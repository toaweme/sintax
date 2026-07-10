package boolean

import (
	"errors"
	"testing"

	"github.com/toaweme/sintax/functions"
)

// The *Baseline functions below reproduce the pre-migration untyped bodies, so a
// run measures the Wrap / Overload coercion overhead against the exact code the
// typed clauses replaced rather than against an abstract ideal. Compare a
// Baseline benchmark with its Wrapped / Overload sibling to read the delta.

var errBenchMissing = errors.New("requires at least one parameter")

func gtBaseline(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errBenchMissing
	}
	val, err := functions.ValueNumber(value)
	if err != nil {
		return nil, err
	}
	than, err := functions.ValueNumber(params[0])
	if err != nil {
		return nil, err
	}
	return val > than, nil
}

func eqBaseline(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errBenchMissing
	}
	other := params[0]
	if value == nil || other == nil {
		return value == other, nil
	}
	valNum, errVal := functions.ValueNumber(value)
	otherNum, errOther := functions.ValueNumber(other)
	if errVal == nil && errOther == nil {
		return valNum == otherNum, nil
	}
	valStr, okVal := value.(string)
	otherStr, okOther := other.(string)
	if okVal && okOther {
		return valStr == otherStr, nil
	}
	return value == other, nil
}

func Benchmark_Gt_Baseline(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		_, _ = gtBaseline(91, []any{90.5})
	}
}

func Benchmark_Gt_Wrapped(b *testing.B) {
	gt := gtModifier
	b.ReportAllocs()
	for range b.N {
		_, _ = gt(91, []any{90.5})
	}
}

func Benchmark_Eq_Baseline_Number(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		_, _ = eqBaseline(5, []any{5})
	}
}

// Benchmark_Eq_Overload_Number is the best case: the nil guard declines and the
// first typed clause (EqNumber) accepts.
func Benchmark_Eq_Overload_Number(b *testing.B) {
	eq := eqModifier
	b.ReportAllocs()
	for range b.N {
		_, _ = eq(5, []any{5})
	}
}

func Benchmark_Eq_Baseline_String(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		_, _ = eqBaseline("active", []any{"active"})
	}
}

// Benchmark_Eq_Overload_String is a mid case: the nil guard declines, EqNumber
// rejects on a failed coercion, and EqString accepts on the second clause.
func Benchmark_Eq_Overload_String(b *testing.B) {
	eq := eqModifier
	b.ReportAllocs()
	for range b.N {
		_, _ = eq("active", []any{"active"})
	}
}

func Benchmark_Eq_Baseline_Any(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		_, _ = eqBaseline(true, []any{true})
	}
}

// Benchmark_Eq_Overload_Any is the worst case: every earlier clause rejects and
// only EqAny, listed last, accepts, so the value falls through the full chain.
func Benchmark_Eq_Overload_Any(b *testing.B) {
	eq := eqModifier
	b.ReportAllocs()
	for range b.N {
		_, _ = eq(true, []any{true})
	}
}
