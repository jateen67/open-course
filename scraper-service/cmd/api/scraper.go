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

type OrderPayload struct {
	CourseID          int     `json:"courseId"`
	CourseCode        string  `json:"courseCode"`
	CourseTitle       string  `json:"courseTitle"`
	Semester          string  `json:"semester"`
	Section           string  `json:"section"`
	OpenSeats         int     `json:"openSeats"`
	WaitlistAvailable int     `json:"waitlistAvailable"`
	WaitlistCapacity  int     `json:"waitlistCapacity"`
	Orders            []Order `json:"orders"`
}

type Order struct {
	OrderID int    `json:"orderId"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
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

	var orders []OrderPayload

	if err := json.Unmarshal(body, &orders); err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	ch := make(chan []OrderPayload, len(orders))

	for _, order := range orders {
		wg.Add(1)
		go scrape(&wg, order, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for orderList := range ch {
		sendToNotifier(conn, orderList)
	}
}

func sendToNotifier(conn *amqp.Connection, orderList []OrderPayload) {
	for _, order := range orderList {
		if order.OpenSeats > 0 || order.WaitlistAvailable > 0 {
			pushToQueue(conn, order)
		}
	}
}

func scrape(wg *sync.WaitGroup, order OrderPayload, ch chan<- []OrderPayload) {
	defer wg.Done()

	url := ""
	c := colly.NewCollector()
	var _course XMLCourse
	orderList := []OrderPayload{}

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
			if _course.Section[i] == order.Section {
				var newOrder OrderPayload
				newOrder.CourseID = order.CourseID
				newOrder.CourseCode = order.CourseCode
				newOrder.CourseTitle = order.CourseTitle
				newOrder.Semester = order.Semester
				newOrder.Section = _course.Section[i]
				newOrder.OpenSeats, _ = strconv.Atoi(_course.OpenSeats[i])
				newOrder.WaitlistAvailable, _ = strconv.Atoi(_course.WaitlistAvailable[i])
				newOrder.WaitlistCapacity, _ = strconv.Atoi(_course.WaitlistCapacity[i])
				newOrder.Orders = order.Orders
				orderList = append(orderList, newOrder)
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

	ch <- orderList
}

func pushToQueue(conn *amqp.Connection, order OrderPayload) error {
	emitter, q, err := event.NewEventEmitter(conn)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err = encoder.Encode(order)
	if err != nil {
		return err
	}

	err = emitter.Push(&q, b.Bytes())
	if err != nil {
		return err
	}

	return nil
}
