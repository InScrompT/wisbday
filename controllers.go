package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	WriteAsJSON(w, "Hello World. This is anonymous structure")
}

func ShowAuthLogin(w http.ResponseWriter, r *http.Request) {
	WriteAsJSON(w, "[GET] -> /auth/login")
}

func ShowAuthRegister(w http.ResponseWriter, r *http.Request) {
	WriteAsJSON(w, "[GET] -> /auth/register")
}

func HandleAuthLogin(w http.ResponseWriter, r *http.Request) {
	type loginForm struct {
		Email, Password string
	}

	theUser := &User{}
	theForm := &loginForm{
		Email: r.FormValue("email"), // Could be either a username or an email ID
		Password: r.FormValue("password"),
	}

	if DB.Where(&User{Email: theForm.Email}).Or(&User{Username: theForm.Email}).First(&theUser).RecordNotFound() {
		WriteAsJSON(w, "The user associated with this email or username is not found")
		return
	}

	if !CheckPasswordHash(theForm.Password, theUser.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		WriteAsJSON(w, "The password you entered is wrong")

		return
	}

	WriteAsJSON(w, "You are identified. Logged in (Dummy)")
}

func HandleAuthRegister(w http.ResponseWriter, r *http.Request) {
	type registrationForm struct {
		Email, Username, Password string
	}

	if err = json.NewEncoder(w).Encode(registrationForm{
		r.FormValue("email"),
		r.FormValue("username"),
		r.FormValue("password"),
	}); err != nil {
		fmt.Println("Couldn't respond back with JSON value")
	}
}
