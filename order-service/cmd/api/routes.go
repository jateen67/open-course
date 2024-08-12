package main

func (s *server) routes() {
	s.Router.Get("/courses", s.getAllCourses)
	//s.Router.Get("/scrapercourses", s.getAllScraperCourses)
	s.Router.Get("/orders", s.getOrders)
	s.Router.Get("/orderbyid/{id}", s.getOrderByID)
	s.Router.Get("/ordersbyemail/{email}", s.getOrdersByEmail)
	s.Router.Get("/ordersbycourseid/{courseId}", s.getOrdersByCourseID)
	s.Router.Post("/orders", s.createOrder)
	s.Router.Put("/orders", s.editOrder)
	s.Router.Put("/orderstatus", s.updateOrderStatus)
}
