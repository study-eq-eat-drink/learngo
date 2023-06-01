package main

import "fmt"

// go에는 Array와 java, python에서 list와 비슷한 slice가 있다.
// go array는 java array와 동일하게 크기가 생성할때 지정된다.
// slice는 list와 같이 값을 append(slice, object)로 추가 할 수 있다.
// append를 하여도 기존 slice 객체에 추가되는것이 아니라 값이 추가된 새로운 slice가 반환된다.

func main() {
	arrayNames := [5]string{"주완", "다혜", "지온"}
	arrayNames[4] = "누구인가"
	fmt.Println(arrayNames)

	sliceNames := []string{"복순", "주완", "다혜"}
	fmt.Println(&sliceNames)
	sliceNames = append(sliceNames, "누구냐고")
	fmt.Println(&sliceNames)
	fmt.Println(sliceNames)

}
