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