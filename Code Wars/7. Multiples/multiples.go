package main

//https://www.codewars.com/kata/514b92a657cdc65150000006/train/go

func main() {

}

func Multiple3And5(number int) int {
	var sum = 0
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
