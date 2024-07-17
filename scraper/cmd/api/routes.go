package main

func (s *server) routes() {
	s.Router.Post("/users", s.signup)
}
