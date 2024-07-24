package main

func (s *server) routes() {
	s.Router.Get("/courses", s.getAllCourses)
	s.Router.Get("/orders", s.getOrder)
	s.Router.Post("/orders", s.createOrder)
	s.Router.Put("/orders", s.editOrder)
}
