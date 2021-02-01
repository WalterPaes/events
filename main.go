package main

import (
	"events/pkg/controllers"
	account2 "events/pkg/repositories/account"
	"events/pkg/services/account"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	repository := account2.NewAccountRepository()
	accountController := controllers.NewAccountController(
		account.NewService(repository))

	eventController := controllers.NewEventController(
		account.NewService(repository))

	r := mux.NewRouter()
	r.HandleFunc("/reset", accountController.Reset).Methods("POST")
	r.HandleFunc("/balance", accountController.Balance).Methods("GET")
	r.HandleFunc("/event", eventController.Event).Methods("POST")
	http.Handle("/", r)

	log.Println("Server is Ok!")
	log.Fatal(http.ListenAndServe(":3000", r))
}
