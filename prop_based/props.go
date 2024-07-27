package proptests

import "strings"

type RomanNumeral struct {
	ArabicValue uint16
	Symbol      string
}

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

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.ArabicValue {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.ArabicValue
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	arabic := 0
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += int(numeral.ArabicValue)
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return uint16(arabic)
}
