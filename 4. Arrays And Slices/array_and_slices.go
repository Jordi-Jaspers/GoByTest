package arrays_and_slices

// The slice type which allows us to have collections of any size.
// The syntax is very similar to arrays, you just omit the size when declaring them.
// Go has slices which do not encode the size of the collection and instead can have any size
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// len gets the length of the array
// make initializes the specified type with give length
func SumAll(arr ...[]int) []int {
	amountOfArrays := len(arr)
	sum := make([]int, amountOfArrays)

	for i, numbers := range arr {
		sum[i] = Sum(numbers)
	}
	return sum
}

// append: makes a slice and a new value, then returns a new slice with all the items in it.
// The following code will return a runtime error whenever an empty slice is given.
func SumAllTails(arr ...[]int) []int {
	var newArr []int

	for _, list := range arr {
		if len(list) == 0 {
			newArr = append(newArr, 0)
		} else {
			newArr = append(newArr, Sum(list[1:]))
		}
	}
	return newArr
}
