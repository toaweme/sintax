// Package trim provides whitespace and affix trimming modifiers.
package trim

import (
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameTrim is the template name for the Trim modifier.
const ModifierNameTrim functions.ModifierName = "trim"

// ModifierNameTrimPrefix is the template name for the TrimPrefix modifier.
const ModifierNameTrimPrefix functions.ModifierName = "trim_prefix"

// ModifierNameTrimSuffix is the template name for the TrimSuffix modifier.
const ModifierNameTrimSuffix functions.ModifierName = "trim_suffix"

// affixCutset is the whitespace stripped by the no-argument prefix/suffix trims.
const affixCutset = "\n \t"

// The trim modifiers each decompose into four clauses: a value may be a string
// or []byte, and the cutset param may be present or absent. The no-param clause
// is a distinct operation (whitespace stripping) rather than a defaulted cutset,
// so it cannot be folded into the cutset clause: trim:'' must strip nothing,
// which a "" default would turn into whitespace stripping. Overload lists the
// cutset (higher-arity) clause first so the no-param clause is the fallback.

// TrimString strips leading and trailing whitespace from a string.
func TrimString(s string) (string, error) {
	return strings.TrimSpace(s), nil
}

// TrimStringSet strips any of the cutset characters from both ends of a string.
func TrimStringSet(s, cutset string) (string, error) {
	return strings.Trim(s, cutset), nil
}

// TrimBytes strips leading and trailing whitespace from a []byte.
func TrimBytes(b []byte) ([]byte, error) {
	return []byte(strings.TrimSpace(string(b))), nil
}

// TrimBytesSet strips any of the cutset characters from both ends of a []byte.
func TrimBytesSet(b []byte, cutset string) ([]byte, error) {
	return []byte(strings.Trim(string(b), cutset)), nil
}

// TrimPrefixString strips leading whitespace from a string.
func TrimPrefixString(s string) (string, error) {
	return strings.TrimLeft(s, affixCutset), nil
}

// TrimPrefixStringArg removes the given prefix once from the start of a string,
// returning it unchanged when it does not start with that prefix.
func TrimPrefixStringArg(s, prefix string) (string, error) {
	return strings.TrimPrefix(s, prefix), nil
}

// TrimPrefixBytes strips leading whitespace from a []byte.
func TrimPrefixBytes(b []byte) ([]byte, error) {
	return []byte(strings.TrimLeft(string(b), affixCutset)), nil
}

// TrimPrefixBytesArg removes the given prefix once from the start of a []byte.
func TrimPrefixBytesArg(b []byte, prefix string) ([]byte, error) {
	return []byte(strings.TrimPrefix(string(b), prefix)), nil
}

// TrimSuffixString strips trailing whitespace from a string.
func TrimSuffixString(s string) (string, error) {
	return strings.TrimRight(s, affixCutset), nil
}

// TrimSuffixStringArg removes the given suffix once from the end of a string,
// returning it unchanged when it does not end with that suffix.
func TrimSuffixStringArg(s, suffix string) (string, error) {
	return strings.TrimSuffix(s, suffix), nil
}

// TrimSuffixBytes strips trailing whitespace from a []byte.
func TrimSuffixBytes(b []byte) ([]byte, error) {
	return []byte(strings.TrimRight(string(b), affixCutset)), nil
}

// TrimSuffixBytesArg removes the given suffix once from the end of a []byte.
func TrimSuffixBytesArg(b []byte, suffix string) ([]byte, error) {
	return []byte(strings.TrimSuffix(string(b), suffix)), nil
}
