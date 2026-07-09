package sintax

import (
	"os"
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Banking_Template(t *testing.T) {
	want, err := os.ReadFile("_data/templates/banking.xml")
	assert.NoError(t, err)

	type amounts struct {
		a1, a2, a3 any
	}

	mkBasic := func(am amounts) map[string]any {
		return map[string]any{
			"msg_id":         "-20260409161255",
			"cre_dt_tm":      "2026-04-09T16:12:55",
			"total_txs":      "3",
			"total_ctrl_sum": "28596.05",
			"debtor_name":    "ACME Holdings, MB",
			"debtor_iban":    "LT857300010000000001",
			"debtor_bic":     "HABALT22",
			"groups": []any{
				map[string]any{
					"pmt_inf_id": "SEPA-2026-04-10-0",
					"is_sepa":    true,
					"exec_date":  "2026-04-10",
					"nb_of_txs":  "3",
					"ctrl_sum":   "28596.05",
					"txs":        mkTxs(am.a1, am.a2, am.a3),
				},
			},
		}
	}
	mkAdvanced := func(am amounts) map[string]any {
		return map[string]any{
			"msg_id":      "-20260409161255",
			"cre_dt_tm":   "2026-04-09T16:12:55",
			"debtor_name": "ACME Holdings, MB",
			"debtor_iban": "LT857300010000000001",
			"debtor_bic":  "HABALT22",
			"groups": []any{
				map[string]any{
					"pmt_inf_id": "SEPA-2026-04-10-0",
					"is_sepa":    true,
					"exec_date":  "2026-04-10",
					"txs":        mkTxs(am.a1, am.a2, am.a3),
				},
			},
		}
	}

	funcs := builtins()

	cases := []struct {
		name     string
		template string
		vars     map[string]any
	}{
		{
			name:     "basic template with string amounts",
			template: "_data/templates/banking.tpl.xml",
			vars:     mkBasic(amounts{"4235.00", "14070.00", "10291.05"}),
		},
		{
			name:     "advanced template with float64 amounts",
			template: "_data/templates/banking-advanced.tpl.xml",
			vars:     mkAdvanced(amounts{4235.00, 14070.00, 10291.05}),
		},
		{
			name:     "advanced template with int and float amounts",
			template: "_data/templates/banking-advanced.tpl.xml",
			vars:     mkAdvanced(amounts{int(4235), int(14070), 10291.05}),
		},
		{
			name:     "advanced template with mixed numeric types",
			template: "_data/templates/banking-advanced.tpl.xml",
			vars:     mkAdvanced(amounts{int64(4235), float32(14070.00), 10291.05}),
		},
		{
			name:     "advanced template with string amounts (parsed by | decimal)",
			template: "_data/templates/banking-advanced.tpl.xml",
			vars:     mkAdvanced(amounts{"4235.00", "14070.00", "10291.05"}),
		},
		{
			name:     "pretty template uses {{- -}} trim markers",
			template: "_data/templates/banking-pretty.tpl.xml",
			vars:     mkAdvanced(amounts{4235.00, 14070.00, 10291.05}),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tpl, err := os.ReadFile(tc.template)
			assert.NoError(t, err)

			got, err := Render(string(tpl), tc.vars, funcs)
			assert.NoError(t, err)

			gotStr, ok := got.(string)
			assert.True(t, ok, "expected string, got %T", got)

			assert.Equal(t, string(want), gotStr)
		})
	}
}

func mkTxs(a1, a2, a3 any) []any {
	return []any{
		map[string]any{
			"e2e_id":          "20260409-1",
			"is_sepa":         true,
			"currency":        "EUR",
			"amount":          a1,
			"creditor_bic":    "",
			"creditor_name":   "ACME Consulting MB",
			"creditor_iban":   "LT777300010000000002",
			"remittance_info": "INV-001",
		},
		map[string]any{
			"e2e_id":          "20260409-2",
			"is_sepa":         true,
			"currency":        "EUR",
			"amount":          a2,
			"creditor_bic":    "",
			"creditor_name":   "Wide Foreign Sp. z o.o.",
			"creditor_iban":   "PL861050150000000000000003",
			"remittance_info": "INV-002",
		},
		map[string]any{
			"e2e_id":          "20260409-3",
			"is_sepa":         true,
			"currency":        "EUR",
			"amount":          a3,
			"creditor_bic":    "",
			"creditor_name":   "Jane Doe",
			"creditor_iban":   "LT357300010000000004",
			"remittance_info": "INV-003",
		},
	}
}

