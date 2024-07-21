package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	event "github.com/jateen67/order-service/rabbit"
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
	IsActive bool   `json:"is_active"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (s *server) createOrder(w http.ResponseWriter, r *http.Request) {
	var reqPayload orderCreationPayload

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

	err = s.pushToQueue()
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

	err = s.pushToQueue()
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	m := MailPayload{
		From:    "from@gmail.com",
		To:      "to@gmail.com",
		Subject: "test subject",
		Message: "test message",
	}
	err = s.sendMail(w, m)
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

func (s *server) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := s.CourseDB.GetCourses()
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(courses)
}

func (s *server) pushToQueue() error {
	emitter, q, err := event.NewEventEmitter(s.Rabbit)
	if err != nil {
		return err
	}

	err = emitter.Push(&q)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) sendMail(w http.ResponseWriter, msg MailPayload) error {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	request, err := http.NewRequest("POST", "http://mailer-service/mail", bytes.NewBuffer(jsonData))
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		s.errorJSON(w, errors.New("error calling mail service"), res.StatusCode)
		return err
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Message sent to " + msg.To + "!"

	//s.writeJSON(w, payload, http.StatusOK)

	return nil
}
