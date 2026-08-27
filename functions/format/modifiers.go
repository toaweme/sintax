package format

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine. date, from_date,
// length, line_numbers, and decimal are Overloads over their value shapes and
// optional params.
var (
	dateModifier = functions.Overload(
		functions.WrapOne(Date),
		functions.Wrap(DateDefault),
	)
	lengthModifier = functions.Overload(
		functions.Wrap(LengthString),
		functions.Wrap(LengthBytes),
		functions.Wrap(LengthReflect),
	)
	lineNumbersModifier = functions.Overload(
		lineNumbersNilEmpty,
		functions.WrapOne(LineNumbersFrom),
		functions.Wrap(LineNumbers),
	)
	decimalModifier = functions.Overload(
		functions.WrapOne(DecimalPlaces),
		functions.Wrap(DecimalDefault),
	)
	currencyModifier = functions.WrapTwo(Currency)
	fromDateModifier = functions.Overload(
		functions.WrapOne(FromDate),
		functions.Wrap(FromDateGuess),
	)
)

// Modifiers returns the value formatting modifiers, and the date reader that
// shares their layout language, keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameDate):        dateModifier,
		string(ModifierNameLength):      lengthModifier,
		string(ModifierNameLineNumbers): lineNumbersModifier,
		string(ModifierNameDecimal):     decimalModifier,
		string(ModifierNameCurrency):    currencyModifier,
		string(ModifierNameFromDate):    fromDateModifier,
	}
}
