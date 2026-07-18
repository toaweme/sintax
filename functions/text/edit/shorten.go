package edit

import (
	"fmt"
	"strconv"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameShorten is the template name for the Shorten modifier.
const ModifierNameShorten functions.ModifierName = "shorten"

// Shorten truncates the input to at most a given number of bytes, returning it
// unchanged when it is already shorter. The limit counts bytes, not characters,
// so cutting inside a multi-byte character can leave a broken final byte.
func Shorten(s string, length int) (string, error) {
	if len(s) > length {
		return s[:length], nil
	}
	return s, nil
}

// ShortenParse is the shorten clause for a numeric-string length such as
// shorten:'30', parsing it before delegating to Shorten. A non-numeric length is
// an error.
func ShortenParse(s, length string) (string, error) {
	n, err := strconv.Atoi(length)
	if err != nil {
		return "", fmt.Errorf("failed to parse shorten length %q: %w", length, err)
	}
	return Shorten(s, n)
}
