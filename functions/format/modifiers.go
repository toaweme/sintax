package format

import "github.com/toaweme/sintax/functions"

// Each modifier is a named, composed GlobalModifier so it can be referenced
// directly (in tests, or by a consumer wanting one modifier) without building
// the whole map. Modifiers assembles them for the engine. format, length,
// line_numbers, and decimal are Overloads over their value shapes and optional
// params.
var (
	formatModifier = functions.Overload(
		formatStringPassthrough,
		functions.WrapOne(FormatTime),
		functions.Wrap(FormatTimeDefault),
	)
	lengthModifier = functions.Overload(
		functions.Wrap(LengthString),
		functions.Wrap(LengthBytes),
		functions.Wrap(LengthReflect),
	)
	lineNumbersModifier = functions.Overload(
		lineNumbersNilEmpty,
		functions.Wrap(LineNumbers),
	)
	decimalModifier = functions.Overload(
		functions.WrapOne(DecimalPlaces),
		functions.Wrap(DecimalDefault),
	)
	currencyModifier = functions.WrapTwo(Currency)
)

// Modifiers returns the value-formatting modifiers keyed by their template
// names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFormat):      formatModifier,
		string(ModifierNameLength):      lengthModifier,
		string(ModifierNameLineNumbers): lineNumbersModifier,
		string(ModifierNameDecimal):     decimalModifier,
		string(ModifierNameCurrency):    currencyModifier,
	}
}
