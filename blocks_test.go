package sintax

import (
	"errors"
	"slices"
	"testing"
)

func parseTokens(t *testing.T, template string) []Token {
	t.Helper()

	tokens, err := NewStringParser().Parse(template)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	return tokens
}

func Test_BlockEnd(t *testing.T) {
	tests := []struct {
		name     string
		template string
		start    int
		wantElse int
		wantEnd  int
		wantOK   bool
	}{
		{name: "for block", template: "{{ for a in xs }}{{ a }}{{ endfor }}", start: 0, wantElse: -1, wantEnd: 2, wantOK: true},
		{name: "nested for stops at the outer end", template: "{{ for a in xs }}{{ for b in a }}{{ endfor }}{{ endfor }}", start: 0, wantElse: -1, wantEnd: 3, wantOK: true},
		{name: "unterminated for", template: "{{ for a in xs }}{{ a }}", start: 0, wantElse: -1, wantEnd: -1},
		{name: "if block", template: "{{ if a }}x{{ endif }}", start: 0, wantElse: -1, wantEnd: 2, wantOK: true},
		{name: "if block with else", template: "{{ if a }}x{{ else }}y{{ endif }}", start: 0, wantElse: 2, wantEnd: 4, wantOK: true},
		{name: "nested else is not the outer else", template: "{{ if a }}{{ if b }}x{{ else }}y{{ endif }}{{ endif }}", start: 0, wantElse: -1, wantEnd: 6, wantOK: true},
		{name: "an else inside a for is not a block else", template: "{{ for a in xs }}{{ else }}{{ endfor }}", start: 0, wantElse: -1, wantEnd: 2, wantOK: true},
		{name: "unterminated if", template: "{{ if a }}x", start: 0, wantElse: -1, wantEnd: -1},
		{name: "a variable opens no block", template: "{{ a }}", start: 0, wantElse: -1, wantEnd: -1},
		{name: "an out of range start opens no block", template: "{{ a }}", start: 9, wantElse: -1, wantEnd: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := parseTokens(t, tt.template)

			elseIdx, endIdx, ok := blockEnd(tokens, tt.start, len(tokens))
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if elseIdx != tt.wantElse {
				t.Errorf("else index = %d, want %d", elseIdx, tt.wantElse)
			}
			if endIdx != tt.wantEnd {
				t.Errorf("end index = %d, want %d", endIdx, tt.wantEnd)
			}
		})
	}
}

// Test_ParseloopSpec pins the binding names one for tag puts in reach of its
// body, including the spaced pair form, which the parser normalises so nothing
// downstream has to trim the halves for itself.
func Test_ParseLoopSpec(t *testing.T) {
	tests := []struct {
		name         string
		template     string
		wantElement  string
		wantPosition string
		wantNames    []string
	}{
		{
			name:        "single binding",
			template:    "{{ for tx in txs }}{{ endfor }}",
			wantElement: "tx",
			wantNames:   []string{"tx", "tx_index", "tx_first", "tx_last", "tx_key"},
		},
		{
			name:         "paired binding",
			template:     "{{ for k,v in xs }}{{ endfor }}",
			wantElement:  "v",
			wantPosition: "k",
			wantNames:    []string{"k", "v", "v_index", "v_first", "v_last"},
		},
		{
			name:         "spaced paired binding",
			template:     "{{ for k , v in xs }}{{ endfor }}",
			wantElement:  "v",
			wantPosition: "k",
			wantNames:    []string{"k", "v", "v_index", "v_first", "v_last"},
		},
		{
			name:      "a tag naming nothing binds nothing",
			template:  "{{ for , in xs }}{{ endfor }}",
			wantNames: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loop := parseLoopSpec(parseTokens(t, tt.template)[0])
			if loop.Element != tt.wantElement {
				t.Errorf("element = %q, want %q", loop.Element, tt.wantElement)
			}
			if loop.Position != tt.wantPosition {
				t.Errorf("position = %q, want %q", loop.Position, tt.wantPosition)
			}

			names := make([]string, 0, len(loop.bindings()))
			for name := range loop.bindings() {
				names = append(names, name)
			}
			slices.Sort(names)
			want := slices.Clone(tt.wantNames)
			slices.Sort(want)
			if len(names) == 0 && len(want) == 0 {
				return
			}
			if !slices.Equal(names, want) {
				t.Errorf("names = %v, want %v", names, want)
			}
		})
	}
}

// Test_ParseLoopSpec_MatchesRenderedBindings checks that the names the spec
// reports are the names the renderer publishes.
func Test_ParseLoopSpec_MatchesRenderedBindings(t *testing.T) {
	const template = "{{ for k , v in xs }}{{ k }}:{{ v }}:{{ v_index }}:{{ v_first }}:{{ v_last }};{{ endfor }}"

	got, err := RenderString(template, map[string]any{"xs": []any{"a", "b"}})
	if err != nil {
		t.Fatalf("failed to render template: %v", err)
	}
	if want := "0:a:0:true:false;1:b:1:false:true;"; got != want {
		t.Errorf("render = %q, want %q", got, want)
	}
}

func Test_ParseExpr(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		wantTyp TokenType
		wantVar string
		wantErr error
	}{
		{name: "bare variable", expr: "user", wantTyp: VariableToken, wantVar: "user"},
		{name: "dotted variable", expr: " user.name ", wantTyp: VariableToken, wantVar: "user.name"},
		{name: "filtered variable", expr: "rows | sort", wantTyp: FilteredVariableToken, wantVar: "rows"},
		{name: "literal", expr: "1 + 2", wantErr: ErrNotAnExpression},
		{name: "control tag", expr: "endfor", wantErr: ErrNotAnExpression},
		{name: "empty", expr: "  ", wantErr: ErrNotAnExpression},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := parseExpr(tt.expr)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("error = %v, want %v", err, tt.wantErr)
				}

				return
			}
			if err != nil {
				t.Fatalf("failed to parse expression: %v", err)
			}
			if token.Type() != tt.wantTyp {
				t.Errorf("type = %v, want %v", token.Type(), tt.wantTyp)
			}
			if token.Name() != tt.wantVar {
				t.Errorf("name = %q, want %q", token.Name(), tt.wantVar)
			}
		})
	}
}
