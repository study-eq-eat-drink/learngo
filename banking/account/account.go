package account

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

func CreateAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// go에서 method는 아래와 같이 작성한다.
// go에서는 reciver라고 표현한다.
// go method에서 a Account는 복사 되어 넘어오기 때문에
// 만약 java와 같이 this 즉 자기 자신을 변형 시키려면 *를 추가해야한다.
func (a *Account) Deposit(number int) {

	a.balance += number
}

func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("거지는 돈을 뺄 수 없습니다")

// go에는 try execption이 없다. 헐
// 대신 error라는게 존재하며 에러로 인해 프로그램을 종료하려면 log.Fatalln 과 같은 함수를 사용한다.
// 한마디로 프로그래머가 모든 함수에 대한 에러를 return 받은 후 체크해야한다.
func (a *Account) Withdraw(amount int) error {

	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount

	return nil
}

func (a Account) String() string {
	return fmt.Sprint(a.owner, "는 다음과 같은 금액을 소지하고 있습니다.\nMoney : ", a.balance)
}
