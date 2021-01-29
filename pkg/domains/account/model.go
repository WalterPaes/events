package account

type Account struct {
	id int
	balance float64
}

func Create(id int) *Account {
	return &Account{id: id}
}

func GetAccount(id int) Account {
	return Account{}
}

func (a *Account) GetId() int {
	return a.id
}

func (a *Account) GetBalance() float64 {
	return a.balance
}