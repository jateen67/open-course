package event

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type emitter struct {
	connection *amqp.Connection
}

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

func (e *emitter) Push(q *amqp.Queue, payload []byte) error {
	ch, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(payload),
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s\n", payload)
	return nil
}

func NewEventEmitter(conn *amqp.Connection) (emitter, amqp.Queue, error) {
	em := emitter{
		connection: conn,
	}

	ch, err := em.connection.Channel()
	if err != nil {
		return emitter{}, amqp.Queue{}, err
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
		return emitter{}, amqp.Queue{}, err
	}

	return em, q, nil
}
