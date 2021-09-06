package two_sum

// https://www.codewars.com/kata/52c31f8e6605bcc646000082/train/go

func TwoSum(numbers []int, target int) [2]int {
	for i, _ := range numbers {
		var result = ArrayRecursion(numbers, target, i)
		if result[0] != -1 {
			return result
		}
	}
	return [2]int{}
}

func ArrayRecursion(numbers []int, target int, lastIndex int) [2]int {
	for i, _ := range numbers {
		if target == (numbers[lastIndex]+numbers[i]) && lastIndex != i {
			return [2]int{lastIndex, i}
		}
	}
	return [2]int{-1, -1}
}
