package account

import (
	"errors"
	account2 "events/pkg/domains/account"
	"events/pkg/repositories/account"
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

func (svc Service) Deposit(id int, amount float64) error {
	if amount < 0 {
		return InvalidValueOfAmountError
	}

	acc, err := svc.repository.GetById(id)
	if err != nil {
		return err
	}

	acc.Deposit(amount)
	return nil
}

func (svc Service) Withdraw(id int, amount float64) error {
	acc, err := svc.repository.GetById(id)
	if err != nil {
		return err
	}

	if acc.GetBalance() < amount {
		return InsufficientFundsError
	}

	acc.Withdraw(amount)
	return nil
}

func (svc Service) Transfer(amount float64, originId, destinationId int) error {
	origin, err := svc.GetAccount(originId)
	if err != nil {
		return err
	}

	err = svc.Withdraw(origin.GetId(), amount)
	if err != nil {
		return err
	}

	destination, err := svc.GetAccount(destinationId)
	if err != nil {
		return err
	}

	err = svc.Deposit(destination.GetId(), amount)
	if err != nil {
		return err
	}

	return nil
}
