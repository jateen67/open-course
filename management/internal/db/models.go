package db

import "time"

type course struct {
	ID                int       `json:"id"`
	CourseCode        string    `json:"course_code"`
	CourseTitle       string    `json:"course_title"`
	Semester          string    `json:"semester"`
	Credits           string    `json:"credits"`
	Section           string    `json:"section"`
	OpenSeats         int       `json:"open_seats"`
	WaitlistAvailable int       `json:"waitlist_available"`
	WaitlistCapacity  int       `json:"waitlist_capacity"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type order struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CourseID  int       `json:"course_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type notification struct {
	ID                 int       `json:"id"`
	OrderID            int       `json:"order_id"`
	NotificationTypeID int       `json:"notification_type_id"`
	TimeSent           time.Time `json:"time_sent"`
}

type notificationType struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}
