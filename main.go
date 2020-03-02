package main

import (
	"github.com/codegangsta/negroni"
	"github.com/dankusuma/learngolang/Controller"

	"github.com/dankusuma/learngolang/Auth"

	"github.com/gorilla/mux"

	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/GenerateToken", Auth.GenerateToken)
	router.HandleFunc("/Register", Auth.ValidateMiddleware(Controller.CreateCustomer))
	router.HandleFunc("/Login", Auth.ValidateMiddleware(Controller.Login))
	n := negroni.New()
	n.UseHandler(router)

	// port := "8080"
	port := os.Getenv("PORT")
	n.Run(":" + port)
}
