package transform

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameSum is the template name for the Sum modifier.
const ModifierNameSum functions.ModifierName = "sum"

// Sum adds up the elements of a slice and always returns a float. With no
// parameter every element must be a number or a string that parses as one, and
// a nil element counts as zero. With a string field parameter each element is
// treated as a map and the named field is summed instead, which is how you total
// one column across a list of records. A missing field or a value that is not a
// number is an error.
//
// value: array
// param:0?: string (a field name to sum across map elements instead of the elements themselves)
// returns: float
//
// example: total a list of amounts
// in:  amounts = [12.50, 8.00, 4.25]
// tpl: {{ amounts | sum }}
// out: 24.75
//
// example: total an order from line items
// in:  items = [{"name": "Mug", "price": 12}, {"name": "Pen", "price": 3}, {"name": "Pad", "price": 5}]
// tpl: {{ items | sum:'price' }}
// out: 20
//
// example: sum numeric strings coming from a form
// in:  quantities = ["1.5", "2.5"]
// tpl: {{ quantities | sum }}
// out: 4
func Sum(value any, params []any) (any, error) {
	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return float64(0), nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, fmt.Errorf("sum: expected slice or array, got %T", value)
	}

	var field string
	if len(params) > 0 {
		s, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("sum: field parameter must be a string, got %T", params[0])
		}
		field = s
	}

	var total float64
	for i := range rv.Len() {
		elem := rv.Index(i).Interface()
		var n float64
		var err error
		if field != "" {
			n, err = numberFromField(elem, field)
		} else {
			n, err = toFloat(elem)
		}
		if err != nil {
			return nil, fmt.Errorf("sum: element %d: %w", i, err)
		}
		total += n
	}
	return total, nil
}

func numberFromField(elem any, field string) (float64, error) {
	rv := reflect.ValueOf(elem)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return 0, nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Map {
		return 0, fmt.Errorf("expected map for field %q, got %T", field, elem)
	}
	for _, k := range rv.MapKeys() {
		if fmt.Sprint(k.Interface()) == field {
			return toFloat(rv.MapIndex(k).Interface())
		}
	}
	return 0, fmt.Errorf("field %q not found", field)
}

func toFloat(v any) (float64, error) {
	if v == nil {
		return 0, nil
	}
	switch x := v.(type) {
	case float64:
		return x, nil
	case float32:
		return float64(x), nil
	case int:
		return float64(x), nil
	case int8:
		return float64(x), nil
	case int16:
		return float64(x), nil
	case int32:
		return float64(x), nil
	case int64:
		return float64(x), nil
	case uint:
		return float64(x), nil
	case uint8:
		return float64(x), nil
	case uint16:
		return float64(x), nil
	case uint32:
		return float64(x), nil
	case uint64:
		return float64(x), nil
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot parse %q as number", x)
		}
		return f, nil
	}
	return 0, fmt.Errorf("cannot convert %T to number", v)
}
