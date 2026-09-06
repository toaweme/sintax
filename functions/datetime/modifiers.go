package datetime

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine.
var (
	addDaysModifier     = functions.WrapOne(AddDays)
	addMonthsModifier   = functions.WrapOne(AddMonths)
	addYearsModifier    = functions.WrapOne(AddYears)
	daysBetweenModifier = functions.WrapOne(DaysBetween)
)

// Modifiers returns the date arithmetic modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameAddDays):     addDaysModifier,
		string(ModifierNameAddMonths):   addMonthsModifier,
		string(ModifierNameAddYears):    addYearsModifier,
		string(ModifierNameDaysBetween): daysBetweenModifier,
	}
}
