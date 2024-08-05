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

func (s *server) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := s.CourseDB.GetCourses()
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
		courseIDArray[i] = order.CourseID
	}

	courses, err := s.CourseDB.GetCoursesByMultpleIDs(courseIDArray)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	orderMap := make(map[db.Order]int)
	for _, order := range orders {
		orderMap[order] = order.CourseID
	}

	courseMap := make(map[int]db.Course)
	for _, course := range courses {
		courseMap[course.ID] = course
	}

	var rabbits []RabbitPayload

	for _, order := range orders {
		var payload RabbitPayload
		payload.CourseID = order.CourseID
		payload.CourseCode = courseMap[order.CourseID].CourseCode
		payload.CourseTitle = courseMap[order.CourseID].CourseTitle
		payload.Semester = courseMap[order.CourseID].Semester
		payload.Section = courseMap[order.CourseID].Section
		payload.OpenSeats = courseMap[order.CourseID].OpenSeats
		payload.WaitlistAvailable = courseMap[order.CourseID].WaitlistAvailable
		payload.WaitlistCapacity = courseMap[order.CourseID].WaitlistCapacity
		payload.OrderID = order.ID
		payload.Name = order.Name
		payload.Email = order.Email
		payload.Phone = order.Phone
		rabbits = append(rabbits, payload)
	}

	json.NewEncoder(w).Encode(rabbits)
}
