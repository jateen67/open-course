package db

import "time"

type Course struct {
	ID                int       `json:"id"`
	CourseCode        string    `json:"courseCode"`
	CourseTitle       string    `json:"courseTitle"`
	Semester          string    `json:"semester"`
	Credits           string    `json:"credits"`
	Section           string    `json:"section"`
	OpenSeats         int       `json:"openSeats"`
	WaitlistAvailable int       `json:"waitlistAvailable"`
	WaitlistCapacity  int       `json:"waitlistCapacity"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type Order struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CourseID  int       `json:"courseId"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
