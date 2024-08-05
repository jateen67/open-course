package event

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
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

func Listen(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"db_change_queue", // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return err
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalf("could not set QoS: %s", err)
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("could not register consumer: %s", err)
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			split := strings.Split(string(d.Body), ";")
			//courseID, courseCode, courseTitle, semester, section, openSeats, wa, wc, orderID, name, email, phone
			courseID, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatalf("invalid data: %s", err)
			}
			openSeats, err := strconv.Atoi(split[5])
			if err != nil {
				log.Fatalf("invalid data: %s", err)
			}
			wa, err := strconv.Atoi(split[6])
			if err != nil {
				log.Fatalf("invalid data: %s", err)
			}
			wc, err := strconv.Atoi(split[7])
			if err != nil {
				log.Fatalf("invalid data: %s", err)
			}
			orderID, err := strconv.Atoi(split[8])
			if err != nil {
				log.Fatalf("invalid data: %s", err)
			}

			mail := RabbitPayload{
				CourseID:          courseID,
				CourseCode:        split[1],
				CourseTitle:       split[2],
				Semester:          split[3],
				Section:           split[4],
				OpenSeats:         openSeats,
				WaitlistAvailable: wa,
				WaitlistCapacity:  wc,
				OrderID:           orderID,
				Name:              split[9],
				Email:             split[10],
				Phone:             split[11],
			}
			err = sendMail(mail)
			if err != nil {
				log.Fatalf("could not send email: %s", err)
			}
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages")
	<-forever

	return nil
}

func sendMail(msg RabbitPayload) error {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	request, err := http.NewRequest("POST", "http://mailer-service/mail", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}
