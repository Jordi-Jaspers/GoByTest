package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

type testCase struct {
	Number uint16
	Roman  string
}

var cases = []testCase{
	{Number: 1, Roman: "I"},
	{Number: 2, Roman: "II"},
	{Number: 3, Roman: "III"},
	{Number: 4, Roman: "IV"},
	{Number: 5, Roman: "V"},
	{Number: 6, Roman: "VI"},
	{Number: 7, Roman: "VII"},
	{Number: 8, Roman: "VIII"},
	{Number: 9, Roman: "IX"},
	{Number: 10, Roman: "X"},
	{Number: 14, Roman: "XIV"},
	{Number: 18, Roman: "XVIII"},
	{Number: 20, Roman: "XX"},
	{Number: 39, Roman: "XXXIX"},
	{Number: 40, Roman: "XL"},
	{Number: 47, Roman: "XLVII"},
	{Number: 49, Roman: "XLIX"},
	{Number: 50, Roman: "L"},
	{Number: 100, Roman: "C"},
	{Number: 90, Roman: "XC"},
	{Number: 400, Roman: "CD"},
	{Number: 500, Roman: "D"},
	{Number: 900, Roman: "CM"},
	{Number: 1000, Roman: "M"},
	{Number: 1984, Roman: "MCMLXXXIV"},
	{Number: 3999, Roman: "MMMCMXCIX"},
	{Number: 2014, Roman: "MMXIV"},
	{Number: 1006, Roman: "MVI"},
	{Number: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d can be converted to %s", c.Number, c.Roman), func(t *testing.T) {
			got := ConvertToRoman(c.Number)
			want := c.Roman

			AssertRomanError(t, got, want)
		})
		t.Run(fmt.Sprintf("%s can be converted to %d", c.Roman, c.Number), func(t *testing.T) {
			got := ConvertToNumber(c.Roman)
			want := c.Number

			AssertNumberError(t, got, want)
		})
	}
	t.Run("Performing  a quick check with quick.Check package", func(t *testing.T) {
		assertion := func(number uint16) bool {
			if number > 3999 {
				return true
			}
			t.Log("testing", number)
			roman := ConvertToRoman(number)
			fromRoman := ConvertToNumber(roman)
			return fromRoman == number
		}

		if err := quick.Check(assertion, &quick.Config{
			MaxCount: 1000,
		}); err != nil {
			t.Error("failed checks", err)
		}
	})
}

func AssertRomanError(t testing.TB, got string, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertNumberError(t testing.TB, got uint16, want uint16) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
