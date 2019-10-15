package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting Wisbday. Let's buckle-up to handle traffic bois")

	NewDatabase()
	bootstrapWebserver()
}

func bootstrapWebserver() {
	r := mux.NewRouter()
	initRoutes(r)

	if err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r)); err != nil {
		panic("Cannot bind to *:8080 and listen for connections")
	}
}

func initRoutes(r *mux.Router) {
	r.HandleFunc("/", IndexHandler)

	r.HandleFunc("/auth/login", ShowAuthLogin).Methods("GET")
	r.HandleFunc("/auth/register", ShowAuthRegister).Methods("GET")
	r.HandleFunc("/auth/login", HandleAuthLogin).Methods("POST")
	r.HandleFunc("/auth/register", HandleAuthRegister).Methods("POST")
}
