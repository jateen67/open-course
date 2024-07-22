package main

func (s *server) routes() {
	s.Router.Get("/courses", s.getAllCourses)
	s.Router.Get("/scrapercourses", s.getAllScraperCourses)
	s.Router.Post("/orders", s.createOrder)
	s.Router.Put("/orders", s.editOrder)
}
