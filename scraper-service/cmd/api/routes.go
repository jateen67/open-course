package main

func (s *server) routes() {
	s.Router.Get("/test", s.sendMessageToMailer)
}
