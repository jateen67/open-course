package main

import (
	"net/http"

	"github.com/jateen67/mailer-service/internal/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JSONPayload struct {
	From               string `json:"from"`
	To                 string `json:"to"`
	Subject            string `json:"subject"`
	Message            string `json:"message"`
	OrderID            int    `json:"orderId"`
	NotificationTypeID string `json:"notificationTypeId"`
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
	var requestPayload JSONPayload

	err := s.readJSON(w, r, &requestPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	objectId, err := primitive.ObjectIDFromHex("66a862e4b2fddb9ea6768279")
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.logNotification(requestPayload.OrderID, objectId)
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
		Message: "db entry + email notification sent to " + requestPayload.To,
	}

	s.writeJSON(w, payload, http.StatusAccepted)
}
