package main

import "fmt"

func sumNumbers(numbers ...int) int {

	totalNumber := 0

	for _, number := range numbers {
		totalNumber += number
	}

	return totalNumber
}

func sumNumbersIndex(numbers ...int) int {

	totalNumber := 0
	for i := 0; i < len(numbers); i++ {
		totalNumber += i
	}
	return totalNumber
}

func main() {
	fmt.Println(sumNumbers(1, 2, 3, 4, 5))
	fmt.Println(sumNumbersIndex(1, 2, 3, 4, 5))
}
