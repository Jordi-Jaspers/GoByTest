package integers

import "fmt"

// Variadic function can be used in Go.
// Go does not support method overloading or argument defaulting.
func add(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println(add(2, 2))
}
