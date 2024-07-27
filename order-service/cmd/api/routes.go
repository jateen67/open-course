package main

func (s *server) routes() {
	s.Router.Get("/courses", s.getAllCourses)
	s.Router.Get("/scrapercourses", s.getAllScraperCourses)
	s.Router.Get("/orders", s.getOrders)
	s.Router.Get("/orderbyid/{id}", s.getOrderByID)
	s.Router.Get("/orderbyemail/{email}", s.getOrdersByEmail)
	s.Router.Get("/orderbycourseid/{courseId}", s.getOrdersByCourseID)
	s.Router.Post("/orders", s.createOrder)
	s.Router.Put("/orders", s.editOrder)
}
