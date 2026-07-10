package functions

import (
	"errors"
	"fmt"
)

// Overload binds one modifier name to several typed clauses, for the modifiers
// whose behaviour changes with the shape of the piped value or the number of
// params (for example trim over a string versus []byte, or sum over the whole
// slice versus one field). Each clause is an ordinary Wrap* adapter; Overload
// tries them in the given order and returns the first that accepts the inputs.
//
// A clause "rejects" the inputs by returning one of the coercion sentinels
// (ErrInvalidValueType, ErrInvalidParamType, ErrMissingParam), which is exactly
// what a Wrap* adapter emits when the value or a param will not coerce to the
// clause's declared type, or when a required param is absent. That is the only
// signal Overload needs, so clauses carry no separate predicate. Any other error
// is a genuine failure from inside a matched clause and is returned as-is.
//
// Order is significant: list the most specific clause first and the most
// permissive (an any-typed value, or the lowest arity) last, so a broad clause
// cannot shadow a narrower one. With no clause accepting the inputs, Overload
// returns the last rejection so the caller sees a concrete type mismatch.
func Overload(clauses ...GlobalModifier) GlobalModifier {
	return func(value any, params []any) (any, error) {
		// track value-shape and param rejections separately. A param rejection
		// means a clause accepted the value but its params did not fit, which is
		// more informative than a bare value mismatch, so it wins when both occur
		// (for example find on a slice with a non-string key reports the param
		// error, not that the value is not a map).
		var valueReject, paramReject error
		for _, clause := range clauses {
			out, err := clause(value, params)
			if err != nil {
				switch {
				case errors.Is(err, ErrInvalidParamType), errors.Is(err, ErrMissingParam):
					if paramReject == nil {
						paramReject = err
					}
					continue
				case errors.Is(err, ErrInvalidValueType):
					if valueReject == nil {
						valueReject = err
					}
					continue
				}
			}
			return out, err
		}
		// every clause declined; add the concrete detail once here, wrapping the
		// most informative sentinel so errors.Is still identifies the mismatch.
		reject := paramReject
		if reject == nil {
			reject = valueReject
		}
		if reject == nil {
			reject = ErrInvalidValueType
		}
		return nil, fmt.Errorf("no overload accepts a %T value with %d param(s): %w", value, len(params), reject)
	}
}
