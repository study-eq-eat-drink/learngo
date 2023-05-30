package main

import "fmt"

// go는 switch case문을 지원한다.

func isCanDrink(age int) bool {

	switch koreanAge := age + 1; koreanAge {
	case 17:
		return true
	case 18:
		return true
	}

	return false
}

func isCanDrink2(age int) bool {
	switch {
	case age > 18:
		return true
	}
	return false
}

func main() {

	fmt.Println(18, isCanDrink(18))
	fmt.Println(19, isCanDrink(19))
}
