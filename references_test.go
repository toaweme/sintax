package sintax

import (
	"errors"
	"testing"
)

func Test_References(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     []Reference
		wantErr  error
	}{
		{
			name:     "a plain variable",
			template: "{{ name }}",
			want:     []Reference{{Name: "name", Root: "name"}},
		},
		{
			name:     "a pipeline ending in a default answers its own miss",
			template: "{{ name | upper | default:'x' }}",
			want:     []Reference{{Name: "name", Root: "name", Default: true}},
		},
		{
			name:     "a default inside a quoted argument is not a default",
			template: "{{ name | replace:'|default','x' }}",
			want:     []Reference{{Name: "name", Root: "name"}},
		},
		{
			name:     "a modifier argument is a reference with no default of its own",
			template: "{{ name | default:fallback }}",
			want:     []Reference{{Name: "name", Root: "name", Default: true}, {Name: "fallback", Root: "fallback"}},
		},
		{
			name:     "a loop reads its collection and binds its own names",
			template: "{{ for row in rows }}{{ row.title }}{{ row_index }}{{ endfor }}",
			want:     []Reference{{Name: "rows", Root: "rows"}},
		},
		{
			name:     "a paired loop binds the position it named and not the map key",
			template: "{{ for k, v in xs }}{{ k }}{{ v }}{{ v_key }}{{ endfor }}",
			want:     []Reference{{Name: "xs", Root: "xs"}, {Name: "v_key", Root: "v_key"}},
		},
		{
			name:     "an unpaired loop binds the map key",
			template: "{{ for v in xs }}{{ v_key }}{{ endfor }}",
			want:     []Reference{{Name: "xs", Root: "xs"}},
		},
		{
			name:     "a nested loop keeps the outer bindings in reach",
			template: "{{ for a in as }}{{ for b in a.bs }}{{ a }}{{ b }}{{ endfor }}{{ endfor }}",
			want:     []Reference{{Name: "as", Root: "as"}},
		},
		{
			name:     "an if condition is optional and its body is not",
			template: "{{ if ready }}{{ payload }}{{ endif }}",
			want:     []Reference{{Name: "ready", Root: "ready", Optional: true}, {Name: "payload", Root: "payload"}},
		},
		{
			name:     "an else branch is read like the body",
			template: "{{ if ready }}{{ a }}{{ else }}{{ b }}{{ endif }}",
			want:     []Reference{{Name: "ready", Root: "ready", Optional: true}, {Name: "a", Root: "a"}, {Name: "b", Root: "b"}},
		},
		{
			name:     "a block nothing closes opens no block",
			template: "{{ for row in rows }}{{ row }}",
			want:     []Reference{{Name: "row", Root: "row"}},
		},
		{
			name:     "text carries no reference",
			template: "hello there",
		},
		{
			name:     "a condition that is not an expression is refused",
			template: "{{ if 'literal' }}{{ endif }}",
			wantErr:  ErrNotAnExpression,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := parseTokens(t, tt.template)
			got, err := References(tokens)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("error = %v, want %v", err, tt.wantErr)
				}

				return
			}
			if err != nil {
				t.Fatalf("failed to read references: %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("references = %v, want %v", got, tt.want)
			}
			for i, ref := range got {
				if ref != tt.want[i] {
					t.Errorf("reference %d = %+v, want %+v", i, ref, tt.want[i])
				}
			}
		})
	}
}
