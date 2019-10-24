package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type Basic struct {
	Message string `json:"message"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func WriteAsJSON(w io.Writer, message string) {
	if err := json.NewEncoder(w).Encode(Basic{message}); err != nil {
		fmt.Println("Couldn't respond back with JSON value")
	}
}
