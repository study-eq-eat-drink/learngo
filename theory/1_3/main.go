package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func printManyWords(words ...string) {
	fmt.Println(words)
}

func main() {

	var length, upperName = lenAndUpper("juwan")

	println(length, upperName)

	printManyWords("asd", "qwe", "func smail gate")
}
