package db

type CourseDB interface {
	GetCourses() ([]Course, error)
	GetCoursesByInput(string) ([]Course, error)
	GetCourseInfo(int) ([]Course, error)
	GetCoursesByMultpleIDs([]int) ([]Course, error)
	GetCoursesBySemester(int) ([]Course, error)
	CreateCourse(Course) (int, error)
}

type OrderDB interface {
	GetOrders() ([]Order, error)
	GetOrder(int) (*Order, error)
	GetOrdersByUserEmail(string) ([]Order, error)
	GetOrdersByUserPhone(string) ([]Order, error)
	GetOrdersByClassNumber(int) ([]Order, error)
	GetActiveOrders() ([]Order, error)
	CreateOrder(Order) (int, error)
	UpdateOrder(Order) error
	UpdateOrderStatus([]int) error
}
