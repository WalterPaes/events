package event

import (
	"errors"
	"events/pkg/domains/account"
)

type EventType string

const (
	deposit  EventType = "deposit"
	withdraw EventType = "withdraw"
	transfer EventType = "transfer"
)

type Event struct {
	Type        EventType `json:"type"`
	Amount      float64   `json:"amount"`
	Destination int       `json:"destination"`
	Origin      int       `json:"origin"`
}

func (e *Event) Handler() error {
	var err error

	switch e.Type {
	case deposit:
		acc := account.GetAccount(e.Destination)
		err = acc.Deposit(e.Amount)
	case withdraw:
		acc := account.GetAccount(e.Origin)
		err = acc.Withdraw(e.Amount)
	case transfer:
		accOr := account.GetAccount(e.Origin)
		accDest := account.GetAccount(e.Destination)
		err = accOr.Transfer(e.Amount, accDest.GetId())
	default:
		err = errors.New("invalid event")
	}

	return err
}
