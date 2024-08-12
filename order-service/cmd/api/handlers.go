package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jateen67/order-service/internal/db"
)

type OrderPayload struct {
	ID                   int     `json:"Id"`
	CourseID             int     `json:"courseId"`
	Subject              string  `json:"subject"`
	Catalog              string  `json:"catalog"`
	CourseTitle          string  `json:"courseTitle"`
	Semester             string  `json:"semester"`
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

func (s *server) getOrderByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	order, err := s.OrderDB.GetOrder(id)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (s *server) getOrdersByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if email == "" {
		s.errorJSON(w, errors.New("no email passed in url"), http.StatusBadRequest)
		return
	}

	order, err := s.OrderDB.GetOrdersByUserEmail(email)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (s *server) getOrdersByCourseID(w http.ResponseWriter, r *http.Request) {
	courseID, err := strconv.Atoi(chi.URLParam(r, "courseId"))
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	order, err := s.OrderDB.GetOrdersByCourseID(courseID)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (s *server) getOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := s.OrderDB.GetOrders()
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (s *server) createOrder(w http.ResponseWriter, r *http.Request) {
	var reqPayload db.Order

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
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

func (s *server) editOrder(w http.ResponseWriter, r *http.Request) {
	var reqPayload db.Order

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.OrderDB.UpdateOrder(reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resPayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("user %s order for course %d updated successfully", reqPayload.Phone, reqPayload.CourseID),
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("order service: successful update")
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
		Message: fmt.Sprintf("order status for course %d turned to 0 successfully", reqPayload.CourseID),
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

// func (s *server) getAllScraperCourses(w http.ResponseWriter, r *http.Request) {
// 	orders, err := s.OrderDB.GetActiveOrders()
// 	if err != nil {
// 		s.errorJSON(w, err, http.StatusBadRequest)
// 		return
// 	}

// 	courseIDArray := make([]int, len(orders))

// 	for i, order := range orders {
// 		courseIDArray[i] = order.CourseID
// 	}

// 	courses, err := s.CourseDB.GetCoursesByMultpleIDs(courseIDArray)
// 	if err != nil {
// 		s.errorJSON(w, err, http.StatusBadRequest)
// 		return
// 	}

// 	orderMap := make(map[int][]Order)
// 	for _, order := range orders {
// 		if _, ok := orderMap[order.CourseID]; !ok {
// 			newSlice := []Order{{OrderID: order.ID, Name: order.Name, Email: order.Email, Phone: order.Phone}}
// 			orderMap[order.CourseID] = newSlice
// 		} else {
// 			orderMap[order.CourseID] = append(orderMap[order.CourseID],
// 				Order{OrderID: order.ID, Name: order.Name, Email: order.Email, Phone: order.Phone})
// 		}
// 	}

// 	var orderPayload []OrderPayload

// 	for _, course := range courses {
// 		var payload OrderPayload
// 		payload.ID = course.ID
// 		payload.CourseID = course.CourseID
// 		payload.CourseCode = course.CourseCode
// 		payload.CourseTitle = course.CourseTitle
// 		payload.Semester = course.Semester
// 		payload.ComponentCode = course.ComponentCode
// 		payload.Section = course.Section
// 		payload.OpenSeats = course.OpenSeats
// 		payload.WaitlistAvailable = course.WaitlistAvailable
// 		payload.WaitlistCapacity = course.WaitlistCapacity
// 		payload.Orders = orderMap[course.ID]
// 		orderPayload = append(orderPayload, payload)
// 	}

// 	json.NewEncoder(w).Encode(orderPayload)
// }
