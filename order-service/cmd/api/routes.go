package main

func (s *server) routes() {
	s.Router.Get("/courses", s.getAllCourses)
	s.Router.Get("/course/{termCode}/{courseId}", s.getCourseInfo)
	s.Router.Get("/coursesearch/{termCode}/{input}", s.getCourseSearch)
	//s.Router.Get("/scrapercourses", s.getAllScraperCourses)
	s.Router.Get("/orders", s.getOrders)
	s.Router.Get("/orderbyid/{id}", s.getOrderByID)
	s.Router.Get("/ordersbyemail/{email}", s.getOrdersByEmail)
	s.Router.Get("/ordersbycourseid/{classNumber}", s.getOrdersByClassNumber)
	s.Router.Post("/orders", s.createOrder)
	s.Router.Put("/orders", s.editOrder)
	s.Router.Put("/orderstatus", s.updateOrderStatus)
	s.Router.Post("/smsmanage", s.ManageOrders)
}
