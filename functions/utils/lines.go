package utils

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameLineNumbers is the template name for the LineNumbers modifier.
const ModifierNameLineNumbers functions.ModifierName = "line-numbers"

// LineNumbers prepends each line of the string with its zero-based line number.
//
// value: string
// returns: string
//
// example: number lines of a short note
// in:  note = "Buy milk\nWalk the dog\nPay rent"
// tpl: {{ note | line-numbers }}
// out: 0. Buy milk
// out: 1. Walk the dog
// out: 2. Pay rent
//
// example: number a haiku
// in:  poem = "An old silent pond\nA frog jumps into the pond\nSplash! Silence again."
// tpl: {{ poem | line-numbers }}
// out: 0. An old silent pond
// out: 1. A frog jumps into the pond
// out: 2. Splash! Silence again.
func LineNumbers(value any, params []any) (any, error) {
	if value == nil {
		return nil, nil
	}
	if value == "" {
		return nil, nil
	}

	input, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("line_numbers function expected string, got %T", value)
	}

	result := addLineNumbers(input)

	return result, nil
}

func addLineNumbers(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = fmt.Sprintf("%d. %s", i, line)
	}

	return strings.Join(lines, "\n")
}
