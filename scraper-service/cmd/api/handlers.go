package main

import (
	"log"
	"net/http"

	event "github.com/jateen67/scraper-service/rabbit"
)

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (s *server) sendMessageToMailer(w http.ResponseWriter, r *http.Request) {
	err := s.pushToQueue("COMP-250", "Introduction to Computer Science", "202409", "Lec 001")
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resPayload := jsonResponse{
		Error:   false,
		Message: "message pushed to mailer service successfully",
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("scraper service: successful message push to mailer service")
}

func (s *server) pushToQueue(courseCode, courseTitle, semester, section string) error {
	emitter, q, err := event.NewEventEmitter(s.Rabbit)
	if err != nil {
		return err
	}

	err = emitter.Push(&q, courseCode, courseTitle, semester, section)
	if err != nil {
		return err
	}

	return nil
}
