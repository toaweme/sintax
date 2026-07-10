package functions

import "fmt"

// AsText decorates a modifier so a scalar value - a number or a bool - is handed
// to it as its default string form, for text-first modifiers where 42 should
// read as "42". It is opt-in per modifier: wrap the registration, for example
// AsText(Wrap(ToUpper)). Strings are passed through unchanged, and nil, slices,
// maps, and every other value are passed through untouched so the modifier still
// sees and rejects them. Because it only rewrites the value (never the params)
// and runs before the wrapped modifier, it composes with Wrap, Overload, and any
// other GlobalModifier, and it leaves value-type dispatch strict.
func AsText(mod GlobalModifier) GlobalModifier {
	return func(value any, params []any) (any, error) {
		if s, ok := Stringish(value); ok {
			value = s
		}
		return mod(value, params)
	}
}

// Stringish returns the default string form of a scalar primitive (a string,
// bool, or any int/uint/float kind) and true, and reports false for every other
// value - including nil, []byte, slices, and maps - so a caller leaves those
// untouched. It backs AsText and is exported for modifiers such as concat that
// stringify their params too.
func Stringish(v any) (string, bool) {
	switch x := v.(type) {
	case string:
		return x, true
	case bool,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return fmt.Sprint(x), true
	default:
		return "", false
	}
}
