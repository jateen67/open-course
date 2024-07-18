package main

import (
	"fmt"
	"log"
	"net/http"
)

type orderPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	CourseID int    `json:"course_id"`
}

func (s *server) signup(w http.ResponseWriter, r *http.Request) {
	var reqPayload orderPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	id, err := s.OrderDB.CreateOrder(reqPayload.Name, reqPayload.Email, reqPayload.Phone, reqPayload.CourseID)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resPayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("user %d created successfully", id),
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("authentication service: successful signin")
}
