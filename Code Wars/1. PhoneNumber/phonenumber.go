package phonenumber

// https://www.codewars.com/kata/525f50e3b73515a6db000b83/train/go

import (
	"fmt"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var length = len(numbers)
	var telephone strings.Builder

	if length == 10 {
		for i := 0; i < 10; i++ {
			if numbers[i] <= 9 {
				if i == 0 {
					telephone.WriteString("(")
				}
				if i == 3 {
					telephone.WriteString(") ")
				}
				if i == 6 {
					telephone.WriteString("-")
				}
				telephone.WriteString(fmt.Sprintf("%v", numbers[i]))
			}
		}
		return telephone.String()
	}
	return ""
}
