package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Basic struct {
	Message string `json:"message"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Basic{"Hello World. This is anonymous structure"}); err != nil {
		fmt.Println("Couldn't respond back with JSON value")
	}
}

func ShowAuthLogin(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Basic{"[GET] -> /auth/login"}); err != nil {
		fmt.Println("Couldn't respond back with JSON value")
	}
}

func ShowAuthRegister(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Basic{"[GET] -> /auth/register"}); err != nil {
		fmt.Println("Couldn't respond back with JSON value")
	}
}

func HandleAuthLogin(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Basic{"[POST] -> /auth/login"}); err != nil {
		fmt.Println("Couldn't respond back with JSON value")
	}
}

func HandleAuthRegister(w http.ResponseWriter, r *http.Request) {
	type registrationForm struct {
		Email, Username, Password string
	}

	if err := json.NewEncoder(w).Encode(registrationForm{
		r.FormValue("email"),
		r.FormValue("username"),
		r.FormValue("password"),
	}); err != nil {
		fmt.Println("Couldn't respond back with JSON value")
	}
}
