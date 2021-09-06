package gcd

// https://www.codewars.com/kata/54d7660d2daf68c619000d95

import "fmt"

func ConvertFracts(a [][]int) string {
	var denominator int = a[0][1]

	for i, _ := range a {
		fmt.Printf("position i(%d) gives GCD(%d, %d) \n", i, a[i][1], denominator)
		denominator = GCD(a[i][1], denominator)

		if denominator == 1 {
			break
		}
	}

	fmt.Printf("Common denominator is %d \n", denominator)

	for j, _ := range a {
		a[j][0] = a[j][0] * denominator
		a[j][1] = denominator
	}

	return fmt.Sprintf("%v", a)
}

func GCD(a int, b int) int {
	if a == 0 {
		fmt.Printf("returning b as %d \n", b)
		return b
	}
	fmt.Printf("%d modulo %d equals %d \n", b, a, b%a)
	return GCD(b%a, a)
}
