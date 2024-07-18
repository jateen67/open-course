package db

type CourseDB interface {
	GetCourses() ([]course, error)
	GetCourse(courseID int) (*course, error)
	GetCourseByCourseCode(courseCode string) (*course, error)
	GetCoursesBySemester(semester string) ([]course, error)
	GetCoursesBySection(section string) ([]course, error)
	GetOpenCourses() ([]course, error)
	CreateCourse(courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) (int64, error)
	UpdateCourse(id int, courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) error
}

type OrderDB interface {
	GetOrders() ([]order, error)
	GetOrder(orderID int) (*order, error)
	GetOrdersByUserName(name string) ([]order, error)
	GetOrdersByUserEmail(email string) ([]order, error)
	GetOrdersByUserPhone(phone int) ([]order, error)
	GetOrdersByCourseID(courseID int) ([]order, error)
	GetActiveOrders() ([]order, error)
	CreateOrder(name, email string, phone, courseID int) (int64, error)
	UpdateOrder(name, email string, phone int, courseID int, isActive bool) error
}

type NotificationDB interface {
	GetNotifications() ([]notification, error)
	GetNotification(notificationID int) (*notification, error)
	GetNotificationsByOrderID(orderID int) ([]notification, error)
	GetNotificationsByNotificationTypeID(notificationTypeID int) ([]notification, error)
	CreateNotification(orderID, notificationTypeID int) (int64, error)
}

type NotificationTypeDB interface {
	GetNotificationTypes() ([]notificationType, error)
	CreateNotificationType(t string) (int64, error)
}
