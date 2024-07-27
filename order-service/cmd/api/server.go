package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jateen67/order-service/internal/db"
)

type server struct {
	Router             chi.Router
	CourseDB           db.CourseDB
	OrderDB            db.OrderDB
	NotificationDB     db.NotificationDB
	NotificationTypeDB db.NotificationTypeDB
}

func newServer(
	courseDB db.CourseDB,
	orderDB db.OrderDB,
	notificationDB db.NotificationDB,
	notificationTypeDB db.NotificationTypeDB) *server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
	}))

	s := &server{
		Router:             r,
		CourseDB:           courseDB,
		OrderDB:            orderDB,
		NotificationDB:     notificationDB,
		NotificationTypeDB: notificationTypeDB,
	}
	s.routes()

	return s
}
