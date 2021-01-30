package account

import (
	"errors"
	"events/pkg/domains/account"
)

type RepositoryContract interface {
	Create(acc account.Account) (account.Account, error)
	GetById(id int) (account.Account, error)
	Update(acc account.Account) (account.Account, error)
}

type AccountRepository struct {
	data map[int]account.Account
}

func (a AccountRepository) Create(acc account.Account) (account.Account, error) {
	_, exists := a.data[acc.GetId()]
	if exists {
		return acc, errors.New("account already exist")
	}

	a.data[acc.GetId()] = acc

	return acc, nil
}

func (a AccountRepository) GetById(id int) (account.Account, error) {
	var acc account.Account
	acc, exists := a.data[id]
	if !exists {
		return acc, errors.New("account not exist")
	}

	return acc, nil
}

func (a AccountRepository) Update(acc account.Account) (account.Account, error) {
	_, exists := a.data[acc.GetId()]
	if !exists {
		return acc, errors.New("account not exist")
	}

	a.data[acc.GetId()] = acc

	return acc, nil
}