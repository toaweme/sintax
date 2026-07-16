package boolean

import (
	"errors"
	"testing"

	"github.com/toaweme/sintax/functions"
)

// The *Baseline functions below are hand-written untyped modifier bodies, so a
// run measures the Wrap / Overload coercion overhead against equivalent code
// doing the same work by hand rather than against an abstract ideal. Compare a
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

// Each baseline is called through a GlobalModifier value rather than directly,
// and every param slice is hoisted out of its loop. Both details exist to keep
// the comparison honest. A direct call to a known function lets escape analysis
// stack-allocate an inline []any{...} literal, while a call through a func value
// is opaque and forces the same literal onto the heap, so a benchmark that calls
// the baseline directly and the adapter indirectly reports that difference as an
// allocation the adapter added. The engine only ever reaches a modifier through a
// GlobalModifier value anyway, so the indirection is what production does.
var (
	gtBaselineMod functions.GlobalModifier = gtBaseline
	eqBaselineMod functions.GlobalModifier = eqBaseline

	benchGtParams     = []any{90.5}
	benchNumberParams = []any{5}
	benchStringParams = []any{"active"}
	benchAnyParams    = []any{true}
)

func Benchmark_Gt_Baseline(b *testing.B) {
	gt := gtBaselineMod
	b.ReportAllocs()
	for range b.N {
		_, _ = gt(91, benchGtParams)
	}
}

func Benchmark_Gt_Wrapped(b *testing.B) {
	gt := gtModifier
	b.ReportAllocs()
	for range b.N {
		_, _ = gt(91, benchGtParams)
	}
}

func Benchmark_Eq_Baseline_Number(b *testing.B) {
	eq := eqBaselineMod
	b.ReportAllocs()
	for range b.N {
		_, _ = eq(5, benchNumberParams)
	}
}

// Benchmark_Eq_Overload_Number is the best case, where the nil guard declines and
// the first typed clause (EqNumber) accepts.
func Benchmark_Eq_Overload_Number(b *testing.B) {
	eq := eqModifier
	b.ReportAllocs()
	for range b.N {
		_, _ = eq(5, benchNumberParams)
	}
}

func Benchmark_Eq_Baseline_String(b *testing.B) {
	eq := eqBaselineMod
	b.ReportAllocs()
	for range b.N {
		_, _ = eq("active", benchStringParams)
	}
}

// Benchmark_Eq_Overload_String is a mid case, where the nil guard declines,
// EqNumber rejects on a failed coercion, and EqString accepts on the second
// clause.
func Benchmark_Eq_Overload_String(b *testing.B) {
	eq := eqModifier
	b.ReportAllocs()
	for range b.N {
		_, _ = eq("active", benchStringParams)
	}
}

func Benchmark_Eq_Baseline_Any(b *testing.B) {
	eq := eqBaselineMod
	b.ReportAllocs()
	for range b.N {
		_, _ = eq(true, benchAnyParams)
	}
}

// Benchmark_Eq_Overload_Any is the worst case, where every earlier clause rejects
// and only EqAny, listed last, accepts, so the value falls through the full
// chain.
func Benchmark_Eq_Overload_Any(b *testing.B) {
	eq := eqModifier
	b.ReportAllocs()
	for range b.N {
		_, _ = eq(true, benchAnyParams)
	}
}
