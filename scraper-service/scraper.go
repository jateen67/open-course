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

	"github.com/gocolly/colly/v2"
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

type Course struct {
	CourseCode        string `xml:"key,attr"`
	CourseTitle       string `xml:"title,attr"`
	Semester          string `xml:"ssid,attr"`
	Credits           string `xml:"credits,attr"`
	Section           string `xml:"disp,attr"`
	OpenSeats         int    `xml:"os,attr"`
	WaitlistAvailable int    `xml:"ws,attr"`
	WaitlistCapacity  int    `xml:"wc,attr"`
}

type Response struct {
	CourseCode        string `json:"course_code"`
	CourseTitle       string `json:"course_title"`
	Semester          string `json:"semester"`
	Credits           string `json:"credits"`
	Section           string `json:"section"`
	OpenSeats         int    `json:"open_seats"`
	WaitlistAvailable int    `json:"waitlist_available"`
	WaitlistCapacity  int    `json:"waitlist_capacity"`
}

func main() {
	jsonData, _ := json.MarshalIndent("", "", "\t")

	request, err := http.NewRequest("GET", "http://localhost:8081/scrapercourses", bytes.NewBuffer(jsonData))
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
	var courses []Response

	if err := json.Unmarshal(body, &courses); err != nil {
		log.Fatalln(err)
	}

	var urls []string

	for _, c := range courses {
		urls = append(urls, fmt.Sprintf("https://vsb.mcgill.ca/vsb/getclassdata.jsp?term=%s&course_0_0=%s&rq_0_0=null&t=218&e=19&nouser=1&_=1720573081517", c.Semester, c.CourseCode))
	}

	var wg sync.WaitGroup
	ch := make(chan []Course, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go scrape(&wg, url, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for courseList := range ch {
		writeToDB(courseList)
	}
}

func writeToDB(courseList []Course) {
	fmt.Println("course data successfully written to database")
}

func scrape(wg *sync.WaitGroup, url string, ch chan<- []Course) {
	defer wg.Done()

	c := colly.NewCollector()
	var _course XMLCourse
	courseList := []Course{}

	c.OnXML("//errors", func(e *colly.XMLElement) {
		err := e.ChildText("error")
		if err != "" {
			log.Fatal("Error: " + err)
		}
	})

	c.OnXML("//classdata/course", func(e *colly.XMLElement) {
		_course.CourseCode = e.Attr("key")
		_course.CourseTitle = e.Attr("title")
		_course.Semester = e.ChildAttr("uselection/selection", "ssid")
		_course.Credits = e.ChildAttr("uselection/selection", "credits")
		_course.Section = e.ChildAttrs("uselection/selection/block", "disp")
		_course.OpenSeats = e.ChildAttrs("uselection/selection/block", "os")
		_course.WaitlistAvailable = e.ChildAttrs("uselection/selection/block", "ws")
		_course.WaitlistCapacity = e.ChildAttrs("uselection/selection/block", "wc")
		for i := range _course.Section {
			var newCourse Course
			newCourse.CourseCode = _course.CourseCode
			newCourse.CourseTitle = _course.CourseTitle
			newCourse.Semester = _course.Semester
			newCourse.Credits = _course.Credits
			newCourse.Section = _course.Section[i]
			newCourse.OpenSeats, _ = strconv.Atoi(_course.OpenSeats[i])
			newCourse.WaitlistAvailable, _ = strconv.Atoi(_course.WaitlistAvailable[i])
			newCourse.WaitlistCapacity, _ = strconv.Atoi(_course.WaitlistCapacity[i])
			courseList = append(courseList, newCourse)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	ch <- courseList
}
