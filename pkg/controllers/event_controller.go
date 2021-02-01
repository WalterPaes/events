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

	contentType := "text/plain"
	status := http.StatusCreated
	var message interface{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
	}

	var ev *event.Event
	err = json.Unmarshal(body, &ev)

	result, err := ev.Handler(ec.svc)
	if err != nil {
		fmt.Println(err)
		status = http.StatusNotFound
		message = "0"
	}

	if err == nil {
		contentType = "application/json"
		status = http.StatusCreated
		message = result
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	fmt.Fprint(w, message)
}