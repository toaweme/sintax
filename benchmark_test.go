package sintax

import (
	"fmt"
	"testing"
)

// benchRender drives the full public path (parse + render) the way callers do.
// The engine instance is reused across iterations: it keeps no mutable state
// between calls (nesting depth is carried by per-call child renderers), so this
// measures steady-state render cost rather than construction. A warm-up call
// fails fast if the template/vars are wrong, so a benchmark never silently
// measures an error path.
func benchRender(b *testing.B, tmpl string, vars map[string]any) {
	b.Helper()
	s := New(builtins())
	if _, err := s.Render(tmpl, vars); err != nil {
		b.Fatalf("setup render failed: %v", err)
	}
	// report throughput as bytes of template text rendered per second
	b.SetBytes(int64(len(tmpl)))
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if _, err := s.Render(tmpl, vars); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Render_Variable(b *testing.B) {
	benchRender(b, `{{ name }}`, map[string]any{"name": "Ada"})
}

func Benchmark_Render_ModifierChain(b *testing.B) {
	benchRender(b, `{{ text | trim | upper }}`, map[string]any{"text": "  hello world  "})
}

// Benchmark_Render_JSONPipeline is the README headline example: parse JSON, dig
// into it, filter, project, reduce, format. It stresses the modifier-dispatch
// and value-coercion path more than the parser.
func Benchmark_Render_JSONPipeline(b *testing.B) {
	resp := `{"orders":[{"status":"paid","total":10.5},{"status":"pending","total":5},{"status":"paid","total":20.25}]}`
	benchRender(b,
		`{{ response | from_json | key:'orders' | filter:'status','paid' | pluck:'total' | sum | decimal:2 }}`,
		map[string]any{"response": resp},
	)
}

func Benchmark_Render_Loop(b *testing.B) {
	benchRender(b, "{{ for x in items }}- {{ x }}\n{{ endfor }}", map[string]any{"items": benchStrings(100)})
}

func Benchmark_Render_ConditionalLoop(b *testing.B) {
	benchRender(b, "{{ for x in flags }}{{ if x }}1{{ else }}0{{ endif }}{{ endfor }}", map[string]any{"flags": benchBools(100)})
}

// Benchmark_Render_LoopMap iterates a string-keyed map, exercising the key-sort
// path in renderFor.
func Benchmark_Render_LoopMap(b *testing.B) {
	benchRender(b, "{{ for k, v in items }}{{ k }}={{ v }}\n{{ endfor }}", map[string]any{"items": benchStringMap(100)})
}

// Benchmark_Render_TemplateModifier measures the nested-render primitive: the
// `template` modifier re-enters the engine to render a string variable.
func Benchmark_Render_TemplateModifier(b *testing.B) {
	benchRender(b, `{{ tpl | template }}`, map[string]any{
		"tpl":   "Hello {{ name }}, you have {{ count }} new messages",
		"name":  "Ada",
		"count": 7,
	})
}

// Benchmark_Render_TemplateModifier_Deep measures three levels of re-entrant
// rendering (each level allocates a child renderer + re-parses its source).
func Benchmark_Render_TemplateModifier_Deep(b *testing.B) {
	benchRender(b, `{{ l1 | template }}`, map[string]any{
		"l1":   "[{{ l2 | template }}]",
		"l2":   "({{ l3 | template }})",
		"l3":   "{{ leaf }}",
		"leaf": "deep",
	})
}

// benchComplexTemplate mixes every major feature: a modifier chain, a loop with
// per-item field access and first/last handling, and a terminal reduction
// pipeline. Used to compare parse-only vs render-only cost below.
const benchComplexTemplate = `{{ title | upper }}
{{ for o in orders }}- {{ o | key:'name' | trim }} x{{ o | key:'qty' }}{{ if o_last }}.{{ else }},{{ endif }}
{{ endfor }}
paid total: {{ orders | filter:'status','paid' | pluck:'total' | sum | decimal:2 }}`

func benchComplexVars() map[string]any {
	orders := make([]any, 0, 50)
	for i := range 50 {
		status := "paid"
		if i%3 == 0 {
			status = "pending"
		}
		orders = append(orders, map[string]any{
			"name":   fmt.Sprintf("  item-%d  ", i),
			"qty":    i + 1,
			"status": status,
			"total":  float64(i) + 0.5,
		})
	}
	return map[string]any{"title": "invoice", "orders": orders}
}

func Benchmark_Render_Complex(b *testing.B) {
	benchRender(b, benchComplexTemplate, benchComplexVars())
}

// Benchmark_Parse_Complex isolates tokenization cost (no rendering).
func Benchmark_Parse_Complex(b *testing.B) {
	p := NewStringParser()
	b.SetBytes(int64(len(benchComplexTemplate)))
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if _, err := p.Parse(benchComplexTemplate); err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark_RenderTokens_Complex isolates render cost by reusing tokens parsed
// once. Compared against Benchmark_Parse_Complex it shows the parse/render split
// for the same template.
func Benchmark_RenderTokens_Complex(b *testing.B) {
	tokens, err := NewStringParser().Parse(benchComplexTemplate)
	if err != nil {
		b.Fatalf("parse failed: %v", err)
	}
	r := NewTokenRenderer(builtins())
	vars := benchComplexVars()
	if _, err := r.Render(tokens, vars); err != nil {
		b.Fatalf("setup render failed: %v", err)
	}
	b.SetBytes(int64(len(benchComplexTemplate)))
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if _, err := r.Render(tokens, vars); err != nil {
			b.Fatal(err)
		}
	}
}

func benchStrings(n int) []any {
	xs := make([]any, n)
	for i := range xs {
		xs[i] = fmt.Sprintf("value-%d", i)
	}
	return xs
}

func benchBools(n int) []any {
	xs := make([]any, n)
	for i := range xs {
		xs[i] = i%2 == 0
	}
	return xs
}

func benchStringMap(n int) map[string]any {
	m := make(map[string]any, n)
	for i := range n {
		m[fmt.Sprintf("key-%03d", i)] = fmt.Sprintf("value-%d", i)
	}
	return m
}
