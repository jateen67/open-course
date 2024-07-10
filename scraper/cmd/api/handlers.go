package main

import (
	"log"
	"net/http"
)

type authenticationPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

func (s *server) authentication(w http.ResponseWriter, r *http.Request) {
	log.Println("todo implement")
}
