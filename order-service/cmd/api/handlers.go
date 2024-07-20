package main

import (
	"fmt"
	"log"
	"net/http"
)

type orderCreationPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	CourseID int    `json:"course_id"`
}

type orderEditPayload struct {
	Phone    string `json:"phone"`
	CourseID int    `json:"course_id"`
	IsActive bool   `json:"is_empty"`
}

func (s *server) createOrder(w http.ResponseWriter, r *http.Request) {
	var reqPayload orderCreationPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// add course to tbl_Courses

	id, err := s.OrderDB.CreateOrder(reqPayload.Name, reqPayload.Email, reqPayload.Phone, reqPayload.CourseID)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resPayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("user order %d created successfully", id),
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("order service: successful user order creation")
}

func (s *server) editOrder(w http.ResponseWriter, r *http.Request) {
	var reqPayload orderEditPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.OrderDB.UpdateOrder(reqPayload.Phone, reqPayload.CourseID, reqPayload.IsActive)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resPayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("user %s order for course %d updated successfully", reqPayload.Phone, reqPayload.CourseID),
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("order service: successful update")
}
