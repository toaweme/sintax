package sintax

import (
	"errors"
	"strings"
	"testing"

	"github.com/toaweme/sintax/functions"
)

// A modifier that turns a value into a decision refuses nil. A key that is not
// there and a key holding an explicit null are the same value by the time a
// modifier sees them, so the two are driven through the same table to pin them as
// equivalent. Without this, `{{ row.reviews_enabled | not }}` over a misspelled or
// null column answers true and every row imports as reviewed.
func Test_Render_NilDecision_Refuses(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		modifier string
	}{
		{
			name:     "not inverts nothing",
			template: `{{ row.reviews_enabled | not }}`,
			modifier: "not",
		},
		{
			name:     "not feeding or",
			template: `{{ row.reviews_enabled | not | or:approver }}`,
			modifier: "not",
		},
		{
			name:     "when picks a branch on nothing",
			template: `{{ row.reviews_enabled | when:'approved','in_review' }}`,
			modifier: "when",
		},
		{
			name:     "and reads nothing as false",
			template: `{{ row.reviews_enabled | and:approver }}`,
			modifier: "and",
		},
		{
			name:     "or reads nothing as false",
			template: `{{ row.reviews_enabled | or:approver }}`,
			modifier: "or",
		},
		{
			name:     "eq compares nothing to a value",
			template: `{{ row.reviews_enabled | eq:'published' }}`,
			modifier: "eq",
		},
		{
			name:     "neq compares nothing to a value",
			template: `{{ row.reviews_enabled | neq:'published' }}`,
			modifier: "neq",
		},
		{
			name:     "has looks inside nothing",
			template: `{{ row.reviews_enabled | has:'x' }}`,
			modifier: "has",
		},
		{
			name:     "is tests nothing for membership",
			template: `{{ row.reviews_enabled | is:'x' }}`,
			modifier: "is",
		},
		{
			name:     "a decision inside an if condition",
			template: `{{ if row.reviews_enabled | not }}yes{{ else }}no{{ endif }}`,
			modifier: "not",
		},
	}

	// the two ways a boolean column arrives as nothing. Neither reaches a
	// modifier as anything but nil, so both must fail the same way.
	valueSets := map[string]map[string]any{
		"missing key":   {"row": map[string]any{}, "approver": "ada"},
		"explicit null": {"row": map[string]any{"reviews_enabled": nil}, "approver": "ada"},
	}

	for _, tt := range testCases {
		for setName, vars := range valueSets {
			t.Run(tt.name+", "+setName, func(t *testing.T) {
				s := New(builtins())

				_, err := s.Render(tt.template, vars)
				if err == nil {
					t.Fatalf("rendered a decision on nothing without failing")
				}
				if !errors.Is(err, functions.ErrNilDecision) {
					t.Fatalf("got %v, want a nil decision refusal", err)
				}
				var failure *ModifierError
				if !errors.As(err, &failure) {
					t.Fatalf("got %v, want a modifier failure naming the modifier", err)
				}
				if failure.Modifier != tt.modifier {
					t.Fatalf("got modifier %q, want %q", failure.Modifier, tt.modifier)
				}
				if !strings.Contains(err.Error(), "row.reviews_enabled") {
					t.Fatalf("got %v, want the path named", err)
				}
			})
		}
	}
}

// A default ahead of the decision is how a template says what absence means, and
// it answers both shapes of nothing. Written after the decision it is too late,
// since the decision already failed.
func Test_Render_NilDecision_AnsweredByAnEarlierDefault(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		want     any
		wantErr  bool
	}{
		{
			name:     "a default ahead of not",
			template: `{{ row.reviews_enabled | default:true | not }}`,
			want:     false,
		},
		{
			name:     "a default ahead of when",
			template: `{{ row.reviews_enabled | default:false | when:'approved','in_review' }}`,
			want:     "in_review",
		},
		{
			name:     "a default ahead of eq",
			template: `{{ row.status | default:'draft' | eq:'published' }}`,
			want:     false,
		},
		{
			name:     "a default after the decision is too late",
			template: `{{ row.reviews_enabled | not | default:false }}`,
			wantErr:  true,
		},
	}

	valueSets := map[string]map[string]any{
		"missing key":   {"row": map[string]any{}},
		"explicit null": {"row": map[string]any{"reviews_enabled": nil, "status": nil}},
	}

	for _, tt := range testCases {
		for setName, vars := range valueSets {
			t.Run(tt.name+", "+setName, func(t *testing.T) {
				s := New(builtins())

				out, err := s.Render(tt.template, vars)
				if tt.wantErr {
					if !errors.Is(err, functions.ErrNilDecision) {
						t.Fatalf("got %v, want a nil decision refusal", err)
					}
					return
				}
				if err != nil {
					t.Fatalf("failed to render a decision behind a default: %v", err)
				}
				if out != tt.want {
					t.Fatalf("got %v, want %v", out, tt.want)
				}
			})
		}
	}
}

// Reading a value stays soft. Only a decision on nothing refuses, so a missing
// field still renders through a default and a bare miss in a condition is still
// false.
func Test_Render_NilDecision_LeavesReadingAlone(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		want     any
	}{
		{
			name:     "a missing field behind a default",
			template: `{{ sub.body.currency | default:'' }}`,
			want:     "",
		},
		{
			name:     "a bare miss in a condition is false",
			template: `{{ if sub.body.currency }}yes{{ else }}no{{ endif }}`,
			want:     "no",
		},
		{
			name:     "a missing iterable yields no iterations",
			template: `{{ for x in sub.body.rows }}{{ x }}{{ endfor }}done`,
			want:     "done",
		},
	}

	vars := map[string]any{"sub": map[string]any{"body": map[string]any{}}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			out, err := s.Render(tt.template, vars)
			if err != nil {
				t.Fatalf("failed to render a value read: %v", err)
			}
			if out != tt.want {
				t.Fatalf("got %v, want %v", out, tt.want)
			}
		})
	}
}
