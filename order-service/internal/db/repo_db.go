package db

type CourseDB interface {
	GetCourses() ([]Course, error)
	GetCourse(courseID int) (*Course, error)
	GetCourseByCourseCode(courseCode string) (*Course, error)
	GetCoursesBySemester(semester string) ([]Course, error)
	GetCoursesBySection(section string) ([]Course, error)
	GetOpenCourses() ([]Course, error)
	CreateCourse(courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) (int, error)
	UpdateCourse(id int, courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) error
}

type OrderDB interface {
	GetOrders() ([]Order, error)
	GetOrder(orderID int) (*Order, error)
	GetOrdersByUserName(name string) ([]Order, error)
	GetOrdersByUserEmail(email string) ([]Order, error)
	GetOrdersByUserPhone(phone string) ([]Order, error)
	GetOrdersByCourseID(courseID int) ([]Order, error)
	GetActiveOrders() ([]Order, error)
	CreateOrder(name, email, phone string, courseID int) (int, error)
	UpdateOrder(phone string, courseID int, isActive bool) error
}

type NotificationDB interface {
	GetNotifications() ([]Notification, error)
	GetNotification(notificationID int) (*Notification, error)
	GetNotificationsByOrderID(orderID int) ([]Notification, error)
	GetNotificationsByNotificationTypeID(notificationTypeID int) ([]Notification, error)
	CreateNotification(orderID, notificationTypeID int) (int, error)
}

type NotificationTypeDB interface {
	GetNotificationTypes() ([]NotificationType, error)
	CreateNotificationType(t string) (int, error)
}
