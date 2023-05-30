package main

// naked function은 return할 변수를 미리 정하고 따로 return에 변수를 지정하지 않는다.

// defer는 function이 return된 이후로 실행된다.
// function 실행후 buffer나 connection을 close하거나 특정 호출을 할 때 쓰자
// 이후 동시성일까...?

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, upperName string) {
	length = len(name)
	upperName = strings.ToUpper(name)

	return
}

func printManyWords(words ...string) {
	defer fmt.Println("i'm done")
	fmt.Println(words)
}

func main() {

	var length, upperName = lenAndUpper("juwan")

	println(length, upperName)

	printManyWords("asd", "qwe", "func smail gate")
}
