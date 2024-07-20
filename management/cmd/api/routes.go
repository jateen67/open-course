package main

func (s *server) routes() {
	s.Router.Post("/orders", s.signup)
	s.Router.Put("/orders", s.editOrder)
}
