package db

type UserDB interface {
	GetUsers() ([]*user, error)
	GetUser() (*user, error)
	GetUserByName(name string) ([]*user, error)
	GetUserByEmail(email string) (*user, error)
	GetUserByPhone(phone int) (*user, error)
	CreateUser(name, email string, phone int) error
}

type CourseDB interface {
	GetCourses() ([]*course, error)
	GetCourse() (*course, error)
	GetCourseByCourseCode(courseCode string) (*course, error)
	GetCoursesBySemester(semester string) ([]*course, error)
	GetCoursesBySection(section string) ([]*course, error)
	GetOpenCourses() ([]*course, error)
	CreateCourse(courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) error
	UpdateCourse(id, courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) error
}

type OrderDB interface {
	GetOrders() ([]*order, error)
	GetOrder() (*order, error)
	GetOrdersByUserID(userID int) (*order, error)
	GetOrdersByCourseID(courseID int) ([]*order, error)
	GetActiveOrders() ([]*order, error)
	CreateOrder(userID, courseID int) error
	UpdateOrder(id, userID, courseID int) error
}

type NotificationDB interface {
	GetNotifications() ([]*notification, error)
	GetNotification() (*notification, error)
	GetNotificationsByOrderID(orderID int) ([]*notification, error)
	GetNotificationsByNotificationTypeID(notificationTypeID int) ([]*notification, error)
	CreateNotification(orderID, notificationTypeID int) error
}

type NotificationTypeDB interface {
	GetNotificationTypes() ([]*notificationType, error)
	CreateNotificationType(t string) error
}
