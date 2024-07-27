package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type server struct {
	Router chi.Router
	Rabbit *amqp.Connection
}

func newServer(c *amqp.Connection) *server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
	}))

	s := &server{
		Router: r,
		Rabbit: c,
	}
	s.routes()

	return s
}
