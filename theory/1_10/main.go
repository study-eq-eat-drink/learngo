package main

import "fmt"

// go에는 구조를 미리 정의하는 Struct가 있다.
// Scala type class나 python pydantic의 model과 같이 데이터 구조를 미리 정의할 때 유용해보인다.

type person struct {
	name string
	age  int
	mans []string
}

func main() {
	nameLooksAge := map[string]string{"주완": "29", "다혜": "28"}
	fmt.Println(nameLooksAge)

	dMans := []string{"주완"}
	dahye := person{"다혜", 28, dMans}
	fmt.Println(dahye)

	dMans = []string{"다혜"}
	juwan := person{name: "주완", age: 29, mans: dMans}
	fmt.Println(juwan)

}
