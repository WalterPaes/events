package controllers

import (
	"events/pkg/services/account"
	"fmt"
	"net/http"
	"strconv"
)

type AccountController struct {
	svc account.Service
}

func NewAccountController(svc account.Service) *AccountController {
	return &AccountController{svc: svc}
}

func (ac AccountController) Reset(w http.ResponseWriter, _ *http.Request) {
	ac.svc.Reset()
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func (ac AccountController) Balance(w http.ResponseWriter, r *http.Request) {
	contentType := "text/plain"
	status := http.StatusOK
	var message interface{}

	accountId, err := strconv.Atoi(r.URL.Query().Get("account_id"))
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
	}

	acc, err := ac.svc.GetAccount(accountId)
	if err != nil {
		fmt.Println(err)
		status = http.StatusNotFound
		message = "0"
	}

	if err == nil {
		message = acc.GetBalance()
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	fmt.Fprint(w, message)
}
