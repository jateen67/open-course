package main

import (
	"database/sql"
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

	seed(database)

	courseDB := db.NewCourseDBImpl(database)
	orderDB := db.NewOrderDBImpl(database)
	notificationDB := db.NewNotificationDBImpl(database)
	notificationTypeDB := db.NewNotificationTypeDBImpl(database)
	srv := newServer(courseDB, orderDB, notificationDB, notificationTypeDB).Router
	log.Println("starting order service...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("order service closed")
	} else if err != nil {
		log.Println("error starting order service: ", err)
		os.Exit(1)
	}
}

func seed(database *sql.DB) {
	addCourse(database, "MATH-323", "Probability", "202409", "Lec 001", "3.0", 8, 0, 0)
	addCourse(database, "COMP-250", "Introduction to Computer Science", "202409", "Lec 001", "3.0", 10, 0, 0)
	addCourse(database, "COMP-251", "Algorithms and Data Structures", "202409", "Lec 001", "3.0", 10, 3, 6)
	addCourse(database, "COMP-273", "Introduction to Computer Systems", "202501", "Lec 001", "3.0", 10, 10, 18)
	addCourse(database, "COMP-302", "Programming Languages and Paradigms", "202409", "Lec 001", "3.0", 117, 20, 20)
	addCourse(database, "COMP-303", "Software Design", "202409", "Lec 001", "3.0", 10, 4, 5)
	addCourse(database, "COMP-421", "Database Systems", "202501", "Lec 001", "3.0", 10, 2, 15)
	addCourse(database, "SOCI-213", "Deviance", "202501", "Lec 001", "3.0", 10, 0, 0)
	addOrder(database, "danny", "dannymousa@cae.com", "5143430343", 3)
	addOrder(database, "danny", "dannymousa@cae.com", "5143430343", 2)
	addOrder(database, "danny", "dannymousa@cae.com", "5143430343", 7)
	addOrder(database, "danny", "dannymousa@cae.com", "5143430343", 1)
	addOrder(database, "rei", "reikong@gmail.com", "5143430343", 6)
	addOrder(database, "rei", "reikong@gmail.com", "5143430343", 2)
	addOrder(database, "rei", "reikong@gmail.com", "5143430343", 1)
	addOrder(database, "rei", "reikong@gmail.com", "5143430343", 4)
	addOrder(database, "p drizzy", "pdrizzy@hotmail.com", "6969696969", 8)
	addOrder(database, "jateen", "kalsijatin67@icloud.com", "4389893868", 5)
	addOrder(database, "jateen", "kalsijatin67@icloud.com", "4389893868", 2)
	addOrder(database, "jateen", "kalsijatin67@icloud.com", "4389893868", 6)
	addOrder(database, "jateen", "kalsijatin67@icloud.com", "4389893868", 4)

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
}

func addCourse(database *sql.DB, courseCode, courseTitle, semester, section, credits string, openSeats, wa, wc int) {
	courseExists, err := db.CourseExists(database, courseCode, semester, section)
	if err != nil {
		log.Fatalf("error checking if course exists: %v", err)
	}

	if !courseExists {
		err = db.CreateDefaultCourse(database, courseCode, courseTitle, semester, section, credits, openSeats, wa, wc)
		if err != nil {
			log.Fatalf("error inserting course: %v", err)
		}
		log.Println("course inserted successfully")
	} else {
		log.Println("course already inserted")
	}
}

func addOrder(database *sql.DB, name, email, phone string, courseID int) {
	orderExists, err := db.OrderExists(database, name, email, phone, courseID)
	if err != nil {
		log.Fatalf("error checking if order exists: %v", err)
	}

	if !orderExists {
		err = db.CreateDefaultOrder(database, name, email, phone, courseID)
		if err != nil {
			log.Fatalf("error inserting order: %v", err)
		}
		log.Println("order inserted successfully")
	} else {
		log.Println("order already inserted")
	}
}
