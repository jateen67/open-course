package event

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type emitter struct {
	connection *amqp.Connection
}

func (e *emitter) Push(q *amqp.Queue, courseCode, courseTitle, semester, section string) error {
	ch, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := fmt.Sprintf("%s - %s (%s) has opened up for %s. Hurry and register!", courseCode, courseTitle, section, semester)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s\n", body)
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
