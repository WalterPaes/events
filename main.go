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
	accountController := controllers.NewAccountController(
		account.NewService(&account2.AccountRepository{}))

	eventController := controllers.NewEventController(
		account.NewService(&account2.AccountRepository{}))

	r := mux.NewRouter()
	//r.HandleFunc("/reset", nil).Methods("POST")
	r.HandleFunc("/balance", accountController.Balance).Methods("GET")
	r.HandleFunc("/event", eventController.Event).Methods("POST")
	http.Handle("/", r)

	http.ListenAndServe(":8081", r)
	log.Println("Server is Ok!")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
