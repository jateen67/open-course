package db

import "time"

type Course struct {
	ClassNumber          int    `json:"classNumber"`
	CourseID             int    `json:"courseId"`
	TermCode             int    `json:"termCode"`
	Session              string `json:"session"`
	Subject              string `json:"subject"`
	Catalog              string `json:"catalog"`
	Section              string `json:"section"`
	ComponentCode        string `json:"componentCode"`
	ComponentDescription string `json:"componentDescription"`
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
	ID          int       `json:"Id"`
	Phone       string    `json:"phone"`
	ClassNumber int       `json:"classNumber"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
