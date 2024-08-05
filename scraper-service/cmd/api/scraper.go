package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gocolly/colly/v2"
	event "github.com/jateen67/scraper-service/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
)

type XMLCourse struct {
	CourseCode        string   `xml:"key,attr"`
	CourseTitle       string   `xml:"title,attr"`
	Semester          string   `xml:"ssid,attr"`
	Credits           string   `xml:"credits,attr"`
	Section           []string `xml:"disp,attr"`
	OpenSeats         []string `xml:"os,attr"`
	WaitlistAvailable []string `xml:"ws,attr"`
	WaitlistCapacity  []string `xml:"wc,attr"`
}

type RabbitPayload struct {
	CourseID          int    `json:"courseId"`
	CourseCode        string `json:"courseCode"`
	CourseTitle       string `json:"courseTitle"`
	Semester          string `json:"semester"`
	Section           string `json:"section"`
	OpenSeats         int    `json:"openSeats"`
	WaitlistAvailable int    `json:"waitlistAvailable"`
	WaitlistCapacity  int    `json:"waitlistCapacity"`
	OrderID           int    `json:"orderId"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
}

func scraperMain(conn *amqp.Connection) {
	jsonData, _ := json.MarshalIndent("", "", "\t")

	request, err := http.NewRequest("GET", "http://order-service/scrapercourses", bytes.NewBuffer(jsonData))
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

	var rabbits []RabbitPayload

	if err := json.Unmarshal(body, &rabbits); err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	ch := make(chan []RabbitPayload, len(rabbits))

	for _, rabbit := range rabbits {
		wg.Add(1)
		go scrape(&wg, rabbit, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for rabbitList := range ch {
		sendToMailer(conn, rabbitList)
	}
}

func sendToMailer(conn *amqp.Connection, rabbitList []RabbitPayload) {
	for _, rabbit := range rabbitList {
		if rabbit.OpenSeats > 0 || rabbit.WaitlistAvailable > 0 {
			pushToQueue(conn, rabbit)
		}
	}
}

func scrape(wg *sync.WaitGroup, rabbit RabbitPayload, ch chan<- []RabbitPayload) {
	defer wg.Done()

	url := ""
	c := colly.NewCollector()
	var _course XMLCourse
	rabbitList := []RabbitPayload{}

	c.OnXML("//errors", func(e *colly.XMLElement) {
		err := e.ChildText("error")
		if err != "" {
			log.Fatal("error: " + err)
		}
	})

	c.OnXML("//classdata/course", func(e *colly.XMLElement) {
		_course.Section = e.ChildAttrs("uselection/selection/block", "disp")
		_course.OpenSeats = e.ChildAttrs("uselection/selection/block", "os")
		_course.WaitlistAvailable = e.ChildAttrs("uselection/selection/block", "ws")
		_course.WaitlistCapacity = e.ChildAttrs("uselection/selection/block", "wc")
		for i := range _course.Section {
			if _course.Section[i] == rabbit.Section {
				var newRabbit RabbitPayload
				newRabbit.CourseID = rabbit.CourseID
				newRabbit.CourseCode = rabbit.CourseCode
				newRabbit.CourseTitle = rabbit.CourseTitle
				newRabbit.Semester = rabbit.Semester
				newRabbit.Section = _course.Section[i]
				newRabbit.OpenSeats, _ = strconv.Atoi(_course.OpenSeats[i])
				newRabbit.WaitlistAvailable, _ = strconv.Atoi(_course.WaitlistAvailable[i])
				newRabbit.WaitlistCapacity, _ = strconv.Atoi(_course.WaitlistCapacity[i])
				newRabbit.OrderID = rabbit.OrderID
				newRabbit.Name = rabbit.Name
				newRabbit.Email = rabbit.Email
				newRabbit.Phone = rabbit.Phone
				rabbitList = append(rabbitList, newRabbit)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting: ", r.URL)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	ch <- rabbitList
}

func pushToQueue(conn *amqp.Connection, rabbit RabbitPayload) error {
	emitter, q, err := event.NewEventEmitter(conn)
	if err != nil {
		return err
	}

	err = emitter.Push(&q, rabbit.CourseID, rabbit.CourseCode, rabbit.CourseTitle, rabbit.Semester,
		rabbit.Section, rabbit.OpenSeats, rabbit.WaitlistAvailable, rabbit.WaitlistCapacity, rabbit.OrderID,
		rabbit.Name, rabbit.Email, rabbit.Phone)
	if err != nil {
		return err
	}

	return nil
}
