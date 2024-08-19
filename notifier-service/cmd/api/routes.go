package main

func (s *server) routes() {
	s.Router.Post("/notify", s.SendNotifications)
}
