package main

import (
	"fmt"
	"net/http"

	"github.com/jateen67/mailer-service/internal/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JSONPayload struct {
	CourseID          int    `json:"courseId"`
	CourseCode        string `json:"courseCode"`
	CourseTitle       string `json:"courseTitle"`
	Semester          string `json:"semester"`
	Section           string `json:"section"`
	OpenSeats         int    `json:"openSeats"`
	WaitlistAvailable int    `json:"waitlistAvailable"`
	WaitlistCapacity  int    `json:"waitlistCapacity"`
	OrderID           int    `json:"orderId"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
}

func (s *server) logNotification(orderId int, notificationTypeId primitive.ObjectID) error {
	event := data.LogNotification{
		OrderID:            orderId,
		NotificationTypeId: notificationTypeId,
	}

	err := s.Models.LogNotification.Insert(event)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) SendMail(w http.ResponseWriter, r *http.Request) {
	var reqPayload JSONPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	msg := Message{
		From:    "opencourse@gmail.com",
		To:      reqPayload.Email,
		Subject: fmt.Sprintf("%s Seat Opened!", reqPayload.CourseCode),
		Data: fmt.Sprintf("Hi %s,\nA seat in %s - %s (%s) has opened up for the %s semester. Sign up quickly!",
			reqPayload.Name, reqPayload.CourseCode, reqPayload.CourseTitle, reqPayload.Section, reqPayload.Semester),
	}

	objectId, err := primitive.ObjectIDFromHex("66a862e4b2fddb9ea6768279")
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.logNotification(reqPayload.OrderID, objectId)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.Mailer.SendSMTPMessage(msg)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "db entry + email notification sent to " + reqPayload.Email,
	}

	s.writeJSON(w, payload, http.StatusAccepted)
}
