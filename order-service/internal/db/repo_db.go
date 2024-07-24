package db

type CourseDB interface {
	GetCourses() ([]Course, error)
	GetCourse(courseID int) (*Course, error)
	GetCourseByCourseCode(courseCode string) (*Course, error)
	GetCoursesBySemester(semester string) ([]Course, error)
	GetCoursesBySection(section string) ([]Course, error)
	GetOpenCourses() ([]Course, error)
	CreateCourse(Course) (int, error)
	UpdateCourse(Course) error
}

type OrderDB interface {
	GetOrders() ([]Order, error)
	GetOrder(orderID int) (*Order, error)
	GetOrdersByUserName(name string) ([]Order, error)
	GetOrdersByUserEmail(email string) ([]Order, error)
	GetOrdersByUserPhone(phone string) ([]Order, error)
	GetOrdersByCourseID(courseID int) ([]Order, error)
	GetActiveOrders() ([]Order, error)
	CreateOrder(Order) (int, error)
	UpdateOrder(Order) error
}

type NotificationDB interface {
	GetNotifications() ([]Notification, error)
	GetNotification(notificationID int) (*Notification, error)
	GetNotificationsByOrderID(orderID int) ([]Notification, error)
	GetNotificationsByNotificationTypeID(notificationTypeID int) ([]Notification, error)
	CreateNotification(Notification) (int, error)
}

type NotificationTypeDB interface {
	GetNotificationTypes() ([]NotificationType, error)
	CreateNotificationType(NotificationType) (int, error)
}
