package main

func (s *server) routes() {
	s.Router.Post("/authentication", s.authentication)
}
