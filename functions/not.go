package functions

var Not = func(value any, _ []any) (any, error) {
	return ConditionIsTrue(value) == false, nil
}
