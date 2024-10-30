package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		// I implemented for ADDITION
		{"1 converts to I", 1, "I"},
		// {"2 converts to II", 2, "II"},

		// V implemented for SUBTRACTION from at 4
		{"4 converts to IV", 4, "IV"},
		// TODO: introduction of new symbol for subtraction is causing breaks
		// in an terative approach look for pattern in the messy switch

		// V implemented for ADDITION
		{"5 converts to V", 5, "V"},

		// X implemented for SUBTRACTION from at 9
		{"9 converts to IX", 9, "IX"},
		// X implemented for ADDITION at 10
		{"10 converts to X", 10, "X"},

		// L implemented for SUBTRACTION from at 40
		{"40 converts to XL", 40, "XL"},
		// check previous implementation
		{"44 converts to XLIV", 44, "XLIV"},
		{"49 converts to XLIX", 49, "XLIX"},

		// L implemented for ADDITION at 50
		{"50 converts to L", 50, "L"},
		// check previous implementation
		{"54 converts to LIV", 54, "LIV"},
		{"59 converts to LIX", 59, "LIX"},
		// right before edge of introduction of C for subtraction
		{"89 converts to LXXXIX", 89, "LXXXIX"},

		// C implemented for SUBTRACTION at 90
		{"90 converts to XC", 90, "XC"},
		// check previous breaking points
		{"94 converts to XCIV", 94, "XCIV"},
		{"99 converts to XCIX", 99, "XCIX"},

		// C implemented for ADDITION at 100
		{"100 converts to C", 100, "C"},
		// check for subtraction regressions due to new digit
		{"104 converts to CIV", 104, "CIV"},
		{"109 converts to CIX", 109, "CIX"},
		// push the envelope
		{"149 converts to CXLIX", 149, "CXLIX"},
		{"249 converts to CCXLIX", 249, "CCXLIX"},
		{"333 converts to CCCXXXIII", 333, "CCCXXXIII"},
		// nothing should break until we need D for subtraction right?
		{"399 converts to CCCXCIX", 399, "CCCXCIX"},
		// we knew needing the D would trip us up here
		// implements D symbol for SUBTRACTION
		{"400 converts to CD", 400, "CD"},
		{"404 converts to CDIV", 404, "CDIV"},
		{"409 converts to CDIX", 409, "CDIX"},
		{"433 converts to CDXXXIII", 433, "CDXXXIII"},
		{"499 converts to CDXCIX", 499, "CDXCIX"},

		// NOW we add the D for ADDITION
		{"500 converts to D", 500, "D"},
		{"504 converts to DIV", 504, "DIV"},
		{"599 converts to DXCIX", 599, "DXCIX"},
		{"649 converts to DCXLIX", 649, "DCXLIX"},
		{"899 converts to DCCCXCIX", 899, "DCCCXCIX"},

		// implements M symbol for SUBTRACTION
		{"900 converts to CM", 900, "CM"},
		{"994 converts to CMXCIV", 994, "CMXCIV"},
		{"1000 converts to M", 1000, "M"},
		{"3999 converts to MMMCMXCIX", 3999, "MMMCMXCIX"},
		// 4k and beyond requires new symbol for 'large 4k'
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}

func TestRomanNumeralsBigEndian(t *testing.T) {
	cases := []struct {
		Arabic    int
		Assertion string
		Want      string
	}{
		// addition through the symbols
		{1000, "converts to M", "M"},
		{2000, "converts to MM", "MM"},
		{3000, "converts to MMM", "MMM"},
		{3500, "converts to MMMD", "MMMD"},
		{3600, "converts to MMMDC", "MMMDC"},
		{3700, "converts to MMMDCC", "MMMDCC"},
		{3800, "converts to MMMDCCC", "MMMDCCC"},
		{3850, "converts to MMMDCCCL", "MMMDCCCL"},
		{3900, "converts to MMMCM", "MMMCM"},
		{3910, "converts to MMMCMX", "MMMCMX"},
		{3920, "converts to MMMCMXX", "MMMCMXX"},
		{3930, "converts to MMMCMXXX", "MMMCMXXX"},
		{3935, "converts to MMMCMXXXV", "MMMCMXXXV"},
		{3936, "converts to MMMCMXXXVI", "MMMCMXXXVI"},
		{3938, "converts to MMMCMXXXVIII", "MMMCMXXXVIII"},
		// time to bring in subtraction
		{3400, "converts to MMMCD", "MMMCD"},
		{3440, "converts to MMMCDXL", "MMMCDXL"},
		{3444, "converts to MMMCDXLIV", "MMMCDXLIV"},
		// did we do it?  Let's PLAY!
		{3442, "converts to MMMCDXLII", "MMMCDXLII"},
		{3441, "converts to MMMCDXLI", "MMMCDXLI"},
		// let is begin again near the lost trail of 3444
		// now that subtraction is implemented lets lean back into the bigendian thing
		// and continue toward 3999 on the 4s and 9s
		{3990, "converts to MMMCMXC", "MMMCMXC"},
		{3994, "converts to MMMCMXCIV", "MMMCMXCIV"},
		{3999, "converts to MMMCMXCIX", "MMMCMXCIX"},

		// check backwards from 1000
		{999, "converts to CMXCIX", "CMXCIX"},
		{895, "converts to DCCCXCV", "DCCCXCV"},
		{749, "converts to DCCXLIX", "DCCXLIX"},
		{694, "converts to DCXCIV", "DCXCIV"},
		{544, "converts to DXLIV", "DXLIV"},
		// {3940, "converts to MMMCMXL", "MMMCMXL"},
		// {3944, "converts to MMMCMXLIV", "MMMCMXLIV"},
		// {3945, "converts to MMMCMXLV", "MMMCMXLV"},
		// {3949, "converts to MMMCMXLIX", "MMMCMXLIX"},
		// dare we try the TIPPY TOP now that subrtaction has been (hopefully) wroked in?
		// yes, yes we should (documention, no, no we should not have, we ended up back at 3444)
		// {3999, "converts to MMMCMXCIX", "MMMCMXCIX"},
		// predictably borkied 😂
	}
	for _, test := range cases {
		t.Run(test.Assertion, func(t *testing.T) {
			got := ConvertToRomanBigEndian(test.Arabic)
			if got != test.Want {
				t.Errorf("got arabic: %d %q, want %q", test.Arabic, got, test.Want)
			}
		})
	}
}
