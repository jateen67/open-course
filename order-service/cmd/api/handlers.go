package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jateen67/order-service/internal/db"
	event "github.com/jateen67/order-service/rabbit"
)

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
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

	course, err := s.CourseDB.GetCourse(reqPayload.CourseID)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.pushToQueue(course.CourseCode, course.CourseTitle, course.Semester, course.Section)
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

	course, err := s.CourseDB.GetCourse(reqPayload.CourseID)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.pushToQueue(course.CourseCode, course.CourseTitle, course.Semester, course.Section)
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

	json.NewEncoder(w).Encode(courses)
}

func (s *server) pushToQueue(courseCode, courseTitle, semester, section string) error {
	emitter, q, err := event.NewEventEmitter(s.Rabbit)
	if err != nil {
		return err
	}

	err = emitter.Push(&q, courseCode, courseTitle, semester, section)
	if err != nil {
		return err
	}

	return nil
}
