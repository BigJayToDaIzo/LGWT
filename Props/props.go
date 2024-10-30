package main

import (
	"strings"
)

func ConvertToRomanBigEndian(arabic int) string {
	var res strings.Builder
	for arabic > 0 {
		switch {
		case arabic/1000 >= 1:
			for i := 0; i < arabic/1000; i++ {
				res.WriteString("M")
				arabic -= 1000
			}
		case arabic/900 >= 1:
			res.WriteString("CM")
			arabic -= 900
		case arabic/500 >= 1:
			for i := 0; i < arabic/500; i++ {
				res.WriteString("D")
				arabic -= 500
			}
		case arabic/400 >= 1:
			res.WriteString("CD")
			arabic -= 400
		case arabic/100 >= 1:
			res.WriteString("C")
			arabic -= 100
		case arabic/50 >= 1:
			res.WriteString("L")
			arabic -= 50
		case arabic/10 >= 1:
			res.WriteString("X")
			arabic -= 10
		case arabic/5 >= 1:
			res.WriteString("V")
			arabic -= 5
		default:
			for i := 0; i < arabic; i++ {
				res.WriteString("I")
				arabic--
			}
		}
	}
	// don't want to return till we run out of arabic
	return res.String()
}

// Romans decreed NO symbol may repeat >3x
func ConvertToRoman(arabic int) string {
	// this switch length is a smell
	// parsing the int with mod may be better
	var res strings.Builder
	// base case for recursion?
	for arabic > 0 {
		// 5000 and above is ↁ or V^overline
		// this test will only cover up to 4999
		switch {
		case arabic >= 1000:
			res.WriteString("M")
			arabic -= 1000
		case arabic >= 900: // CM
			res.WriteString("CM")
			arabic -= 900
		case arabic >= 500: // D
			res.WriteString("D")
			arabic -= 500
		case arabic >= 400:
			res.WriteString("CD")
			arabic -= 400
		case arabic >= 100: // C
			res.WriteString("C")
			arabic -= 100
		case arabic >= 90:
			res.WriteString("XC")
			arabic -= 90
		case arabic >= 50: // L
			res.WriteString("L")
			arabic -= 50
		case arabic >= 40:
			res.WriteString("XL")
			arabic -= 40
		case arabic >= 10:
			res.WriteString("X")
			arabic -= 10
		case arabic >= 9:
			res.WriteString("IX")
			arabic -= 9
		case arabic >= 5:
			res.WriteString("V")
			arabic -= 5
		case arabic > 3:
			res.WriteString("IV")
			arabic -= 4
		default:
			res.WriteString(switchOnSmallArabic(arabic))
			arabic -= arabic
		}
	}
	return res.String()
}

// Rule for Roman is NO symbol may repeat >3x
func switchOnSmallArabic(arabic int) string {
	s := ""
	for arabic > 0 {
		s += "I"
		arabic--
	}
	return s
}
