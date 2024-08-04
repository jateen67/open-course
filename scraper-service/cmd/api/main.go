package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	// scraper start
	go func() {
		for {
			time.Sleep(3 * time.Second)
			scraperMain()
		}
	}()

	srv := newServer(conn).Router
	log.Println("starting scraper service...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("order service closed")
	} else if err != nil {
		log.Println("error starting order service: ", err)
		os.Exit(1)
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
