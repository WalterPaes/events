package account

import (
	"errors"
	account2 "events/pkg/domains/account"
	"events/pkg/repositories/account"
	"fmt"
)

var (
	InvalidValueOfAmountError = errors.New("invalid value of amount")
	InsufficientFundsError    = errors.New("insufficient funds")
)

type Service struct {
	repository account.RepositoryContract
}

func NewService(repository account.RepositoryContract) Service {
	return Service{repository: repository}
}

func (svc Service) Create(id int) (account2.Account, error) {
	acc, err := svc.repository.Create(account2.New(id))
	if err != nil {
		return acc, err
	}
	return acc, err
}

func (svc Service) GetAccount(id int) (account2.Account, error) {
	acc, err := svc.repository.GetById(id)
	if err != nil {
		return acc, err
	}
	return acc, err
}

func (svc Service) Deposit(id int, amount float64) (account2.Account, error) {
	fmt.Println("Deposit", id, amount)
	var acc account2.Account

	if amount < 0 {
		return acc, InvalidValueOfAmountError
	}

	acc, err := svc.repository.GetById(id)
	if err != nil {
		acc, err = svc.Create(id)
		if err != nil {
			return acc, err
		}
	}

	acc.Deposit(amount)
	acc, err = svc.repository.Update(acc)
	if err != nil {
		return acc, err
	}

	return acc, nil
}

func (svc Service) Withdraw(id int, amount float64) (account2.Account, error) {
	var acc account2.Account

	acc, err := svc.repository.GetById(id)
	if err != nil {
		return acc, err
	}

	if acc.GetBalance() < amount {
		return acc, InsufficientFundsError
	}

	acc.Withdraw(amount)
	acc, err = svc.repository.Update(acc)
	if err != nil {
		return acc, err
	}
	return acc, nil
}

func (svc Service) Transfer(amount float64, originId, destinationId int) (account2.Account, account2.Account, error) {
	var origin account2.Account
	var destination account2.Account

	origin, err := svc.Withdraw(originId, amount)
	if err != nil {
		return origin, destination, err
	}

	destination, err = svc.Deposit(destinationId, amount)
	if err != nil {
		return origin, destination, err
	}

	origin, err = svc.repository.Update(origin)
	if err != nil {
		return origin, destination, err
	}

	destination, err = svc.repository.Update(destination)
	if err != nil {
		return origin, destination, err
	}
	return origin, destination, nil
}

func (svc Service) Reset() {
	svc.repository.Reset()
}
