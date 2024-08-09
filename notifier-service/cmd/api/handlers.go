package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jateen67/notifier-service/internal/data"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RabbitPayload struct {
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

func (s *server) ManageOrders(w http.ResponseWriter, r *http.Request) {}

func (s *server) SendNotifications(w http.ResponseWriter, r *http.Request) {
	var reqPayload RabbitPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	msg := Message{
		From:    os.Getenv("MAIL_FROM_ADDRESS"),
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

	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetFrom(os.Getenv("TWILIO_FROM_PHONE_NUMBER"))
	params.SetTo(reqPayload.Phone)
	params.SetBody(fmt.Sprintf("Hi %s,\nA seat in %s - %s (%s) has opened up for the %s semester. Sign up quickly!",
		reqPayload.Name, reqPayload.CourseCode, reqPayload.CourseTitle, reqPayload.Section, reqPayload.Semester))

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if resp.Body != nil {
			fmt.Println(*resp.Body)
		} else {
			fmt.Println(resp.Body)
		}
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("db entry + sms/email notification sent to %s (%s)", reqPayload.Email, reqPayload.Phone),
	}

	s.writeJSON(w, payload, http.StatusAccepted)
}
