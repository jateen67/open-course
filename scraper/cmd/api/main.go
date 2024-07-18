package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jateen67/scraper/internal/db"
)

const port = "80"

func main() {
	database, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("could not connect to postgres: %v", err)
	}
	defer database.Close()

	log.Println("connected to postgres successfully")

	err = db.CreateTables(database)
	if err != nil {
		log.Fatalf("could not create users table: %v", err)
	}

	log.Println("users table created successfully")

	userExists, err := db.UserExists(database, "admin@example.com")
	if err != nil {
		log.Fatalf("error checking if user exists: %v", err)
	}

	if !userExists {
		err = db.CreateDefaultUser(database, "john", "johndoe@test.com", 6789998212)
		if err != nil {
			log.Fatalf("error inserting user: %v", err)
		}
		log.Println("user inserted successfully")
	} else {
		log.Println("user already inserted")
	}

	courseDB := db.NewCourseDBImpl(database)
	orderDB := db.NewOrderDBImpl(database)
	notificationDB := db.NewNotificationDBImpl(database)
	notificationTypeDB := db.NewNotificationTypeDBImpl(database)
	srv := newServer(courseDB, orderDB, notificationDB, notificationTypeDB).Router
	log.Println("starting scraper server...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("scraper server closed")
	} else if err != nil {
		log.Println("error starting scraper server: ", err)
		os.Exit(1)
	}
}
