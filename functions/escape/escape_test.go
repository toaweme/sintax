package escape

import (
	"strings"
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// modifier is the shared signature of the three typed escapers.
type modifier func(any) (string, error)

// mustEscape runs an escaper and fails the test on error.
func mustEscape(t *testing.T, fn modifier, value any) string {
	t.Helper()
	out, err := fn(value)
	assert.NoError(t, err)
	return out
}

// xssVectors are payloads drawn from the OWASP XSS Filter Evasion cheat sheet
// and common real-world breakout attempts. Each escaper must neutralize every
// vector for the context it targets.
var xssVectors = []string{
	`<script>alert(1)</script>`,
	`<img src=x onerror=alert(1)>`,
	`"><script>alert(document.cookie)</script>`,
	`'><svg/onload=alert(1)>`,
	`<svg><script>alert&#40;1&#41;</script>`,
	`<a href="javascript:alert(1)">x</a>`,
	`<!--<script>alert(1)</script>-->`,
	`</textarea><script>alert(1)</script>`,
	"`-alert(1)-`",
	`</script><script>alert(1)</script>`,
}

func Test_HTML(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"ampersand", "&", "&amp;"},
		{"less than", "<", "&lt;"},
		{"greater than", ">", "&gt;"},
		{"double quote", `"`, "&#34;"},
		{"single quote", "'", "&#39;"},
		{"tag", "<b>hi</b>", "&lt;b&gt;hi&lt;/b&gt;"},
		{"mixed", `a<b>&"'`, `a&lt;b&gt;&amp;&#34;&#39;`},
		{"quoted attribute value", `a "quoted" word`, `a &#34;quoted&#34; word`},
		{"plain passes through", "hello world 123", "hello world 123"},
		{"unicode passes through", "café ☕ 日本", "café ☕ 日本"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, mustEscape(t, HTML, tt.value))
		})
	}
}

// Test_HTML_SinglePass proves the escaper does not double-encode. An
// already-escaped-looking input has only its literal "&" encoded once.
func Test_HTML_SinglePass(t *testing.T) {
	assert.Equal(t, "&amp;lt;", mustEscape(t, HTML, "&lt;"))
	assert.Equal(t, "&amp;amp;", mustEscape(t, HTML, "&amp;"))
}

// Test_HTML_NeutralizesXSS asserts no raw tag or quote delimiter survives, so a
// value dropped into an HTML text node or a quoted attribute cannot open a tag
// or break out of the surrounding quotes.
func Test_HTML_NeutralizesXSS(t *testing.T) {
	for _, v := range xssVectors {
		out := mustEscape(t, HTML, v)
		for _, bad := range []string{"<", ">", `"`, "'"} {
			if strings.Contains(out, bad) {
				t.Errorf("html escape of %q left raw %q: %q", v, bad, out)
			}
		}
	}
}

// Test_HTML_BoundaryUnquotedAttr locks in a documented limitation. html
// escaping is not enough for an unquoted attribute value, because spaces and
// "=" pass through, which leaves `<div class={{ v|escape_html }}>` injectable.
// If this ever changes, the package doc must change with it.
func Test_HTML_BoundaryUnquotedAttr(t *testing.T) {
	out := mustEscape(t, HTML, "x onerror=alert(1)")
	if !strings.Contains(out, " onerror=") {
		t.Fatalf("expected space and = to survive html escaping (documented boundary), got %q", out)
	}
}

// Test_HTML_BoundaryURLScheme locks in a documented limitation. html escaping
// does not neutralize a "javascript:" scheme, so it must not be relied on to
// make an href safe.
func Test_HTML_BoundaryURLScheme(t *testing.T) {
	assert.Equal(t, "javascript:alert(1)", mustEscape(t, HTML, "javascript:alert(1)"))
}

func Test_URL(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"space becomes plus", "a b", "a+b"},
		{"ampersand", "tea & coffee", "tea+%26+coffee"},
		{"reserved chars", "a=b&c=d", "a%3Db%26c%3Dd"},
		{"slashes and colon", "http://x/y", "http%3A%2F%2Fx%2Fy"},
		{"angle brackets", "<script>", "%3Cscript%3E"},
		{"unreserved untouched", "aZ0-_.~", "aZ0-_.~"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, mustEscape(t, URL, tt.value))
		})
	}
}

// Test_URL_NeutralizesXSS asserts no HTML-significant delimiter survives query
// escaping, so the result is also safe to drop into a quoted attribute.
func Test_URL_NeutralizesXSS(t *testing.T) {
	for _, v := range xssVectors {
		out := mustEscape(t, URL, v)
		for _, bad := range []string{"<", ">", `"`, "'", "&", " "} {
			if strings.Contains(out, bad) {
				t.Errorf("url escape of %q left raw %q: %q", v, bad, out)
			}
		}
	}
}

