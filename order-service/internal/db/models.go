package db

import "time"

type Course struct {
	CourseID             int    `json:"courseID"`
	TermCode             int    `json:"termCode"`
	Session              string `json:"session"`
	Subject              string `json:"subject"`
	Catalog              string `json:"catalog"`
	Section              int    `json:"section"`
	ComponentCode        string `json:"componentCode"`
	ComponentDescription string `json:"componentDescription"`
	ClassNumber          int    `json:"classNumber"`
	ClassAssociation     int    `json:"classAssociation"`
	CourseTitle          string `json:"courseTitle"`
	ClassStartTime       string `json:"classStartTime"`
	ClassEndTime         string `json:"classEndTime"`
	Mondays              bool   `json:"mondays"`
	Tuesdays             bool   `json:"tuesdays"`
	Wednesdays           bool   `json:"wednesdays"`
	Thursdays            bool   `json:"thursdays"`
	Fridays              bool   `json:"fridays"`
	Saturdays            bool   `json:"saturdays"`
	Sundays              bool   `json:"sundays"`
	ClassStartDate       string `json:"classStartDate"`
	ClassEndDate         string `json:"classEndDate"`
	EnrollmentCapacity   int    `json:"enrollmentCapacity"`
	CurrentEnrollment    int    `json:"currentEnrollment"`
	WaitlistCapacity     int    `json:"waitlistCapacity"`
	CurrentWaitlistTotal int    `json:"currentWaitlistTotal"`
}

type Order struct {
	OrderID   int       `json:"orderId"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CourseID  int       `json:"courseId"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
