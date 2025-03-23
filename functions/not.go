package functions

import (
	"github.com/toaweme/log"
)

var Not = func(s any, _ []any) (any, error) {
	log.Debug("not", "!ConditionIsTrue(s)", !ConditionIsTrue(s), "ConditionIsTrue(s)", ConditionIsTrue(s))
	return !ConditionIsTrue(s), nil
}
