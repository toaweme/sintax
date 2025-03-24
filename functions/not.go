package functions

var Not = func(s any, _ []any) (any, error) {
	return ConditionIsTrue(s) == false, nil
}
