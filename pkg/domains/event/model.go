package event

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
