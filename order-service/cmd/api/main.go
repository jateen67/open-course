package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jateen67/order-service/internal/db"
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
		log.Fatalf("could not create tables: %v", err)
	}

	log.Println("tables created successfully")

	courseExists, err := db.CourseExists(database, "MATH-323", "202409", "Lec 001")
	if err != nil {
		log.Fatalf("error checking if course exists: %v", err)
	}

	if !courseExists {
		err = db.CreateDefaultCourse(database, "MATH-323", "Probability", "202409", "Lec 001", "3.0", 8, 0, 0)
		if err != nil {
			log.Fatalf("error inserting course: %v", err)
		}
		log.Println("course inserted successfully")
	} else {
		log.Println("course already inserted")
	}

	orderExists, err := db.OrderExists(database, "john", "johndoe@test.com", "6789998212", 1)
	if err != nil {
		log.Fatalf("error checking if order exists: %v", err)
	}

	if !orderExists {
		err = db.CreateDefaultOrder(database, "john", "johndoe@test.com", "6789998212", 1)
		if err != nil {
			log.Fatalf("error inserting order: %v", err)
		}
		log.Println("order inserted successfully")
	} else {
		log.Println("order already inserted")
	}

	nTypeExists, err := db.NotificationTypeExists(database, "Open Seat")
	if err != nil {
		log.Fatalf("error checking if notification type exists: %v", err)
	}

	if !nTypeExists {
		err = db.CreateDefaultNotificationType(database, "Open Seat")
		if err != nil {
			log.Fatalf("error inserting notification type: %v", err)
		}
		log.Println("notification type inserted successfully")
	} else {
		log.Println("notification type already inserted")
	}

	courseDB := db.NewCourseDBImpl(database)
	orderDB := db.NewOrderDBImpl(database)
	notificationDB := db.NewNotificationDBImpl(database)
	notificationTypeDB := db.NewNotificationTypeDBImpl(database)
	srv := newServer(courseDB, orderDB, notificationDB, notificationTypeDB).Router
	log.Println("starting order server...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("order server closed")
	} else if err != nil {
		log.Println("error starting order server: ", err)
		os.Exit(1)
	}
}
