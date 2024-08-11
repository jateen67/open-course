package event

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
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
				CourseID:          orderPayload.CourseID,
				CourseCode:        orderPayload.CourseCode,
				CourseTitle:       orderPayload.CourseTitle,
				Semester:          orderPayload.Semester,
				Section:           orderPayload.Section,
				OpenSeats:         orderPayload.OpenSeats,
				WaitlistAvailable: orderPayload.WaitlistAvailable,
				WaitlistCapacity:  orderPayload.WaitlistCapacity,
				Orders:            orderPayload.Orders,
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
