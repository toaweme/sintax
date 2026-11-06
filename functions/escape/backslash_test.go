package escape

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

// android and apple are the sets the two platforms' string files escape.
const (
	android        = `\'"n`
	androidLeading = `@?`
	apple          = `\"nt`
)

func Test_Backslash(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		sets     []string
		expected string
	}{
		{name: "no sets", value: `a\b"c`, expected: `a\b"c`},
		{name: "android", value: "it's \"quoted\"\nnext \\ done", sets: []string{android, androidLeading}, expected: `it\'s \"quoted\"\nnext \\ done`},
		{name: "android leading at", value: "@string/x and @", sets: []string{android, androidLeading}, expected: `\@string/x and @`},
		{name: "android leading question", value: "?attr", sets: []string{android, androidLeading}, expected: `\?attr`},
		{name: "letters stay letters", value: "nature tour", sets: []string{android}, expected: "nature tour"},
		{name: "apple tab", value: "a\tb", sets: []string{apple}, expected: `a\tb`},
		{name: "raw control in set", value: "a\nb", sets: []string{"\n"}, expected: `a\nb`},
		{name: "format and emoji untouched", value: "%1$s 👋 你好", sets: []string{android, androidLeading}, expected: "%1$s 👋 你好"},
		{name: "number", value: 42, sets: []string{android}, expected: "42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Backslash(tt.value, tt.sets...)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Unbackslash(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		sets     []string
		expected string
	}{
		{name: "android", value: `it\'s \"q\"\nnext \\ done`, sets: []string{android}, expected: "it's \"q\"\nnext \\ done"},
		{name: "escaped backslash before n", value: `a\\nb`, sets: []string{android}, expected: `a\nb`},
		{name: "leading set joins", value: `\@x \?y`, sets: []string{android, androidLeading}, expected: "@x ?y"},
		{name: "outside the set is kept", value: `a\qb`, sets: []string{android}, expected: `a\qb`},
		{name: "trailing backslash", value: `a\`, sets: []string{android}, expected: `a\`},
		{name: "unicode", value: `café \U00E9`, sets: []string{"uU"}, expected: "café é"},
		{name: "surrogate pair", value: `👋!`, sets: []string{"u"}, expected: "👋!"},
		{name: "unicode off", value: `é`, sets: []string{android}, expected: `é`},
		{name: "short unicode", value: `\u00`, sets: []string{"u"}, expected: `\u00`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Unbackslash(tt.value, tt.sets...)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Backslash_RoundTrips(t *testing.T) {
	values := []string{
		`it's`, `"q"`, "line\nbreak", `back\slash`, `\n literal`, `\\n`, "@start", "?start", "%1$s and %2$d", "👋 你好 こんにちは", "",
	}
	for _, value := range values {
		t.Run(value, func(t *testing.T) {
			escaped, err := Backslash(value, android, androidLeading)
			assert.NoError(t, err)
			back, err := Unbackslash(escaped, android, androidLeading)
			assert.NoError(t, err)
			assert.Equal(t, value, back)
		})
	}
}

func Test_Backslash_Modifier(t *testing.T) {
	out, err := escapeBackslashModifier("it's @", []any{android, androidLeading})
	assert.NoError(t, err)
	assert.Equal(t, `it\'s @`, out)
	out, err = unescapeBackslashModifier(`it\'s`, []any{android})
	assert.NoError(t, err)
	assert.Equal(t, "it's", out)
}
