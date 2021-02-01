package event

import (
	"encoding/json"
	"errors"
	"events/pkg/services/account"
	"strconv"
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
	Destination string    `json:"destination,omitempty"`
	Origin      string    `json:"origin,omitempty"`
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
		destination, _ := strconv.Atoi(e.Destination)
		acc, err := svc.Deposit(destination, e.Amount)
		if err != nil {
			return "", err
		}

		var result struct {
			Destination struct {
				ID      string     `json:"id"`
				Balance float64 `json:"balance"`
			} `json:"destination"`
		}

		result.Destination.ID = strconv.Itoa(acc.GetId())
		result.Destination.Balance = acc.GetBalance()

		response, err := json.Marshal(result)
		if err != nil {
			return "", err
		}

		return string(response), nil
	case withdraw:
		origin, _ := strconv.Atoi(e.Origin)
		acc, err := svc.Withdraw(origin, e.Amount)
		if err != nil {
			return "", err
		}

		var result struct {
			Origin struct {
				ID      string     `json:"id"`
				Balance float64 `json:"balance"`
			} `json:"origin"`
		}

		result.Origin.ID = strconv.Itoa(acc.GetId())
		result.Origin.Balance = acc.GetBalance()

		response, err := json.Marshal(result)
		if err != nil {
			return "", err
		}

		return string(response), nil
	case transfer:
		destinationId, _ := strconv.Atoi(e.Destination)
		originId, _ := strconv.Atoi(e.Origin)

		origin, destination, err := svc.Transfer(e.Amount, originId, destinationId)
		if err != nil {
			return "", err
		}

		var result struct {
			Origin struct {
				ID      string     `json:"id"`
				Balance float64 `json:"balance"`
			} `json:"origin"`
			Destination struct {
				ID      string     `json:"id"`
				Balance float64 `json:"balance"`
			} `json:"destination"`
		}

		result.Origin.ID = strconv.Itoa(origin.GetId())
		result.Origin.Balance = origin.GetBalance()
		result.Destination.ID = strconv.Itoa(destination.GetId())
		result.Destination.Balance = destination.GetBalance()

		response, err := json.Marshal(result)
		if err != nil {
			return "", err
		}

		return string(response), nil
	default:
		return "", errors.New("invalid event")
	}
}
