package account

import "errors"

var (
	InvalidValueOfAmountError = errors.New("invalid value of amount")
	InsufficientFundsError = errors.New("insufficient funds")
)

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

func (a *Account) Deposit(amount float64) error {
	if amount < 0 {
		return InvalidValueOfAmountError
	}
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return InsufficientFundsError
	}
	a.balance -= amount
	return nil
}

func (a *Account) Transfer(amount float64, destinationId int) error {
	err := a.Withdraw(amount)
	if err != nil {
		return err
	}

	destination := GetAccount(destinationId)
	err = destination.Deposit(amount)
	if err != nil {
		return err
	}

	return nil
}