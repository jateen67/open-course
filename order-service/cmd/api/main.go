package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jateen67/order-service/internal/db"
)

type Course struct {
	CourseID             int    `json:"courseID"`
	TermCode             int    `json:"termCode"`
	Session              string `json:"session"`
	Subject              string `json:"subject"`
	Catalog              string `json:"catalog"`
	Section              int    `json:"section"`
	ComponentCode        string `json:"componentCode"`
	ComponentDescription string `json:"componentDescription"`
	ClassNumber          int    `json:"classNumber"`
	ClassAssociation     int    `json:"classAssociation"`
	CourseTitle          string `json:"courseTitle"`
	ClassStartTime       string `json:"classStartTime"`
	ClassEndTime         string `json:"classEndTime"`
	Mondays              bool   `json:"modays"`
	Tuesdays             bool   `json:"tuesdays"`
	Wednesdays           bool   `json:"wednesdays"`
	Thursdays            bool   `json:"thursdays"`
	Fridays              bool   `json:"fridays"`
	Saturdays            bool   `json:"saturdays"`
	Sundays              bool   `json:"sundays"`
	ClassStartDate       string `json:"classStartDate"`
	ClassEndDate         string `json:"classEndDate"`
	EnrollmentCapacity   int    `json:"enrollmentCapacity"`
	CurrentEnrollment    int    `json:"currentEnrollment"`
	WaitlistCapacity     int    `json:"waitlistCapacity"`
	CurrentWaitlistTotal int    `json:"currentWaitlistTotal"`
}

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

	seedCourses(database, "COMP", 2242) // 2242 = f2024, 2244 = w2025
	seedOrders(database)

	courseDB := db.NewCourseDBImpl(database)
	orderDB := db.NewOrderDBImpl(database)
	srv := newServer(courseDB, orderDB).Router
	log.Println("starting order service...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("order service closed")
	} else if err != nil {
		log.Println("error starting order service: ", err)
		os.Exit(1)
	}
}

func seedCourses(database *sql.DB, subject string, termCode int) {
	coursesTablePopulated, err := db.CoursesTablePopulated(database)
	if err != nil {
		log.Fatalf("error checking if courses table populated: %v", err)
	}

	if !coursesTablePopulated {
		jsonData, _ := json.MarshalIndent("", "", "\t")

		request, err := http.NewRequest("GET",
			fmt.Sprintf("https://opendata.concordia.ca/API/v1/course/scheduleTerm/filter/%s/%v", subject, termCode),
			bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatalf("could not make new http request: %s", err)
		}

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			log.Fatalf("could not do http request: %s", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Fatalf("error code: %v", res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var courses []Course

		if err := json.Unmarshal(body, &courses); err != nil {
			log.Fatalln(err)
		}

		for _, course := range courses {
			addCourse(database, course)
		}

		log.Printf("all courses for %s for term %v inserted successfully", subject, termCode)
	}
}

func seedOrders(database *sql.DB) {
	ordersTablePopulated, err := db.OrdersTablePopulated(database)
	if err != nil {
		log.Fatalf("error checking if orders table populated: %v", err)
	}

	if !ordersTablePopulated {
		addOrder(database, "dannymousa@cae.com", "5143430343", 1)
		addOrder(database, "dannymousa@cae.com", "5143430343", 10)
		addOrder(database, "reikong@gmail.com", "5143430343", 132)
		addOrder(database, "reikong@gmail.com", "5143430343", 45)
		addOrder(database, "reikong@gmail.com", "5143430343", 165)
		addOrder(database, "kalsijatin67@icloud.com", "4389893868", 44)
		addOrder(database, "kalsijatin67@icloud.com", "4389893868", 45)
	}
}

func addCourse(database *sql.DB, course Course) {
	err := db.CreateDefaultCourse(database, course.CourseID, course.TermCode, course.Session, course.Subject, course.Catalog,
		course.Section, course.ComponentCode, course.ComponentDescription, course.ClassNumber, course.ClassAssociation,
		course.CourseTitle, course.ClassStartTime, course.ClassEndTime, course.Mondays, course.Tuesdays, course.Wednesdays,
		course.Thursdays, course.Fridays, course.Saturdays, course.Sundays, course.ClassStartDate, course.ClassEndDate,
		course.EnrollmentCapacity, course.CurrentEnrollment, course.WaitlistCapacity, course.CurrentWaitlistTotal)
	if err != nil {
		log.Fatalf("error inserting course: %v", err)
	}
}

func addOrder(database *sql.DB, email, phone string, FK_courseID int) {
	err := db.CreateDefaultOrder(database, email, phone, FK_courseID)
	if err != nil {
		log.Fatalf("error inserting order: %v", err)
	}
	log.Println("order inserted successfully")
}
