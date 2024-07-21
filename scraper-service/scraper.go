package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func main() {
	file, err := os.Create("courses.json")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
	}
	defer file.Close()

	urls := []string{"https://vsb.mcgill.ca/vsb/getclassdata.jsp?term=202409&course_0_0=MATH-323&rq_0_0=null&t=218&e=19&nouser=1&_=1720573081517"}

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
		writeToJSON(courseList)
	}
}

func writeToJSON(courseList []Course) {
	jsonData, err := json.MarshalIndent(courseList, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}

	file, err := os.Create("courses.json")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data successfully written to courses.json")
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
