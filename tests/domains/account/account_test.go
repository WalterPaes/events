package account

import (
	"events/pkg/domains/account"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	id := 100
	acc := account.Create(id)

	if id != acc.GetId() {
		t.Errorf("Was expected %d, but got %d", id, acc.GetId())
	}
}

func TestAccountGetId(t *testing.T) {
	acc := account.Create(100)

	amount := 250.0
	err := acc.Deposit(amount)
	if err != nil {
		t.Fatal("Errors wasn't expected")
	}

	if acc.GetBalance() != amount {
		t.Errorf("Was expected %.2f, but got %.2f", amount, acc.GetBalance())
	}
}

func TestAccountDeposit(t *testing.T) {
	acc := account.Create(100)

	t.Run("Success Case", func(t *testing.T) {
		amount := 250.0
		err := acc.Deposit(amount)
		if err != nil {
			t.Error("Errors wasn't expected")
		}
	})

	t.Run("Error Case", func(t *testing.T) {
		amount := -250.0
		err := acc.Deposit(amount)
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}

func TestAccountWithdraw(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		acc := account.Create(100)
		err := acc.Deposit(250.0)
		if err != nil {
			t.Fatal("Errors wasn't expected")
		}

		amount := 47.0
		err = acc.Withdraw(amount)
		if err != nil {
			t.Error("Errors wasn't expected")
		}
	})

	t.Run("Error Case", func(t *testing.T) {
		acc := account.Create(100)
		amount := 47.0
		err := acc.Withdraw(amount)
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}

func TestAccountTransfer(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		firstAcc := account.Create(100)
		secondAcc := account.Create(101)

		err := firstAcc.Deposit(786.99)
		if err != nil {
			t.Fatal("Errors wasn't expected")
		}

		amount := 102.0

		err = firstAcc.Transfer(amount, secondAcc.GetId())
		if err != nil {
			t.Error("Errors wasn't expected")
		}
	})

	t.Run("Error Case: InsufficientFundsError", func(t *testing.T) {
		firstAcc := account.Create(100)
		secondAcc := account.Create(101)

		amount := 102.0

		err := firstAcc.Transfer(amount, secondAcc.GetId())
		if err == nil {
			t.Error("Errors wasn't expected")
		}
	})

	t.Run("Error Case: InvalidValueOfAmountError", func(t *testing.T) {
		firstAcc := account.Create(100)
		secondAcc := account.Create(101)

		err := firstAcc.Deposit(786.99)
		if err != nil {
			t.Fatal("Errors wasn't expected")
		}

		amount := -102.0

		err = firstAcc.Transfer(amount, secondAcc.GetId())
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}
