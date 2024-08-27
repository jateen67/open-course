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
	"strconv"

	"github.com/jateen67/order-service/internal/db"
)

type CourseAPI struct {
	ClassNumber          string `json:"classNumber"`
	CourseID             string `json:"courseID"`
	TermCode             string `json:"termCode"`
	Session              string `json:"session"`
	Subject              string `json:"subject"`
	Catalog              string `json:"catalog"`
	Section              string `json:"section"`
	ComponentCode        string `json:"componentCode"`
	ComponentDescription string `json:"componentDescription"`
	ClassAssociation     string `json:"classAssociation"`
	CourseTitle          string `json:"courseTitle"`
	ClassStartTime       string `json:"classStartTime"`
	ClassEndTime         string `json:"classEndTime"`
	Mondays              string `json:"modays"`
	Tuesdays             string `json:"tuesdays"`
	Wednesdays           string `json:"wednesdays"`
	Thursdays            string `json:"thursdays"`
	Fridays              string `json:"fridays"`
	Saturdays            string `json:"saturdays"`
	Sundays              string `json:"sundays"`
	ClassStartDate       string `json:"classStartDate"`
	ClassEndDate         string `json:"classEndDate"`
	EnrollmentCapacity   string `json:"enrollmentCapacity"`
	CurrentEnrollment    string `json:"currentEnrollment"`
	WaitlistCapacity     string `json:"waitlistCapacity"`
	CurrentWaitlistTotal string `json:"currentWaitlistTotal"`
}

type Course struct {
	ClassNumber          int    `json:"classNumber"`
	CourseID             int    `json:"courseID"`
	TermCode             int    `json:"termCode"`
	Session              string `json:"session"`
	Subject              string `json:"subject"`
	Catalog              string `json:"catalog"`
	Section              string `json:"section"`
	ComponentCode        string `json:"componentCode"`
	ComponentDescription string `json:"componentDescription"`
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
		log.Fatalf("could not connect to postgres: %s", err)
	}
	defer database.Close()

	log.Println("connected to postgres successfully")

	err = db.CreateTables(database)
	if err != nil {
		log.Fatalf("could not create tables: %s", err)
	}

	log.Println("tables created successfully")

	coursesTablePopulated, err := db.CoursesTablePopulated(database)
	if err != nil {
		log.Fatalf("error checking if courses table populated: %s", err)
	}
	if !coursesTablePopulated {
		seedCourses(database, "*", 2242) // 2242 = f2024
		seedCourses(database, "*", 2243) // 2243 = f2024/w2025
		seedCourses(database, "*", 2244) // 2244 = w2025
	}

	ordersTablePopulated, err := db.OrdersTablePopulated(database)
	if err != nil {
		log.Fatalf("error checking if orders table populated: %s", err)
	}
	if !ordersTablePopulated {
		seedOrders(database)
	}

	err = db.CreateIndexes(database)
	if err != nil {
		log.Fatalf("error creating indexes: %s", err)
	}

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
	jsonData, _ := json.MarshalIndent("", "", "\t")

	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://opendata.concordia.ca/API/v1/course/scheduleTerm/filter/%s/%v", subject, termCode),
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("could not make new http request: %s", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth("711", "d77946e392ed877022b6e0825cb36aa0")

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

	var courses []CourseAPI

	if err := json.Unmarshal(body, &courses); err != nil {
		log.Fatalln(err)
	}

	for _, course := range courses {
		addCourse(database, course)
	}

	log.Printf("all courses for %s for term %v inserted successfully", subject, termCode)
}

func seedOrders(database *sql.DB) {
	addOrder(database, "kalsijatin67@icloud.com", "+14389893868", 6399)
}

func addCourse(database *sql.DB, course CourseAPI) {
	classNumber, _ := strconv.Atoi(course.ClassNumber)
	exists, err := db.ContainsClassNumber(database, classNumber)
	if err != nil {
		log.Fatalf("error checking if class number exists: %s", err)
		return
	}
	if exists {
		return
	}

	idx := 0
	for _, r := range course.CourseID {
		if r == '0' {
			idx++
		} else {
			break
		}
	}
	courseID, _ := strconv.Atoi(course.CourseID[idx:])
	termCode, _ := strconv.Atoi(course.TermCode)
	classAssociation, _ := strconv.Atoi(course.ClassAssociation)
	enrollmentCapacity, _ := strconv.Atoi(course.EnrollmentCapacity)
	CurrentEnrollment, _ := strconv.Atoi(course.CurrentEnrollment)
	waitlistCapacity, _ := strconv.Atoi(course.WaitlistCapacity)
	currentWaitlistTotal, _ := strconv.Atoi(course.CurrentWaitlistTotal)

	err = db.CreateDefaultCourse(database, classNumber, courseID, termCode, course.Session, course.Subject, course.Catalog,
		course.Section, course.ComponentCode, course.ComponentDescription, classAssociation, course.CourseTitle,
		course.ClassStartTime, course.ClassEndTime, course.Mondays == "Y", course.Tuesdays == "Y", course.Wednesdays == "Y",
		course.Thursdays == "Y", course.Fridays == "Y", course.Saturdays == "Y", course.Sundays == "Y", course.ClassStartDate,
		course.ClassEndDate, enrollmentCapacity, CurrentEnrollment, waitlistCapacity, currentWaitlistTotal)
	if err != nil {
		log.Fatalf("error inserting course: %s", err)
	}
}

func addOrder(database *sql.DB, email, phone string, classNumber int) {
	err := db.CreateDefaultOrder(database, email, phone, classNumber)
	if err != nil {
		log.Fatalf("error inserting order: %s", err)
	}
	log.Println("order inserted successfully")
}