func Test_Sum_Pluck_Flatten_Decimal(t *testing.T) {
	funcs := builtins()

	t.Run("sum slice of numbers", func(t *testing.T) {
		out, err := Render("{{ xs | sum }}", map[string]any{"xs": []any{1.5, 2.5, 3.0}}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, float64(7), out)
	})

	t.Run("sum field across slice of maps", func(t *testing.T) {
		out, err := Render("{{ items | sum:'price' | decimal:2 }}", map[string]any{
			"items": []any{
				map[string]any{"price": 10.0},
				map[string]any{"price": 20.5},
				map[string]any{"price": 0.25},
			},
		}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "30.75", out)
	})

	t.Run("pluck", func(t *testing.T) {
		out, err := Render("{{ items | pluck:'name' | length }}", map[string]any{
			"items": []any{
				map[string]any{"name": "a"},
				map[string]any{"name": "b"},
			},
		}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, 2, out)
	})

	t.Run("flatten", func(t *testing.T) {
		out, err := Render("{{ xs | flatten | sum }}", map[string]any{
			"xs": []any{
				[]any{1.0, 2.0},
				[]any{3.0, 4.0},
			},
		}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, float64(10), out)
	})

	t.Run("pluck flatten sum field", func(t *testing.T) {
		out, err := Render("{{ groups | pluck:'txs' | flatten | sum:'amount' | decimal:2 }}", map[string]any{
			"groups": []any{
				map[string]any{"txs": []any{
					map[string]any{"amount": 1.5},
					map[string]any{"amount": 2.5},
				}},
				map[string]any{"txs": []any{
					map[string]any{"amount": 100.0},
				}},
			},
		}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "104.00", out)
	})

	t.Run("decimal default 2 places", func(t *testing.T) {
		out, err := Render("{{ x | decimal }}", map[string]any{"x": 1.2}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "1.20", out)
	})
}

func Test_For_Loop(t *testing.T) {
	funcs := builtins()

	t.Run("slice single var", func(t *testing.T) {
		out, err := Render("{{ for x in xs }}- {{ x }}\n{{ endfor }}", map[string]any{
			"xs": []any{"a", "b", "c"},
		}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "- a\n- b\n- c\n", out)
	})

	t.Run("slice index, value", func(t *testing.T) {
		out, err := Render("{{ for i, x in xs }}{{ i }}:{{ x }} {{ endfor }}", map[string]any{
			"xs": []any{"a", "b", "c"},
		}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "0:a 1:b 2:c ", out)
	})

	t.Run("map key, value (sorted)", func(t *testing.T) {
		out, err := Render("{{ for k, v in m }}{{ k }}={{ v }};{{ endfor }}", map[string]any{
			"m": map[string]any{"b": "2", "a": "1", "c": "3"},
		}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "a=1;b=2;c=3;", out)
	})

	t.Run("nested for", func(t *testing.T) {
		out, err := Render("{{ for row in rows }}{{ for v in row }}{{ v }}{{ endfor }}|{{ endfor }}", map[string]any{
			"rows": []any{
				[]any{"a", "b"},
				[]any{"c", "d"},
			},
		}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "ab|cd|", out)
	})
}

func Test_If_Else(t *testing.T) {
	funcs := builtins()

	t.Run("if true", func(t *testing.T) {
		out, err := Render("{{ if x }}yes{{ endif }}", map[string]any{"x": true}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "yes", out)
	})

	t.Run("if false (no else)", func(t *testing.T) {
		out, err := Render("a{{ if x }}yes{{ endif }}b", map[string]any{"x": false}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "ab", out)
	})

	t.Run("if/else", func(t *testing.T) {
		out, err := Render("{{ if x }}yes{{ else }}no{{ endif }}", map[string]any{"x": false}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "no", out)
	})

	t.Run("empty string is falsy", func(t *testing.T) {
		out, err := Render("{{ if name }}{{ name }}{{ else }}anon{{ endif }}", map[string]any{"name": ""}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "anon", out)
	})

	t.Run("non-empty string is truthy", func(t *testing.T) {
		out, err := Render("{{ if name }}hi {{ name }}{{ endif }}", map[string]any{"name": "ada"}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "hi ada", out)
	})

	t.Run("filtered condition", func(t *testing.T) {
		out, err := Render("{{ if active | not }}inactive{{ endif }}", map[string]any{"active": false}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "inactive", out)
	})

	t.Run("if inside for", func(t *testing.T) {
		out, err := Render("{{ for x in xs }}{{ if x }}1{{ else }}0{{ endif }}{{ endfor }}", map[string]any{
			"xs": []any{true, false, true, true, false},
		}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "10110", out)
	})
}

func Test_Whitespace_Trim(t *testing.T) {
	funcs := builtins()

	t.Run("control tag alone on line auto-trims", func(t *testing.T) {
		input := "PRE\n{{ if x }}\nyes\n{{ endif }}\nPOST"
		out, err := Render(input, map[string]any{"x": true}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "PRE\nyes\nPOST", out)
	})

	t.Run("for loop alone on line auto-trims", func(t *testing.T) {
		input := "PRE\n{{ for x in xs }}\n- {{ x }}\n{{ endfor }}\nPOST"
		out, err := Render(input, map[string]any{"xs": []any{"a", "b"}}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "PRE\n- a\n- b\nPOST", out)
	})

	t.Run("explicit -}} strips following whitespace", func(t *testing.T) {
		out, err := Render("a{{ x -}}    b", map[string]any{"x": "X"}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "aXb", out)
	})

	t.Run("explicit {{- strips preceding whitespace", func(t *testing.T) {
		out, err := Render("a   {{- x }}b", map[string]any{"x": "X"}, funcs)
		assert.NoError(t, err)
		assert.Equal(t, "aXb", out)
	})
}
