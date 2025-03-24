package functions

func isParam(params []any, index int, name string) bool {
	if len(params) <= index {
		return false
	}

	return params[index] == name
}

func ConditionIsTrue(condition any) bool {
	// log.Error("ConditionIsTrue", "value", condition, "type", fmt.Sprintf("%T", condition))
	if condition == nil {
		return false
	}
	switch v := condition.(type) {
	case bool:
		return v == true
	case string:
		// condition can be rendered before being passed
		// in cases where variable is boolean, we render true/false
		if v == "false" {
			return false
		}
		return len(v) > 0
	case int:
		return v > 0
	case int8:
		return v > 0
	case int16:
		return v > 0
	case int32:
		return v > 0
	case int64:
		return v > 0
	case uint:
		return v > 0
	case uint8:
		return v > 0
	case uint16:
		return v > 0
	case uint32:
		return v > 0
	case uint64:
		return v > 0
	case float32:
		return v > 0
	case float64:
		return v > 0
	default:
		return false
	}
}
