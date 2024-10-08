package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jateen67/order-service/internal/db"
	"github.com/twilio/twilio-go/twiml"
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
	Phone   string `json:"phone"`
}

type Text struct {
	From string
	Body string
}

func (s *server) createOrder(w http.ResponseWriter, r *http.Request) {
	var reqPayload db.Order

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	if reqPayload.Phone == "" {
		s.errorJSON(w, errors.New("cannot have empty phone"), http.StatusBadRequest)
		return
	}

	id, err := s.OrderDB.CreateOrder(reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resPayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("user order %d created successfully", id),
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("order service: successful user order creation")
}

func (s *server) updateOrderStatus(w http.ResponseWriter, r *http.Request) {
	var reqPayload OrderPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var orderIds []int
	for _, o := range reqPayload.Orders {
		orderIds = append(orderIds, o.OrderID)
	}

	err = s.OrderDB.UpdateOrderStatus(orderIds)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resPayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("order status for course %d turned to 0 successfully", reqPayload.ClassNumber),
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("order service: successful isActive turn to 0")
}

func (s *server) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := s.CourseDB.GetCourses()
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(courses)
}

func (s *server) getCourseInfo(w http.ResponseWriter, r *http.Request) {
	termCode, err := strconv.Atoi(chi.URLParam(r, "termCode"))
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	courseID, err := strconv.Atoi(chi.URLParam(r, "courseId"))
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	courses, err := s.CourseDB.GetCourseInfo(courseID, termCode)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(courses)
}

func (s *server) getCourseSearch(w http.ResponseWriter, r *http.Request) {
	termCode, err := strconv.Atoi(chi.URLParam(r, "termCode"))
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	input := chi.URLParam(r, "input")

	courses, err := s.CourseDB.GetCoursesByInput(input, termCode)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(courses)
}

func (s *server) getAllScraperCourses(w http.ResponseWriter, r *http.Request) {
	orders, err := s.OrderDB.GetActiveOrders()
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	courseIDArray := make([]int, len(orders))

	for i, order := range orders {
		courseIDArray[i] = order.ClassNumber
	}

	courses, err := s.CourseDB.GetCoursesByMultpleIDs(courseIDArray)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	orderMap := make(map[int][]Order)
	for _, order := range orders {
		if _, ok := orderMap[order.ClassNumber]; !ok {
			newSlice := []Order{{OrderID: order.ID, Phone: order.Phone}}
			orderMap[order.ClassNumber] = newSlice
		} else {
			orderMap[order.ClassNumber] = append(orderMap[order.ClassNumber],
				Order{OrderID: order.ID, Phone: order.Phone})
		}
	}

	var orderPayload []OrderPayload

	for _, course := range courses {
		var payload OrderPayload
		payload.ClassNumber = course.ClassNumber
		payload.Subject = course.Subject
		payload.Catalog = course.Catalog
		payload.CourseTitle = course.CourseTitle
		payload.TermCode = course.TermCode
		payload.ComponentCode = course.ComponentCode
		payload.Section = course.Section
		payload.EnrollmentCapacity = course.EnrollmentCapacity
		payload.CurrentEnrollment = course.CurrentEnrollment
		payload.WaitlistCapacity = course.WaitlistCapacity
		payload.CurrentWaitlistTotal = course.CurrentWaitlistTotal
		payload.Orders = orderMap[course.ClassNumber]
		orderPayload = append(orderPayload, payload)
	}

	json.NewEncoder(w).Encode(orderPayload)
}

// slop
func (s *server) ManageOrders(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR - could not read SMS command: ", err)
		return
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		log.Println("ERROR - could not parse SMS command: ", err)
		return
	}
	phoneNumber := values.Get("From")

	if phoneNumber == "" {
		log.Println("ERROR - could not retrieve phone number from received twilio message: ", err)
		return
	}

	var message *twiml.MessagingMessage
	var command string
	var classNumber int

	input := strings.Split(values.Get("Body"), " ")
	if len(input) == 1 {
		if input[0] != "ORDERS" {
			message = &twiml.MessagingMessage{
				Body: "Error: Please enter a valid command",
			}
		} else {
			orders, err := s.OrderDB.GetOrdersByUserPhone(phoneNumber)
			if err != nil {
				message = &twiml.MessagingMessage{
					Body: "Internal server error: Could not retrieve all your orders. Please try again later",
				}
			} else {
				var classNumbers []int
				for _, i := range orders {
					classNumbers = append(classNumbers, i.ClassNumber)
				}
				courses, err := s.CourseDB.GetCoursesByMultpleIDs(classNumbers)
				if err != nil {
					message = &twiml.MessagingMessage{
						Body: "Internal server error: Could not retrieve all course info for your orders. Please try again later",
					}
				} else {
					var buffer bytes.Buffer
					for _, i := range courses {
						buffer.WriteString(fmt.Sprintf("(%v) %s%s - %s %s %s\n\n", i.ClassNumber, i.Subject, i.Catalog, i.CourseTitle, i.ComponentCode, i.Section))
					}
					if len(courses) == 0 {
						buffer.WriteString("You currently have no orders")
					}
					message = &twiml.MessagingMessage{
						Body: buffer.String(),
					}
				}
			}
		}
	} else {
		if (len(input) != 2) || (input[0] != "START" && input[0] != "STOP") {
			message = &twiml.MessagingMessage{
				Body: "Error: Please enter a valid command",
			}
		} else {
			if _, err := strconv.Atoi(input[1]); err != nil {
				message = &twiml.MessagingMessage{
					Body: "Error: Please enter a valid class number",
				}
			} else {
				command = input[0]
				classNumber, _ = strconv.Atoi(input[1])
				err = s.OrderDB.UpdateOrderStatusTwilio(classNumber, phoneNumber, command == "START")
				if err != nil {
					message = &twiml.MessagingMessage{
						Body: "Internal server error: Could not change order status. Please try again later",
					}
				} else {
					if command == "START" {
						message = &twiml.MessagingMessage{
							Body: "Order successfully enabled!",
						}
					} else {
						message = &twiml.MessagingMessage{
							Body: "Order successfully disabled!",
						}
					}
				}
			}
		}
	}

	twimlResult, err := twiml.Messages([]twiml.Element{message})
	if err != nil {
		log.Println("ERROR - could not send construct twilio message: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/xml")
	_, err = w.Write([]byte(twimlResult))
	if err != nil {
		log.Println("ERROR - could not retrieve phone number from received twilio message: ", err)
		return
	}
}
