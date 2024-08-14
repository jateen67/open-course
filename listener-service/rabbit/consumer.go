package event

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderPayload struct {
	ID                   int     `json:"Id"`
	ClassNumber          int     `json:"classNumber"`
	Subject              string  `json:"subject"`
	Catalog              string  `json:"catalog"`
	CourseTitle          string  `json:"courseTitle"`
	Semester             string  `json:"semester"`
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
			var orderPayload OrderPayload
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err := decoder.Decode(&orderPayload)
			if err != nil {
				log.Fatalf("error deserializing message: %s", err)
			}

			notifInfo := OrderPayload{
				ID:                   orderPayload.ID,
				ClassNumber:          orderPayload.ClassNumber,
				Subject:              orderPayload.Subject,
				Catalog:              orderPayload.Catalog,
				CourseTitle:          orderPayload.CourseTitle,
				Semester:             orderPayload.Semester,
				ComponentCode:        orderPayload.ComponentCode,
				Section:              orderPayload.Section,
				EnrollmentCapacity:   orderPayload.EnrollmentCapacity,
				CurrentEnrollment:    orderPayload.CurrentEnrollment,
				CurrentWaitlistTotal: orderPayload.CurrentWaitlistTotal,
				WaitlistCapacity:     orderPayload.WaitlistCapacity,
				Orders:               orderPayload.Orders,
			}
			err = sendNotification(notifInfo)
			if err != nil {
				log.Fatalf("could not send notification: %s", err)
			}
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages")
	<-forever

	return nil
}

func sendNotification(msg OrderPayload) error {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	request, err := http.NewRequest("POST", "http://notifier-service/notify", bytes.NewBuffer(jsonData))
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
