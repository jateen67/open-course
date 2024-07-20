package main

func (s *server) routes() {
	s.Router.Post("/order", s.createOrder)
	s.Router.Put("/order", s.editOrder)
}
