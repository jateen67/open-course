package db

type CourseDB interface {
	GetCourses() ([]Course, error)
	GetCourse(courseID int) (*Course, error)
	GetCoursesByMultpleIDs(courseIDs []int) ([]Course, error)
	GetCoursesBySemester(semester int) ([]Course, error)
	CreateCourse(Course) (int, error)
}

type OrderDB interface {
	GetOrders() ([]Order, error)
	GetOrder(orderID int) (*Order, error)
	GetOrdersByUserEmail(email string) ([]Order, error)
	GetOrdersByUserPhone(phone string) ([]Order, error)
	GetOrdersByCourseID(courseID int) ([]Order, error)
	GetActiveOrders() ([]Order, error)
	CreateOrder(Order) (int, error)
	UpdateOrder(Order) error
	UpdateOrderStatus([]int) error
}
