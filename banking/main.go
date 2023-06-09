package main

import (
	"fmt"

	"github.com/study-eq-eat-drink/learngo/banking/account"
)

func main() {
	account := account.CreateAccount("juwan")
	fmt.Println(account)

	account.Deposit(100)
	fmt.Println(account.Balance())

	// drawError := account.Withdraw(200)
	// if drawError != nil {
	// 	log.Fatalln(drawError)
	// }

	fmt.Println(account)

}
