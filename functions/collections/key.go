package collections

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Key(value any, params []any) (any, error) {
	val, err := key(value, params)
	if err != nil {
		// maybe this shouldn't be an error
		return nil, nil
	}
	return val, err
}

func key(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("key function requires a key parameter")
	}

	if value == nil {
		return nil, fmt.Errorf("key function: value is nil")
	}

	rv := reflect.ValueOf(value)
	rt := rv.Type()

	switch rt.Kind() {
	case reflect.Map:
		return handleMap(rv, params)
	case reflect.Slice, reflect.Array:
		return handleSlice(rv, params)
	case reflect.Ptr:
		if rv.IsNil() {
			return nil, fmt.Errorf("key function: pointer is nil")
		}
		return key(rv.Elem().Interface(), params)
	default:
		return nil, fmt.Errorf("key function expected map or slice, got %T", value)
	}
}

func handleMap(rv reflect.Value, params []any) (any, error) {
	parts, err := keyParts(params)
	if err != nil {
		return nil, err
	}

	current := rv
	for i, part := range parts {
		if current.Type().Kind() != reflect.Map {
			return nil, fmt.Errorf("key function: path segment %q is not a map; cannot continue nested lookup", part)
		}

		keyValue, err := convertToMapKeyType(part, current.Type().Key())
		if err != nil {
			return nil, fmt.Errorf("key function: cannot convert key %q to map key type %v: %w", part, current.Type().Key(), err)
		}

		nextVal := current.MapIndex(keyValue)
		if !nextVal.IsValid() {
			return nil, fmt.Errorf("key function: path segment %q not found in map", part)
		}

		if i == len(parts)-1 {
			return nextVal.Interface(), nil
		}

		current = nextVal
	}

	return nil, fmt.Errorf("key function: no value found for key path %v", parts)
}

func handleSlice(rv reflect.Value, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("key function: slice access requires an index parameter")
	}

	index, err := convertToInt(params[0])
	if err != nil {
		return nil, fmt.Errorf("key function: index for slice must be convertible to int: %w", err)
	}

	length := rv.Len()
	if index < 0 || index >= length {
		return nil, fmt.Errorf("key function: index %d out of range [0:%d)", index, length)
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

func convertToInt(value any) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
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
		return nil, fmt.Errorf("key function requires a string key parameter")
	}

	parts := strings.Split(keyPath, ".")
	return parts, nil
}
