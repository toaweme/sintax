package functions

func Default(varValue any, params []any) (any, error) {
	if varValue == nil {
		return params[0], nil
	}
	if varValue == "" {
		return params[0], nil
	}

	return varValue, nil
}
