package account

import (
	"errors"
	"events/pkg/domains/account"
)

// Success mock
type RepositorySuccess struct{}

func (a RepositorySuccess) Create(acc account.Account) (account.Account, error) {
	ac := account.New(acc.GetId())
	ac.Deposit(250.0)
	return ac, nil
}

func (a RepositorySuccess) GetById(id int) (account.Account, error) {
	ac := account.New(id)
	ac.Deposit(250.0)
	return ac, nil
}

func (a RepositorySuccess) Update(acc account.Account) (account.Account, error) {
	return acc, nil
}

// Fail Mock
type RepositoryFail struct{}

func (a RepositoryFail) Create(acc account.Account) (account.Account, error) {
	return acc, errors.New("error")
}

func (a RepositoryFail) GetById(_ int) (account.Account, error) {
	return account.Account{}, errors.New("error")
}

func (a RepositoryFail) Update(acc account.Account) (account.Account, error) {
	return acc, errors.New("error")
}
