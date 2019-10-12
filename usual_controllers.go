package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{"Hello World. This is anonymous structure",}); err != nil {
		fmt.Println("Couldn't respond back with JSON value")
	}
}
