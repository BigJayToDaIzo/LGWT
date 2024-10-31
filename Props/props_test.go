package main

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Arabic uint16
		Roman  string
	}{
		// big endian approach works slightly better than
		// little endian approach, imo
		// addition through the symbols
		{Arabic: 1000, Roman: "M"},
		{Arabic: 2000, Roman: "MM"},
		{Arabic: 3000, Roman: "MMM"},
		{Arabic: 3500, Roman: "MMMD"},
		{Arabic: 3600, Roman: "MMMDC"},
		{Arabic: 3700, Roman: "MMMDCC"},
		{Arabic: 3800, Roman: "MMMDCCC"},
		{Arabic: 3850, Roman: "MMMDCCCL"},
		// time to bring in subtraction
		{Arabic: 3900, Roman: "MMMCM"},
		{Arabic: 3400, Roman: "MMMCD"},
		// bring in X
		{Arabic: 3910, Roman: "MMMCMX"},
		{Arabic: 3920, Roman: "MMMCMXX"},
		{Arabic: 3930, Roman: "MMMCMXXX"},
		// X subtraction
		{Arabic: 3440, Roman: "MMMCDXL"},
		// bring in V
		{Arabic: 3935, Roman: "MMMCMXXXV"},
		// bring in I
		{Arabic: 3936, Roman: "MMMCMXXXVI"},
		{Arabic: 3938, Roman: "MMMCMXXXVIII"},
		{Arabic: 3444, Roman: "MMMCDXLIV"},
		// did we do it?  Let's PLAY!
		{Arabic: 3442, Roman: "MMMCDXLII"},
		{Arabic: 3441, Roman: "MMMCDXLI"},
		// let is begin again near the lost trail of 3444
		// now that subtraction is implemented lets lean back into the bigendian thing
		// and continue toward 3999 on the 4s and 9s
		{Arabic: 3990, Roman: "MMMCMXC"},
		{Arabic: 3994, Roman: "MMMCMXCIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},

		// check backwards from 1000
		{Arabic: 999, Roman: "CMXCIX"},
		{Arabic: 895, Roman: "DCCCXCV"},
		{Arabic: 749, Roman: "DCCXLIX"},
		{Arabic: 694, Roman: "DCXCIV"},
		{Arabic: 544, Roman: "DXLIV"},
		{Arabic: 499, Roman: "CDXCIX"},
		{Arabic: 440, Roman: "CDXL"},
		{Arabic: 199, Roman: "CXCIX"},
		{Arabic: 99, Roman: "XCIX"},
		{Arabic: 94, Roman: "XCIV"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 1, Roman: "I"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got arabic: %d %q, want %q", test.Arabic, got, test.Roman)
			}
		})
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got roman: %q %d, want %d", test.Roman, got, test.Arabic)
			}
		})
	}
}

// testing/quick will point out weaknesses in our code with
// randomized data to fit the function signatures
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic < 0 || arabic > 3999 {
			log.Println("arabic out of range", arabic)
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}
	if err := quick.Check(
		assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
	// passing nil as default config
	// if err := quick.Check(assertion, nil); err != nil {
	// 	t.Error("failed checks", err)
	// }
}
