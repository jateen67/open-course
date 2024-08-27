package main

func (s *server) routes() {
	s.Router.Get("/courses", s.getAllCourses)
	s.Router.Get("/course/{termCode}/{courseId}", s.getCourseInfo)
	s.Router.Get("/coursesearch/{termCode}/{input}", s.getCourseSearch)
	s.Router.Get("/scrapercourses", s.getAllScraperCourses)
	s.Router.Post("/orders", s.createOrder)
	s.Router.Put("/orderstatus", s.updateOrderStatus)
	s.Router.Post("/smsmanage", s.ManageOrders)
}
