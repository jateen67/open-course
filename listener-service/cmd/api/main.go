package main

import (
	"fmt"
	"log"
	"os"
	"time"

	event "github.com/jateen67/listener-service/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
)

const port = "80"

func main() {
	log.Println("starting rabbitmq server...")
	conn, err := connectToRabbitMQ()
	if err != nil {
		log.Fatalf("could not connect to rabbitmq: %s", err)
	}
	defer conn.Close()

	err = event.Listen(conn)
	if err != nil {
		log.Fatalf("could not listen to queue: %s", err)
	}
}

func connectToRabbitMQ() (*amqp.Connection, error) {
	count := 0

	for {
		conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONNECTION_STRING"))
		if err != nil {
			fmt.Println("rabbitmq not yet ready...")
			count++
		} else {
			log.Println("connected to rabbitmq successfully")
			return conn, nil
		}

		if count > 10 {
			log.Println(err)
			return nil, err
		}

		log.Println("retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}
