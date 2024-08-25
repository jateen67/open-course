package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type OrderPayload struct {
	ClassNumber          int     `json:"classNumber"`
	Subject              string  `json:"subject"`
	Catalog              string  `json:"catalog"`
	CourseTitle          string  `json:"courseTitle"`
	TermCode             int     `json:"termCode"`
	ComponentCode        string  `json:"componentCode"`
	Section              string  `json:"section"`
	EnrollmentCapacity   int     `json:"enrollmentCapacity"`
	CurrentEnrollment    int     `json:"currentEnrollment"`
	WaitlistCapacity     int     `json:"waitlistCapacity"`
	CurrentWaitlistTotal int     `json:"currentWaitlistTotal"`
	Orders               []Order `json:"orders"`
}

type Order struct {
	OrderID int    `json:"orderId"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func (s *server) logNotification(orderIDs []int, notificationType string) error {
	err := s.Models.LogNotification.Insert(orderIDs, notificationType)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) SendNotifications(w http.ResponseWriter, r *http.Request) {
	var reqPayloads []OrderPayload

	err := s.readJSON(w, r, &reqPayloads)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	for _, reqPayload := range reqPayloads {

		var emails []string
		for _, i := range reqPayload.Orders {
			emails = append(emails, i.Email)
		}

		var orderIDs []int
		for _, i := range reqPayload.Orders {
			orderIDs = append(orderIDs, i.OrderID)
		}

		msg := Message{
			From:    os.Getenv("MAIL_FROM_ADDRESS"),
			To:      emails,
			Subject: fmt.Sprintf("%s-%s Seat Opened!", reqPayload.Subject, reqPayload.Catalog),
			Data: fmt.Sprintf("Hi,\nA seat in %s-%s - %s (%s %s) has opened up for the %v semester. Sign up quickly!",
				reqPayload.Subject, reqPayload.Catalog, reqPayload.CourseTitle, reqPayload.ComponentCode,
				reqPayload.Section, reqPayload.TermCode),
		}

		// == SEND MAIL ==
		err = s.Mailer.SendSMTPMessage(msg)
		if err != nil {
			s.errorJSON(w, err, http.StatusBadRequest)
			return
		}

		// == LOG ALL EMAIL NOTIFICATIONS TO MONGO ==
		err = s.logNotification(orderIDs, "Email")
		if err != nil {
			s.errorJSON(w, err, http.StatusBadRequest)
			return
		}

		// == SEND SMS ==
		twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: os.Getenv("TWILIO_ACCOUNT_SID"),
			Password: os.Getenv("TWILIO_AUTH_TOKEN"),
		})

		for _, i := range reqPayload.Orders {
			params := &twilioApi.CreateMessageParams{}
			params.SetFrom(os.Getenv("TWILIO_FROM_PHONE_NUMBER"))
			params.SetTo(i.Phone)
			params.SetBody(fmt.Sprintf("Hi,\nA seat in %s-%s - %s (%s %s) has opened up for the %v semester. Sign up quickly!",
				reqPayload.Subject, reqPayload.Catalog, reqPayload.CourseTitle, reqPayload.ComponentCode,
				reqPayload.Section, reqPayload.TermCode))

			_, err := twilioClient.Api.CreateMessage(params)
			if err != nil {
				log.Println(err.Error())
			} else {
				log.Printf("sent sms to %s", i.Phone)
			}
		}

		// == LOG ALL SMS NOTIFICATIONS TO MONGO ==
		err = s.logNotification(orderIDs, "SMS")
		if err != nil {
			s.errorJSON(w, err, http.StatusBadRequest)
			return
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
			Message: fmt.Sprintf("db entry + order update + sms/email notification sent for course %v", reqPayload.ClassNumber),
		}

		s.writeJSON(w, payload, http.StatusAccepted)
	}
}
