package account

type Account struct {
	owner   string
	balance int
}

func CreateAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
