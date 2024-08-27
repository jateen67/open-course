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
	Phone   string `json:"phone"`
}

func (s *server) logNotification(orderID int, notificationType string) error {
	err := s.Models.LogNotification.Insert(orderID, notificationType)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) SendNotifications(w http.ResponseWriter, r *http.Request) {
	var reqPayloads []OrderPayload

	err := s.readJSON(w, r, &reqPayloads)
	if err != nil {
		log.Println("ERROR - could not read /notify request body: ", err)
		return
	}

	for _, reqPayload := range reqPayloads {
		terms := map[int]string{
			2242: "Fall 2024",
			2243: "Fall 2024/Winter 2025",
			2244: "Winter 2025",
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
			params.SetBody(fmt.Sprintf("Hi,\n\nA seat in %s-%s - %s (%s %s) has opened up for %s. Sign up quickly!\n\n-OpenCourse",
				reqPayload.Subject, reqPayload.Catalog, reqPayload.CourseTitle, reqPayload.ComponentCode,
				reqPayload.Section, terms[reqPayload.TermCode]))

			_, err := twilioClient.Api.CreateMessage(params)
			if err != nil {
				log.Printf("ERROR - could not send SMS for order %v of class %v: %s\n", i.OrderID, reqPayload.ClassNumber, err.Error())
			} else {
				err = s.logNotification(i.OrderID, "SMS")
				if err != nil {
					log.Printf("ERROR - could not log notif for order %v of class %v: %s\n", i.OrderID, reqPayload.ClassNumber, err)
				}
			}
		}

		// == DISABLE ORDER STATUSES SO THEY DONT GET FUTURE NOTIFS UNTIL MANUALLY SET AGAIN BY USER ==
		jsonData, _ := json.MarshalIndent(reqPayload, "", "\t")

		request, err := http.NewRequest("POST", "http://order-service/orderstatus", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("ERROR - could not make request to disable all orders of class %v: %s\n", reqPayload.ClassNumber, err)
			continue
		}

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			log.Printf("ERROR - could not do request to disable all orders of class %v: %s\n", reqPayload.ClassNumber, err)
			continue
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusAccepted {
			log.Printf("ERROR - could not disable all orders of class %v -- status %v\n", reqPayload.ClassNumber, res.StatusCode)
			continue
		}

		var jsonFromService jsonResponse

		err = json.NewDecoder(res.Body).Decode(&jsonFromService)
		if err != nil {
			log.Printf("ERROR - could not decode /orderstatus response: %s\n", jsonFromService.Message)
			continue
		}

		if jsonFromService.Error {
			log.Printf("ERROR - could not disable all orders of class %v: %s\n", reqPayload.ClassNumber, jsonFromService.Message)
		}
	}

	payload := jsonResponse{
		Error:   false,
		Message: "db entries + order updates + sms notifications sent",
	}

	s.writeJSON(w, payload, http.StatusAccepted)
}
