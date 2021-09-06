package two_oldest_ages

import "sort"

// https://www.codewars.com/kata/511f11d355fe575d2c000001

func TwoOldestAges1(ages []int) [2]int {
	var max1 = 0
	var max2 = 0

	for i, _ := range ages {
		if max1 < ages[i] {
			max2 = max1
			max1 = ages[i]
		} else if max2 < ages[i] {
			max2 = ages[i]
		}
	}

	return [2]int{max2, max1}
}

func TwoOldestAges2(ages []int) [2]int {
	sort.Sort(sort.Reverse(sort.IntSlice(ages)))
	return [2]int{ages[1], ages[0]}
}
