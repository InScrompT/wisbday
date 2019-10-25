package main

import (
	"html"
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
		Email: html.EscapeString(r.FormValue("email")), // Could be either a username or an email ID
		Password: html.EscapeString(r.FormValue("password")),
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
	hashedPassword, err := HashPassword(r.FormValue("password"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		WriteAsJSON(w, "Unable to create an account")

		return
	}

	DB.Create(User{
		Email: html.EscapeString(r.FormValue("email")),
		Username: html.EscapeString(r.FormValue("username")),
		Password: hashedPassword,
	})

	WriteAsJSON(w, "User created successfully")
}
