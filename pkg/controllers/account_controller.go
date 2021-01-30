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

func (ac AccountController) Balance(w http.ResponseWriter, r *http.Request) {
	accountId, err := strconv.Atoi(r.URL.Query().Get("account_id"))
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}

	acc, err := ac.svc.GetAccount(accountId)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, 0)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, acc.GetBalance())
}
