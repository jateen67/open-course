package main

func (s *server) routes() {
	s.Router.Post("/mail", s.SendMail)
}
