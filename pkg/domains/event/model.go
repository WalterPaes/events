package event

import (
	"encoding/json"
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
	Destination int       `json:"destination,omitempty"`
	Origin      int       `json:"origin,omitempty"`
}

type EventResponse struct {
	Destination struct {
		ID      int     `json:"id"`
		Balance float64 `json:"balance"`
	} `json:"destination,omitempty"`
	Origin struct {
		ID      int     `json:"id"`
		Balance float64 `json:"balance"`
	} `json:"origin,omitempty"`
}

func (e *Event) Handler(svc account.Service) (string, error) {
	switch e.Type {
	case deposit:
		acc, err := svc.Deposit(e.Destination, e.Amount)
		if err != nil {
			return "", err
		}

		er := &EventResponse{
			Destination: struct {
				ID      int     `json:"id"`
				Balance float64 `json:"balance"`
			}{
				acc.GetId(),
				acc.GetBalance(),
			},
		}

		response, err := json.Marshal(er)
		if err != nil {
			return "", err
		}

		return string(response), nil
	case withdraw:
		acc, err := svc.Withdraw(e.Origin, e.Amount)
		if err != nil {
			return "", err
		}

		er := &EventResponse{
			Origin: struct {
				ID      int     `json:"id"`
				Balance float64 `json:"balance"`
			}{
				acc.GetId(),
				acc.GetBalance(),
			},
		}

		response, err := json.Marshal(er)
		if err != nil {
			return "", err
		}

		return string(response), nil
	case transfer:
		origin, destination, err := svc.Transfer(e.Amount, e.Origin, e.Destination)
		if err != nil {
			return "", err
		}

		er := &EventResponse{
			Destination: struct {
				ID      int     `json:"id"`
				Balance float64 `json:"balance"`
			}{
				destination.GetId(),
				destination.GetBalance(),
			},
			Origin: struct {
				ID      int     `json:"id"`
				Balance float64 `json:"balance"`
			}{
				origin.GetId(),
				origin.GetBalance(),
			},
		}

		response, err := json.Marshal(er)
		if err != nil {
			return "", err
		}

		return string(response), nil
	default:
		return "", errors.New("invalid event")
	}
}
