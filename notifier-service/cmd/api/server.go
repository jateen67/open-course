package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jateen67/notifier-service/internal/data"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	Router chi.Router
	Models data.Models
}

func newServer(client *mongo.Client) *server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
	}))

	s := &server{
		Router: r,
		Models: data.New(client),
	}
	s.routes()

	return s
}
