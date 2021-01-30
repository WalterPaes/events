package account

type Account struct {
	id      int
	balance float64
}

func New(id int) Account {
	return Account{id: id}
}

func (a *Account) GetId() int {
	return a.id
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	a.balance -= amount
}
