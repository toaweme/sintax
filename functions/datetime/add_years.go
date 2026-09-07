package datetime

import (
	"time"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameAddYears is the template name for the AddYears modifier.
const ModifierNameAddYears functions.ModifierName = "add_years"

// AddYears shifts a date by a whole number of years, forward on a positive
// amount and back on a negative one. A warranty and a contract term are stated
// in years, and the count is usually a field rather than a literal, so writing
// it as twelve times that field would need arithmetic inside an argument, which
// an argument cannot hold.
//
// The 29th of February in a year that has one lands on the 28th when the target
// year does not, on the same rule AddMonths clamps by.
func AddYears(t time.Time, years int) (time.Time, error) {
	return shiftMonths(t, years*12), nil
}
