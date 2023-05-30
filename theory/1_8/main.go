package main

import "fmt"

// java나 python에서 a = b를 하면 얕은 복사가 되지만 go는 바로 깊은 복사가 되버린다.
// 따라서 메모리 주소값만 복사하고 싶다면 a = &b를 하면 된다.
// 또한 println으로 확인할때 메모리 주소를 보고 싶다면 fmt.Println(&a) 를 하면 된다.
// 그리고 메모리 주소가 아닌 value를 복사 하고 싶다면 c = *a 하면 된다.

func main() {
	a := 2
	b := &a
	c := *b

	a = 10
	fmt.Println(a, *b, c)

	fmt.Println(&a, b, &c)
}
