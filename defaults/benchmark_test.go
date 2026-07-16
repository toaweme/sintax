package defaults_test

import (
	"testing"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/defaults"
)

// All rebuilds every group's modifier map on each call, and New then merges them
// into the engine's own map. Neither is on a render path, but both are on the
// one-shot path below, so their cost is worth reading directly.
func Benchmark_Defaults_All(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		_ = defaults.All()
	}
}

func Benchmark_New_WithDefaults(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		_ = sintax.New(defaults.All())
	}
}

// The two Render benchmarks below are the same work reached two ways, so the
// delta is construction. sintax.Render builds a whole engine per call, which the
// batteries-included option makes expensive, and ExampleRender documents that
// path, so the gap against a reused engine should stay visible rather than
// surface as a puzzling profile in a caller that renders in a loop.
func Benchmark_Render_OneShot(b *testing.B) {
	vars := map[string]any{"name": "ada"}
	b.ReportAllocs()
	for range b.N {
		if _, err := sintax.Render(`{{ name | upper }}`, vars, defaults.All()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Render_ReusedEngine(b *testing.B) {
	s := sintax.New(defaults.All())
	vars := map[string]any{"name": "ada"}
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		if _, err := s.Render(`{{ name | upper }}`, vars); err != nil {
			b.Fatal(err)
		}
	}
}
