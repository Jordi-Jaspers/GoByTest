package digpow

// https://www.codewars.com/kata/5552101f47fc5178b1000050

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	var length = len(strconv.Itoa(n))
	var sum = 0.0
	var base = n

	// Modulus returns the rest after dividing a number.
	// This is a convinient way to split numbers to digits if you use a decimal baseline (power of 10).
	for i := 0; i < length; i++ {
		sum += math.Pow(float64(base%10), float64(length+p-1-i))
		base = int(base / 10)
	}

	// Using trunc to get rid of the irrational part of the result.
	// When the truncation of a float equals original result means,
	// there was no irrational part and the result is defined as an integer.
	var result = sum / float64(n)
	if result == math.Trunc(result) {
		return int(result)
	}
	return -1
}
