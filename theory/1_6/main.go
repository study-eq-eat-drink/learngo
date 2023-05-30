package main

import "fmt"

// go는 if문 조건절에 변수를 만들 수 있다.

func isCanDrink(age int) bool {
	if koreanAge := age + 1; koreanAge > 19 {
		return true
	}

	return false
}

func main() {

	fmt.Println(18, isCanDrink(18))
	fmt.Println(19, isCanDrink(19))
}
