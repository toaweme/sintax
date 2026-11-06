package escape

import (
	"strconv"
	"strings"
	"unicode/utf16"
)

// Backslash escapes a value by putting a backslash before every character of
// the first set, and before the value's first character when it is in the
// second set. Android's resource strings escape a leading @ or ? that way,
// because only there do they read as a reference.
//
// A set lists characters as they follow the backslash, so n, t and r stand
// for a line break, a tab and a carriage return and are written as \n, \t
// and \r. The letters themselves are never escaped. A backslash in a set
// escapes the backslash itself, which every set that escapes anything needs,
// since otherwise a value ending in one runs into the next character.
func Backslash(value any, sets ...string) (string, error) {
	str, err := stringify(value)
	if err != nil {
		return "", err
	}
	anywhere := backslashSet(sets, 0)
	leading := backslashSet(sets, 1)
	var b strings.Builder
	b.Grow(len(str))
	for index, r := range str {
		escaped, named := backslashName(r)
		if !named && strings.ContainsRune("ntruU", r) {
			b.WriteRune(r)
			continue
		}
		if anywhere[escaped] || (index == 0 && leading[escaped]) {
			b.WriteByte('\\')
			b.WriteRune(escaped)
			continue
		}
		b.WriteRune(r)
	}
	return b.String(), nil
}

// Unbackslash reverses Backslash for every character named in any of the sets,
// in one pass, so an escaped backslash followed by an n reads as a backslash
// and an n rather than as a line break. A u or a U in a set also reads \uXXXX
// and \UXXXX as the character those four hex digits name, joining a surrogate
// pair into one character. A backslash before anything outside the sets is
// kept as written.
func Unbackslash(value any, sets ...string) (string, error) {
	str, err := stringify(value)
	if err != nil {
		return "", err
	}
	known := backslashSet(sets, -1)
	runes := []rune(str)
	var b strings.Builder
	b.Grow(len(str))
	for i := 0; i < len(runes); i++ {
		if runes[i] != '\\' || i+1 >= len(runes) {
			b.WriteRune(runes[i])
			continue
		}
		next := runes[i+1]
		if (next == 'u' || next == 'U') && known[next] {
			if code, ok := hexRune(runes, i+2); ok {
				i += 5
				if utf16.IsSurrogate(code) && i+2 < len(runes) && runes[i+1] == '\\' && (runes[i+2] == 'u' || runes[i+2] == 'U') {
					if low, ok := hexRune(runes, i+3); ok {
						if joined := utf16.DecodeRune(code, low); joined != '�' {
							b.WriteRune(joined)
							i += 6
							continue
						}
					}
				}
				b.WriteRune(code)
				continue
			}
			b.WriteRune(runes[i])
			continue
		}
		if !known[next] {
			b.WriteRune(runes[i])
			continue
		}
		i++
		switch next {
		case 'n':
			b.WriteRune('\n')
		case 't':
			b.WriteRune('\t')
		case 'r':
			b.WriteRune('\r')
		default:
			b.WriteRune(next)
		}
	}
	return b.String(), nil
}

// backslashName answers the character written after the backslash for r,
// which is a letter for the three control characters that have one.
func backslashName(r rune) (rune, bool) {
	switch r {
	case '\n':
		return 'n', true
	case '\t':
		return 't', true
	case '\r':
		return 'r', true
	}
	return r, false
}

// backslashSet reads set index of sets, or every set joined when index is
// negative. A control character written raw counts as its letter.
func backslashSet(sets []string, index int) map[rune]bool {
	chosen := sets
	if index >= 0 {
		if index >= len(sets) {
			return nil
		}
		chosen = sets[index : index+1]
	}
	out := map[rune]bool{}
	for _, set := range chosen {
		for _, r := range set {
			name, _ := backslashName(r)
			out[name] = true
		}
	}
	return out
}

func hexRune(runes []rune, start int) (rune, bool) {
	if start+4 > len(runes) {
		return 0, false
	}
	code, err := strconv.ParseUint(string(runes[start:start+4]), 16, 32)
	if err != nil {
		return 0, false
	}
	return rune(code), true
}
