package functions

import "errors"

type Type string

const (
	TypeAny    Type = "any"
	TypeString Type = "string"
	TypeBytes  Type = "bytes"
	TypeInt    Type = "int"
	TypeFloat  Type = "float"
	TypeBool   Type = "bool"
	TypeArray  Type = "array"
	TypeMap    Type = "map"
)

type Param struct {
	Type     Type
	Required bool
	Index    int
	// Value is the expected value of the parameter
	// empty string means any value is accepted
	Value any
}

type ModifierDefinition struct {
	Description    string
	Func           func(value any, params []any) (any, error)
	AcceptedValue  []Type
	AcceptedParams []Param

	value  any
	params []any
}

type ModifierError struct {
	Template     string
	Function     string
	Value        any
	ValueType    Type
	Params       []Param
	Err          error
	ErrLocalized string
}

func (e *ModifierError) Error() string {
	return e.ErrLocalized
}

var ErrInvalidValueType = errors.New("invalid value type")
var ErrInvalidParamType = errors.New("invalid param type")
var ErrMissingParam = errors.New("missing param")
var ErrInvalidParamValue = errors.New("invalid param value")

func (md *ModifierDefinition) Validate(value any, params []any) error {
	if !isTypeAllowed(md.AcceptedValue, value) {
		return &ModifierError{
			Value:        value,
			ValueType:    detectType(value),
			Params:       nil,
			Err:          ErrInvalidValueType,
			ErrLocalized: "",
		}
	}
	//
	// for _, p := range md.AcceptedParams {
	// 	if len(params) < p.Index {
	// 		return &ModifierError{
	// 			Value:        value,
	// 			ValueType:    detectType(value),
	// 			Params:       nil,
	// 			Err:          ErrInvalidValueType,
	// 			ErrLocalized: "",
	// 		}
	// 	}
	// }
	return nil
}

func isTypeAllowed(types []Type, value any) bool {
	valType := detectType(value)
	for _, t := range types {
		if t == TypeAny || t == valType {
			return true
		}
	}
	return false
}

func detectType(v any) Type {
	switch v.(type) {
	case string:
		return TypeString
	case []byte:
		return TypeBytes
	case int, int64, int32:
		return TypeInt
	case float32, float64:
		return TypeFloat
	case bool:
		return TypeBool
	case []any:
		return TypeArray
	case map[string]any:
		return TypeMap
	default:
		return TypeAny
	}
}

func isParamType(t Type, p any) bool {
	return detectType(p) == t
}
