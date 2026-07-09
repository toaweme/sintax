// Package escape provides modifiers that make a string safe to embed in a
// target output context, one modifier per context (HTML, URL, JavaScript).
//
// Escaping here is deliberately explicit, unlike html/template's contextual
// auto-escaping. html/template inspects where a value lands (an HTML text node,
// an attribute, a URL, a script) and picks the right escaper automatically.
// These modifiers cannot see the surrounding markup, so the template author
// picks the modifier for where the value lands. The author stays responsible
// for two things the engine cannot check. One is picking the modifier that
// matches where the value actually lands. The other is keeping the value inside
// quotes when the modifier assumes it, which escape_html (for attributes) and
// escape_js do. These modifiers are the counterpart to the convert modifiers.
// convert changes a value's representation, while escape leaves the text as
// text and only defuses the characters the destination would otherwise
// interpret. See each modifier for its safety boundary.
package escape

import (
	"fmt"
	"html"
	"net/url"
	"strconv"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// Modifier names, one per output context.
const (
	// ModifierNameHTML is the template name for the HTML modifier.
	ModifierNameHTML functions.ModifierName = "escape_html"
	// ModifierNameURL is the template name for the URL modifier.
	ModifierNameURL functions.ModifierName = "escape_url"
	// ModifierNameJS is the template name for the JS modifier.
	ModifierNameJS functions.ModifierName = "escape_js"
)

// HTML escapes a value for an HTML text node or a quoted attribute value. It is
// not safe for an unquoted attribute, an href or src URL, or content inside
// <script>, <style>, or an HTML comment. Reach for escape_js, escape_url, or a
// dedicated sanitizer in those places.
//
// value: string (scalars are coerced to their string form)
// returns: string
//
// example: escape a comment before dropping it into HTML
// in:  comment = "<b>hi</b>"
// tpl: {{ comment | escape_html }}
// out: &lt;b&gt;hi&lt;/b&gt;
func HTML(value any, _ []any) (any, error) {
	str, err := stringify(value)
	if err != nil {
		return nil, err
	}
	return html.EscapeString(str), nil
}

// URL escapes a value for use as a query-string value. It does not validate URL
// schemes, so it never turns an attacker-controlled string into a safe href on
// its own. A "javascript:" payload stays a live scheme if you drop the whole URL
// into href unescaped.
//
// value: string (scalars are coerced to their string form)
// returns: string
//
// example: escape a search term for a query string
// in:  term = "tea & coffee"
// tpl: https://example.com/s?q={{ term | escape_url }}
// out: https://example.com/s?q=tea+%26+coffee
func URL(value any, _ []any) (any, error) {
	str, err := stringify(value)
	if err != nil {
		return nil, err
	}
	return url.QueryEscape(str), nil
}

// JS escapes a value for use inside a quoted JavaScript string literal. It is
// not safe as bare JavaScript, because escaping leaves a ";", a "(", or an
// identifier untouched, so a value spliced outside a string literal is still
// live code.
//
// value: string (scalars are coerced to their string form)
// returns: string
//
// example: escape a value embedded in a script literal
// in:  name = "a\"; drop()"
// tpl: <script>var n = "{{ name | escape_js }}";</script>
// out: <script>var n = "a\"; drop()";</script>
func JS(value any, _ []any) (any, error) {
	str, err := stringify(value)
	if err != nil {
		return nil, err
	}
	return escapeJS(str), nil
}

// stringify coerces a scalar value to the string form the escapers will encode.
// nil becomes the empty string, and the numeric and boolean kinds are formatted
// directly. Composite values (slices, maps, structs) are rejected. Escaping
// their Go rendering is almost always a mistake, so it surfaces as an error
// rather than silently producing garbage that merely looks escaped.
func stringify(v any) (string, error) {
	switch s := v.(type) {
	case nil:
		return "", nil
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return fmt.Sprintf("%v", s), nil
	default:
		return "", fmt.Errorf("escape function cannot escape a %T value: %w", v, functions.ErrInvalidValueType)
	}
}

// escapeJS neutralizes the characters that would let a value break out of a
// JavaScript string literal or the surrounding <script> element. It escapes
// quotes, backslashes, and forward slashes. It escapes the characters that
// could open or close a script or comment tag (< > &), backticks, every control
// character (bytes 0 to 31), and the delete character (127). It also escapes
// U+2028 and U+2029, two invisible line separators that end a string literal in
// older JavaScript engines. The result is meant to sit inside an existing
// quoted string, so it does not add the surrounding quotes itself.
func escapeJS(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch r {
		case '\\':
			b.WriteString("\\\\")
		case '"':
			b.WriteString("\\\"")
		case '\'':
			b.WriteString("\\'")
		case '/':
			b.WriteString("\\/")
		case '`':
			b.WriteString("\\u0060")
		case '<':
			b.WriteString("\\u003C")
		case '>':
			b.WriteString("\\u003E")
		case '&':
			b.WriteString("\\u0026")
		case '\n':
			b.WriteString("\\n")
		case '\r':
			b.WriteString("\\r")
		case '\t':
			b.WriteString("\\t")
		case '\b':
			b.WriteString("\\b")
		case '\f':
			b.WriteString("\\f")
		case '\u2028':
			b.WriteString("\\u2028")
		case '\u2029':
			b.WriteString("\\u2029")
		default:
			if r < 0x20 || r == 0x7f {
				fmt.Fprintf(&b, "\\u%04X", r)
			} else {
				b.WriteRune(r)
			}
		}
	}
	return b.String()
}
