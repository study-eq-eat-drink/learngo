package main

import (
	"fmt"

	"github.com/study-eq-eat-drink/learngo/banking/account"
)

func main() {
	account := account.CreateAccount("juwan")
	fmt.Println(account)
}
