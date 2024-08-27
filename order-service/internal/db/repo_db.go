package db

type CourseDB interface {
	GetCourses() ([]Course, error)
	GetCoursesByInput(string, int) ([]Course, error)
	GetCourseInfo(int, int) ([]Course, error)
	GetCoursesByMultpleIDs([]int) ([]Course, error)
	GetCoursesBySemester(int) ([]Course, error)
}

type OrderDB interface {
	GetOrdersByUserPhone(string) ([]Order, error)
	GetActiveOrders() ([]Order, error)
	CreateOrder(Order) (int, error)
	UpdateOrderStatus([]int) error
	UpdateOrderStatusTwilio(int, string, bool) error
}
