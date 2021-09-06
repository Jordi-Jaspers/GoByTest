package mexican_wave

// https://www.codewars.com/kata/58f5c63f1e26ecda7e000029

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

func wave(words string) []string {
	var baseString = strings.ToLower(words)
	var chars = []rune(baseString)
	var solution = make([]string, GetAlphaNumericLength(baseString))
	var j = 0

	fmt.Printf("The length of the string is: %d \n", len(baseString))
	fmt.Printf("The amount of charachters are: %d \n", GetAlphaNumericLength(baseString))

	for i := 0; i <= len(baseString)-1; i++ {
		fmt.Printf("the 'i' counter is: %d \n", i)
		if strings.Contains(alpha, strings.ToLower(string(chars[i]))) {
			fmt.Printf("the 'j' counter is: %d \n", j)
			chars[i] = unicode.ToUpper(chars[i])
			solution[j] = string(chars)
			chars[i] = unicode.ToLower(chars[i])
			j++
		}
	}
	return solution
}

func GetAlphaNumericLength(words string) int {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := reg.ReplaceAllString(words, "")
	return len(processedString)
}
