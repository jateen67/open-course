package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	event "github.com/jateen67/order-service/rabbit"
)

type orderCreationPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	CourseID int    `json:"course_id"`
}

type orderEditPayload struct {
	Phone    string `json:"phone"`
	CourseID int    `json:"course_id"`
	IsActive bool   `json:"is_active"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (s *server) createOrder(w http.ResponseWriter, r *http.Request) {
	var reqPayload orderCreationPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	id, err := s.OrderDB.CreateOrder(reqPayload.Name, reqPayload.Email, reqPayload.Phone, reqPayload.CourseID)
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
	var reqPayload orderEditPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.OrderDB.UpdateOrder(reqPayload.Phone, reqPayload.CourseID, reqPayload.IsActive)
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
