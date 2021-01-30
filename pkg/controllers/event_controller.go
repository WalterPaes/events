package controllers

import (
	"encoding/json"
	"events/pkg/domains/event"
	"events/pkg/services/account"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EventController struct {
	svc account.Service
}

func NewEventController(svc account.Service) *EventController {
	return &EventController{svc: svc}
}

func (ec EventController) Event(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}

	var ev *event.Event
	err = json.Unmarshal(body, &ev)

	ev.Handler(ec.svc)

}