package functions

import "reflect"

// The Wrap family adapts a typed modifier into the engine's untyped
// GlobalModifier signature. A modifier author writes a plain, testable function
// (for example func(string) (string, error)) and Wrap coerces the piped value
// and params to the declared types, so the boilerplate that each modifier would
// otherwise hand-roll (and get subtly wrong) lives in one audited place. The
// type arguments are inferred from the passed function, so call sites read
// Wrap(ToLower) and WrapTwo(ReplaceString) with no explicit brackets.
//
// There is one Wrap per parameter arity because Go has no variadic type
// parameters. WrapVariadic covers the "one repeated param type" case. A modifier
// whose param types change with the value's shape does not fit a single typed
// signature; register it as an Overload of several typed clauses, or leave it a
// plain GlobalModifier.
//
// On a coercion miss every Wrap returns the bare type sentinel
// (ErrInvalidValueType, ErrInvalidParamType, ErrMissingParam) with no formatting,
// which keeps rejection allocation-free. That matters because rejection is not
// only an error path: it is the normal signal Overload reads to fall through to
// the next clause, so a modifier such as eq rejects its numeric clause on every
// string argument. The concrete detail (which value, which type) is added once,
// lazily, by whoever surfaces the failure - Overload wraps its terminal
// no-clause-matched error with the value type and param count.

// Wrap adapts a no-parameter typed modifier, func(In) (Out, error). It matches
// its arity exactly: a clause that declares no params rejects a call that passes
// any, so under Overload a stray or mistyped argument falls through to another
// clause (or surfaces as an error) rather than being silently ignored.
func Wrap[In, Out any](fn func(In) (Out, error)) GlobalModifier {
	return func(value any, params []any) (any, error) {
		if len(params) > 0 {
			return nil, ErrInvalidParamType
		}
		in, ok := coerce[In](value)
		if !ok {
			return nil, ErrInvalidValueType
		}
		return fn(in)
	}
}

// WrapOne adapts a one-parameter typed modifier, func(In, P0) (Out, error). It
// matches exactly one param, rejecting a call that passes more.
func WrapOne[In, P0, Out any](fn func(In, P0) (Out, error)) GlobalModifier {
	return func(value any, params []any) (any, error) {
		if len(params) > 1 {
			return nil, ErrInvalidParamType
		}
		in, ok := coerce[In](value)
		if !ok {
			return nil, ErrInvalidValueType
		}
		p0, err := param[P0](params, 0)
		if err != nil {
			return nil, err
		}
		return fn(in, p0)
	}
}

// WrapTwo adapts a two-parameter typed modifier, func(In, P0, P1) (Out, error).
// It matches exactly two params, rejecting a call that passes more.
func WrapTwo[In, P0, P1, Out any](fn func(In, P0, P1) (Out, error)) GlobalModifier {
	return func(value any, params []any) (any, error) {
		if len(params) > 2 {
			return nil, ErrInvalidParamType
		}
		in, ok := coerce[In](value)
		if !ok {
			return nil, ErrInvalidValueType
		}
		p0, err := param[P0](params, 0)
		if err != nil {
			return nil, err
		}
		p1, err := param[P1](params, 1)
		if err != nil {
			return nil, err
		}
		return fn(in, p0, p1)
	}
}

// WrapVariadic adapts a variadic typed modifier, func(In, ...P) (Out, error),
// where every param shares the type P.
func WrapVariadic[In, P, Out any](fn func(In, ...P) (Out, error)) GlobalModifier {
	return func(value any, params []any) (any, error) {
		in, ok := coerce[In](value)
		if !ok {
			return nil, ErrInvalidValueType
		}
		ps := make([]P, len(params))
		for i, raw := range params {
			p, ok := coerce[P](raw)
			if !ok {
				return nil, ErrInvalidParamType
			}
			ps[i] = p
		}
		return fn(in, ps...)
	}
}

// coerce converts an untyped value to T. It first tries a direct assertion, then
// falls back to the engine's coercions for the types that have one (string,
// float64 across the numeric kinds, int across the numeric kinds, and []any
// across slice kinds). []byte and map[string]any succeed only on a direct match,
// which is what keeps a string clause and a []byte clause of the same modifier
// distinct under Overload. Any other T succeeds only on a direct match.
func coerce[T any](v any) (T, bool) {
	if tv, ok := v.(T); ok {
		return tv, true
	}
	var zero T
	if v == nil && reflect.TypeFor[T]().Kind() == reflect.Interface {
		// a nil interface fails the assertion above, but nil legitimately
		// satisfies an interface target (T == any, whose zero value is nil).
		// Concrete targets fall through to the switch, where a type-specific
		// coercion may still accept nil (ValueNumber(nil) == 0 for float64).
		return zero, true
	}
	// each case matched on T's own zero value, so the coerced result is always a
	// T; the comma-ok assertion is for the linter, the ok can never be false here.
	switch any(zero).(type) {
	case string:
		if s, err := ValueString(v); err == nil {
			out, _ := any(s).(T)
			return out, true
		}
	case float64:
		if n, err := ValueNumber(v); err == nil {
			out, _ := any(n).(T)
			return out, true
		}
	case int:
		if n, ok := ValueInt(v); ok {
			out, _ := any(n).(T)
			return out, true
		}
	case []any:
		if sl, err := ValueSlice(v); err == nil {
			out, _ := any(sl).(T)
			return out, true
		}
	}
	return zero, false
}

// param coerces the argument at index i to T, returning the bare ErrMissingParam
// when the argument is absent and ErrInvalidParamType when it is the wrong type.
func param[T any](params []any, i int) (T, error) {
	var zero T
	if i >= len(params) {
		return zero, ErrMissingParam
	}
	v, ok := coerce[T](params[i])
	if !ok {
		return zero, ErrInvalidParamType
	}
	return v, nil
}
