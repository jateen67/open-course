package main

func (s *server) routes() {
	s.Router.Get("/courses", s.getAllCourses)
	s.Router.Get("/orders", s.getOrders)
	s.Router.Get("/orders/{id}", s.getOrderByID)
	s.Router.Get("/orders/{email}", s.getOrdersByEmail)
	s.Router.Get("/orders/{courseId}", s.getOrdersByCourseID)
	s.Router.Post("/orders", s.createOrder)
	s.Router.Put("/orders", s.editOrder)
}
