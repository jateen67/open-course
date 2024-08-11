package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jateen67/notifier-service/internal/data"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderPayload struct {
	CourseID          int     `json:"courseId"`
	CourseCode        string  `json:"courseCode"`
	CourseTitle       string  `json:"courseTitle"`
	Semester          string  `json:"semester"`
	Section           string  `json:"section"`
	OpenSeats         int     `json:"openSeats"`
	WaitlistAvailable int     `json:"waitlistAvailable"`
	WaitlistCapacity  int     `json:"waitlistCapacity"`
	Orders            []Order `json:"orders"`
}

type Order struct {
	OrderID int    `json:"orderId"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
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

func (s *server) ManageOrders(w http.ResponseWriter, r *http.Request) {

}

func (s *server) SendNotifications(w http.ResponseWriter, r *http.Request) {
	var reqPayload OrderPayload

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

	// == LOG NOTIFICATION TO MONGO ==
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

	// == SEND MAIL ==
	err = s.Mailer.SendSMTPMessage(msg)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// == SEND SMS ==
	twilioClient := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetFrom(os.Getenv("TWILIO_FROM_PHONE_NUMBER"))
	params.SetTo(reqPayload.Phone)
	params.SetBody(fmt.Sprintf("Hi %s,\nA seat in %s - %s (%s) has opened up for the %s semester. Sign up quickly!",
		reqPayload.Name, reqPayload.CourseCode, reqPayload.CourseTitle, reqPayload.Section, reqPayload.Semester))

	resp, err := twilioClient.Api.CreateMessage(params)
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

	// == DISABLE ORDER STATUSES SO THEY DONT GET FUTURE NOTIFS UNTIL MANUALLY SET AGAIN BY USER ==
	jsonData, _ := json.MarshalIndent(reqPayload, "", "\t")

	request, err := http.NewRequest("POST", "http://order-service/orderstatus", bytes.NewBuffer(jsonData))
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("db entry + sms/email notification sent for course %s", reqPayload.CourseID),
	}

	s.writeJSON(w, payload, http.StatusAccepted)
}
