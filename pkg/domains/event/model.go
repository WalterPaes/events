package event

import (
	"errors"
	"events/pkg/services/account"
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

func (e *Event) Handler(svc account.Service) error {
	var err error

	switch e.Type {
	case deposit:
		err = svc.Deposit(e.Destination, e.Amount)
	case withdraw:
		err = svc.Withdraw(e.Origin, e.Amount)
	case transfer:
		err = svc.Transfer(e.Amount, e.Origin, e.Destination)
	default:
		err = errors.New("invalid event")
	}

	return err
}
