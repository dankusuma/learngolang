package main

import (
	"go-rest-api/Auth"
	"go-rest-api/Controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/GenerateJwt", Auth.GenerateToken)
	router.HandleFunc("/Register", Auth.ValidateMiddleware(Controller.CreateCustomer))
	router.HandleFunc("/Login", Auth.ValidateMiddleware(Controller.Login))
	log.Fatal(http.ListenAndServe(":8080", router))
}