func Test_JS(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"double quote", "\"", "\\\""},
		{"single quote", "'", "\\'"},
		{"backslash", "\\", "\\\\"},
		{"forward slash", "/", "\\/"},
		{"backtick", "`", "\\u0060"},
		{"less than", "<", "\\u003C"},
		{"greater than", ">", "\\u003E"},
		{"ampersand", "&", "\\u0026"},
		{"tab", string(rune(9)), "\\t"},
		{"newline", string(rune(10)), "\\n"},
		{"carriage return", string(rune(13)), "\\r"},
		{"null", string(rune(0)), "\\u0000"},
		{"escape char", string(rune(27)), "\\u001B"},
		{"del", string(rune(127)), "\\u007F"},
		{"line separator", string(rune(0x2028)), "\\u2028"},
		{"paragraph separator", string(rune(0x2029)), "\\u2029"},
		{"close script tag", "</script>", "\\u003C\\/script\\u003E"},
		{"plain passes through", "hello world 123", "hello world 123"},
		{"unicode passes through", "café 日本", "café 日本"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, mustEscape(t, JS, tt.value))
		})
	}
}

// Test_JS_NeutralizesBreakout asserts nothing that could terminate a JS string
// literal or open/close a script or comment survives. Quotes are checked with
// backslash-parity, since a "\"" the escaper emits is a valid escaped quote that
// cannot close the string, whereas a quote after an even backslash run would.
// Every other significant rune is emitted only as a \u escape, so a raw
// occurrence of it is always a breakout.
func Test_JS_NeutralizesBreakout(t *testing.T) {
	neverRaw := []string{"`", "<", ">", "&", string(rune(10)), string(rune(13)), string(rune(0x2028)), string(rune(0x2029))}
	for _, v := range xssVectors {
		out := mustEscape(t, JS, v)
		for _, bad := range neverRaw {
			if strings.Contains(out, bad) {
				t.Errorf("js escape of %q left raw %q: %q", v, bad, out)
			}
		}
		if hasUnescapedQuote(out) {
			t.Errorf("js escape of %q left an unescaped quote: %q", v, out)
		}
		if strings.Contains(out, "</") || strings.Contains(strings.ToLower(out), "<script") {
			t.Errorf("js escape of %q left a tag close/open sequence: %q", v, out)
		}
	}
}

// hasUnescapedQuote reports whether s contains a single or double quote that is
// preceded by an even number of backslashes, i.e. one that would terminate a
// surrounding JS string literal.
func hasUnescapedQuote(s string) bool {
	backslashes := 0
	for _, r := range s {
		switch r {
		case '\\':
			backslashes++
			continue
		case '"', '\'':
			if backslashes%2 == 0 {
				return true
			}
		}
		backslashes = 0
	}
	return false
}

// Test_JS_AllControlChars checks every control character (bytes 0 to 31, e.g.
// null, tab, newline, escape) and the delete character (127). None may appear
// in the output as-is. Each must come out as a backslash escape.
func Test_JS_AllControlChars(t *testing.T) {
	for i := range 0x20 {
		r := rune(i)
		out := mustEscape(t, JS, string(r))
		if strings.ContainsRune(out, r) {
			t.Errorf("control char %#04x survived js escaping: %q", r, out)
		}
		if !strings.HasPrefix(out, "\\") {
			t.Errorf("control char %#04x not escaped: %q", r, out)
		}
	}
	out := mustEscape(t, JS, string(rune(0x7f)))
	if strings.ContainsRune(out, 0x7f) {
		t.Errorf("DEL survived js escaping: %q", out)
	}
}

// Test_Coercion covers the scalar coercion shared by all three escapers.
func Test_Coercion(t *testing.T) {
	assert.Equal(t, "", mustEscape(t, HTML, nil))
	assert.Equal(t, "42", mustEscape(t, HTML, 42))
	assert.Equal(t, "3.14", mustEscape(t, HTML, 3.14))
	assert.Equal(t, "true", mustEscape(t, HTML, true))
}

// Test_CompositeRejected covers the shared rejection of non-scalar values.
func Test_CompositeRejected(t *testing.T) {
	for _, fn := range []modifier{HTML, URL, JS} {
		_, err := fn([]int{1, 2})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)

		_, err = fn(map[string]int{"a": 1})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	}
}

// Test_Modifiers exercises the registered modifiers through the Wrap adapter, so
// the same string escaping is reached by template name. The adapter rejects a
// nil value and any params before the escaper body runs.
func Test_Modifiers(t *testing.T) {
	names := []functions.ModifierName{ModifierNameHTML, ModifierNameURL, ModifierNameJS}
	mods := Modifiers()
	for _, name := range names {
		fn := mods[string(name)]
		if fn == nil {
			t.Fatalf("modifier %q not registered", name)
		}
		out, err := fn("<a>", nil)
		assert.NoError(t, err)
		if s, ok := out.(string); !ok || s == "<a>" {
			t.Errorf("modifier %q did not escape, got %v", name, out)
		}

		// nil coerces to the any input and stringifies to "", matching the
		// direct HTML(nil) behaviour rather than rejecting.
		nilOut, err := fn(nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, "", nilOut)

		_, err = fn("x", []any{"extra"})
		assert.ErrorIs(t, err, functions.ErrInvalidParamType)
	}

	assert.Equal(t, "&lt;a&gt;", func() string {
		out, err := mods[string(ModifierNameHTML)]("<a>", nil)
		assert.NoError(t, err)
		return out.(string)
	}())
}
