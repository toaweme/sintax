package format

import "github.com/toaweme/sintax/functions"

// Modifiers returns the value-formatting modifiers keyed by their template names.
func Modifiers() map[string]functions.GlobalModifier {
	return map[string]functions.GlobalModifier{
		string(ModifierNameFormat):      Format,
		string(ModifierNameLength):      Length,
		string(ModifierNameLineNumbers): LineNumbers,
		string(ModifierNameDecimal):     Decimal,
		string(ModifierNameCurrency):    Currency,
	}
}
