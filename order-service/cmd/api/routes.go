package main

func (s *server) routes() {
	s.Router.Post("/", s.createOrder)
	s.Router.Put("/", s.editOrder)
}
