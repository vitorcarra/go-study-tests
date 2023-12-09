package numbers

import "strings"

// RomanNumeral represents a Roman numeral with its corresponding value and symbol.
type RomanNumeral struct {
	Value  int    // The numerical value of the Roman numeral.
	Symbol string // The symbol representation of the Roman numeral.
}

// allRomanNumerals is a slice of RomanNumeral structs that represents a collection of Roman numerals.
var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ConvertToRoman converts an integer to its Roman numeral representation.
func ConvertToRoman(n int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for n >= numeral.Value {
			result.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}

	return result.String()
}
