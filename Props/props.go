package main

import (
	"strings"
)

// abstracting of of arabic to roman symbol
// to a struct and ordered slice killed the
// smelly switch statement
type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanSymbols = []RomanNumeral{
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
	if arabic == 0 {
		return ""
	}
	var result strings.Builder
	for _, numeral := range allRomanSymbols {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	if roman == "" {
		return 0
	}
	var result uint16
	// clip symbols big endian and iterate arabic
	for _, numeral := range allRomanSymbols {
		for strings.HasPrefix(roman, numeral.Symbol) {
			result += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return result
}

// Romans decreed NO symbol may repeat >3x
// func ConvertToRoman(arabic int) string {
// 	// this switch length is a smell
// 	// parsing the int with mod may be better
// 	var res strings.Builder
// 	// base case for recursion?
// 	for arabic > 0 {
// 		// 4000 and above is denoted by IV^overline
// 		// 5000 and above is ↁ or V^overline, etc
// 		// this test will only cover up to 4999
// 		switch {
// 		case arabic >= 1000:
// 			res.WriteString("M")
// 			arabic -= 1000
// 		case arabic >= 900: // CM
// 			res.WriteString("CM")
// 			arabic -= 900
// 		case arabic >= 500: // D
// 			res.WriteString("D")
// 			arabic -= 500
// 		case arabic >= 400:
// 			res.WriteString("CD")
// 			arabic -= 400
// 		case arabic >= 100: // C
// 			res.WriteString("C")
// 			arabic -= 100
// 		case arabic >= 90:
// 			res.WriteString("XC")
// 			arabic -= 90
// 		case arabic >= 50: // L
// 			res.WriteString("L")
// 			arabic -= 50
// 		case arabic >= 40:
// 			res.WriteString("XL")
// 			arabic -= 40
// 		case arabic >= 10:
// 			res.WriteString("X")
// 			arabic -= 10
// 		case arabic >= 9:
// 			res.WriteString("IX")
// 			arabic -= 9
// 		case arabic >= 5:
// 			res.WriteString("V")
// 			arabic -= 5
// 		case arabic > 3:
// 			res.WriteString("IV")
// 			arabic -= 4
// 		default:
// 			res.WriteString(switchOnSmallArabic(arabic))
// 			arabic -= arabic
// 		}
// 	}
// 	return res.String()
// }

// Rule for Roman is NO symbol may repeat >3x
// func switchOnSmallArabic(arabic int) string {
// 	s := ""
// 	for arabic > 0 {
// 		s += "I"
// 		arabic--
// 	}
// 	return s
// }
