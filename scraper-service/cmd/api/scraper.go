package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type OrderPayload struct {
	ClassNumber          int     `json:"classNumber"`
	Subject              string  `json:"subject"`
	Catalog              string  `json:"catalog"`
	CourseTitle          string  `json:"courseTitle"`
	TermCode             int     `json:"termCode"`
	ComponentCode        string  `json:"componentCode"`
	Section              string  `json:"section"`
	EnrollmentCapacity   int     `json:"enrollmentCapacity"`
	CurrentEnrollment    int     `json:"currentEnrollment"`
	WaitlistCapacity     int     `json:"waitlistCapacity"`
	CurrentWaitlistTotal int     `json:"currentWaitlistTotal"`
	Orders               []Order `json:"orders"`
}

type Order struct {
	OrderID int    `json:"orderId"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func scraperMain() {
	jsonData, _ := json.MarshalIndent("", "", "\t")

	request, err := http.NewRequest("GET", "http://order-service/scrapercourses", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("could not make new http request: ", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		log.Println("could not do http request: ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("error code: ", res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var orders []OrderPayload

	if err := json.Unmarshal(body, &orders); err != nil {
		log.Println(err)
		return
	}

	if len(orders) == 0 {
		return
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
		err := sendToNotifier(orderList)
		if err != nil {
			log.Println("error sending to notifier: ", err)
		}
	}

}

func scrape(wg *sync.WaitGroup, order OrderPayload, ch chan<- []OrderPayload) {
	defer wg.Done()

	orderList := []OrderPayload{}

	var term int
	if order.TermCode == 2242 {
		term = 1202430
	} else if order.TermCode == 2243 {
		term = 1202440
	} else {
		term = 1202510
	}
	unix := (time.Now().UnixMilli() / 60000) % 1000
	url := fmt.Sprintf(
		"https://vsb.concordia.ca/api/class-data?term=%v&course_0_0=%s-%s&va_0_0=&rq_0_0=&t=%v&e=%v&nouser=1",
		term,
		order.Subject,
		order.Catalog,
		unix,
		unix%3+unix%39+unix%42,
	)

	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage(url)

	errors := page.MustElements("error")
	if errors.First() != nil {
		log.Println("error fetching from vsb...")
		return
	}

	blocks := page.MustElements("block")
	for i := range blocks {
		compCode, _ := blocks[i].Attribute("type")
		sec, _ := blocks[i].Attribute("secNo")
		if *compCode == order.ComponentCode && *sec == order.Section {
			cE, _ := blocks[i].Attribute("os")
			eC, _ := blocks[i].Attribute("me")
			cW, _ := blocks[i].Attribute("ws")
			wC, _ := blocks[i].Attribute("wc")

			var newOrder OrderPayload
			newOrder.ClassNumber = order.ClassNumber
			newOrder.Subject = order.Subject
			newOrder.Catalog = order.Catalog
			newOrder.CourseTitle = order.CourseTitle
			newOrder.TermCode = order.TermCode
			newOrder.ComponentCode = order.ComponentCode
			newOrder.Section = order.Section
			newOrder.CurrentEnrollment, _ = strconv.Atoi(*cE)
			newOrder.EnrollmentCapacity, _ = strconv.Atoi(*eC)
			newOrder.CurrentWaitlistTotal, _ = strconv.Atoi(*cW)
			newOrder.WaitlistCapacity, _ = strconv.Atoi(*wC)
			newOrder.Orders = order.Orders
			orderList = append(orderList, newOrder)
		}
	}

	ch <- orderList
}

func sendToNotifier(orders []OrderPayload) error {
	ordersFiltered := filter(orders)
	jsonData, err := json.MarshalIndent(ordersFiltered, "", "\t")
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", "http://notifier-service/notify", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}

func filter(orders []OrderPayload) []OrderPayload {
	var ordersFiltered []OrderPayload
	for _, order := range orders {
		if order.CurrentEnrollment > 0 || order.CurrentWaitlistTotal > 0 {
			ordersFiltered = append(ordersFiltered, order)
		}
	}

	return ordersFiltered
}
