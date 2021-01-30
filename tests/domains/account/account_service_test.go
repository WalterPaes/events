package account

import (
	account3 "events/pkg/services/account"
	account2 "events/tests/repositories/account"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		service := account3.NewService(&account2.RepositorySuccess{})
		id := 100
		acc, err := service.Create(id)
		if err != nil {
			t.Fatal("Errors was not expected")
		}

		if id != acc.GetId() {
			t.Errorf("Was expected %d, but got %d", id, acc.GetId())
		}
	})

	t.Run("Error case", func(t *testing.T) {
		service := account3.NewService(&account2.RepositoryFail{})
		id := 100
		_, err := service.Create(id)
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}

func TestAccountGetId(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		service := account3.NewService(&account2.RepositorySuccess{})
		_, err := service.GetAccount(100)
		if err != nil {
			t.Error("Errors wasn't expected")
		}
	})
}

func TestAccountDeposit(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		service := account3.NewService(&account2.RepositorySuccess{})
		_, err := service.Deposit(100, 250.0)
		if err != nil {
			t.Error("Errors wasn't expected")
		}
	})

	t.Run("Error Case", func(t *testing.T) {
		service := account3.NewService(&account2.RepositoryFail{})
		_, err := service.Deposit(100, 250.0)
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}

func TestAccountWithdraw(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		service := account3.NewService(&account2.RepositorySuccess{})
		_, err := service.Withdraw(100, 250.0)
		if err != nil {
			t.Error("Errors wasn't expected")
		}
	})

	t.Run("Error Case", func(t *testing.T) {
		service := account3.NewService(&account2.RepositoryFail{})
		_, err := service.Withdraw(100, 250.0)
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}

func TestAccountTransfer(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		service := account3.NewService(&account2.RepositorySuccess{})
		_, _, err := service.Transfer(250.0, 100, 101)
		if err != nil {
			t.Error("Errors wasn't expected")
		}
	})
}
