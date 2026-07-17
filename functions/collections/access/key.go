package access

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameKey is the template name for the Key modifier.
const ModifierNameKey functions.ModifierName = "key"

// Key reads one value out of a map by key or out of a slice by index. Pass a
// string to look up a map key, and use dot notation in that string to walk into
// nested maps (for example 'database.host'). Pass a number to index into a
// slice.
//
// A lookup that finds nothing is a miss rather than a failure. A missing key, a
// path that runs out partway, an out-of-range index, and a nil value all report
// one. That is catchable, so `| default:'x'` supplies a fallback, an if reads it
// as false, and a for iterates nothing. Uncaught, it fails the render rather
// than quietly rendering empty, since a misspelled key that renders nothing is
// indistinguishable from data that was genuinely absent.
//
// Being handed something that cannot be looked up at all is a different thing
// and stays terminal. No key parameter, a non-string key, or a value that is
// neither map nor slice means the template is wrong, and no default rescues it.
func Key(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("key requires a key parameter: %w", functions.ErrMissingParam)
	}

	if value == nil {
		return nil, functions.Miss("key found no value to look in")
	}

	rv := reflect.ValueOf(value)

	switch rv.Type().Kind() {
	case reflect.Map:
		return handleMap(rv, params)
	case reflect.Slice, reflect.Array:
		return handleSlice(rv, params)
	case reflect.Pointer:
		if rv.IsNil() {
			return nil, functions.Miss("key found no value to look in")
		}
		return Key(rv.Elem().Interface(), params)
	default:
		return nil, fmt.Errorf("key expected a map or slice, got %T: %w", value, functions.ErrInvalidValueType)
	}
}

func handleMap(rv reflect.Value, params []any) (any, error) {
	parts, err := keyParts(params)
	if err != nil {
		return nil, err
	}

	current := rv
	for i, part := range parts {
		// Dereference pointers and unwrap interfaces
		for current.Kind() == reflect.Pointer || current.Kind() == reflect.Interface {
			if current.IsNil() {
				return nil, functions.Miss("key path stops at %q, which holds nothing", part)
			}
			current = current.Elem()
		}

		// the path runs out rather than the template being wrong. this data simply
		// does not nest that deep, which is the same shape of absence as a key that
		// is not there.
		if current.Kind() != reflect.Map {
			return nil, functions.Miss("key path segment %q is not a map, cannot look deeper", part)
		}

		keyValue, err := convertToMapKeyType(part, current.Type().Key())
		if err != nil {
			return nil, functions.Miss("key %q cannot exist in a map keyed by %v", part, current.Type().Key())
		}

		nextVal := current.MapIndex(keyValue)
		if !nextVal.IsValid() {
			return nil, functions.Miss("key %q not found", part)
		}

		if i == len(parts)-1 {
			return nextVal.Interface(), nil
		}

		current = nextVal
	}

	return nil, functions.Miss("key found no value for path %v", parts)
}

func handleSlice(rv reflect.Value, params []any) (any, error) {
	index, err := convertToInt(params[0])
	if err != nil {
		return nil, fmt.Errorf("key expected an index for a slice, got %T: %w", params[0], functions.ErrInvalidParamType)
	}

	length := rv.Len()
	if index < 0 || index >= length {
		return nil, functions.Miss("key index %d is outside the slice's %d element(s)", index, length)
	}

	return rv.Index(index).Interface(), nil
}

func convertToMapKeyType(key string, keyType reflect.Type) (reflect.Value, error) {
	switch keyType.Kind() {
	case reflect.String:
		return reflect.ValueOf(key), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("cannot convert %q to int: %w", key, err)
		}
		return reflect.ValueOf(intVal).Convert(keyType), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal, err := strconv.ParseUint(key, 10, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("cannot convert %q to uint: %w", key, err)
		}
		return reflect.ValueOf(uintVal).Convert(keyType), nil
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(key, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("cannot convert %q to float: %w", key, err)
		}
		return reflect.ValueOf(floatVal).Convert(keyType), nil
	default:
		return reflect.Value{}, fmt.Errorf("unsupported map key type: %v", keyType.Kind())
	}
}

// convertToInt is the lenient index coercion for slice access: it reuses the
// strict functions.ValueInt for the numeric kinds, then adds the leniency an
// array index wants but Wrap's int slot must not have - truncating a fractional
// float and parsing a numeric string.
func convertToInt(value any) (int, error) {
	if n, ok := functions.ValueInt(value); ok {
		return n, nil
	}
	switch v := value.(type) {
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("cannot convert %T to int", value)
	}
}

func keyParts(params []any) ([]string, error) {
	keyPath, ok := params[0].(string)
	if !ok {
		return nil, fmt.Errorf("key expected a string key, got %T: %w", params[0], functions.ErrInvalidParamType)
	}

	parts := strings.Split(keyPath, ".")
	return parts, nil
}
