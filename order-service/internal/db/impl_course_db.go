package db

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type CourseDBImpl struct {
	DB *sql.DB
}

func NewCourseDBImpl(db *sql.DB) *CourseDBImpl {
	return &CourseDBImpl{DB: db}
}

func (d *CourseDBImpl) GetCourses() ([]Course, error) {
	query := "SELECT * FROM tbl_Courses"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.CourseCode, &course.CourseTitle,
			&course.Semester, &course.Credits, &course.Section, &course.OpenSeats,
			&course.WaitlistAvailable, &course.WaitlistCapacity, &course.CreatedAt,
			&course.UpdatedAt); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func (d *CourseDBImpl) GetCourse(courseID int) (*Course, error) {
	query := "SELECT * FROM tbl_Courses where id = $1"
	var course Course
	if err := d.DB.QueryRow(query, courseID).Scan(
		&course.ID,
		&course.CourseCode,
		&course.CourseTitle,
		&course.Semester,
		&course.Section,
		&course.Credits,
		&course.OpenSeats,
		&course.WaitlistAvailable,
		&course.WaitlistCapacity,
		&course.CreatedAt,
		&course.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &course, nil
}

func (d *CourseDBImpl) GetCoursesByMultpleIDs(courseIDs []int) ([]Course, error) {
	query := "SELECT * FROM tbl_Courses WHERE id = ANY($1)"
	rows, err := d.DB.Query(query, pq.Array(courseIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.CourseCode, &course.CourseTitle,
			&course.Semester, &course.Credits, &course.Section, &course.OpenSeats,
			&course.WaitlistAvailable, &course.WaitlistCapacity, &course.CreatedAt,
			&course.UpdatedAt); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func (d *CourseDBImpl) GetCourseByCourseCode(courseCode string) (*Course, error) {
	query := "SELECT * FROM tbl_Courses WHERE courseCode = $1"
	var course Course
	if err := d.DB.QueryRow(query, courseCode).Scan(&course.ID, &course.CourseCode, &course.CourseTitle,
		&course.Semester, &course.Credits, &course.Section, &course.OpenSeats,
		&course.WaitlistAvailable, &course.WaitlistCapacity, &course.CreatedAt,
		&course.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &course, nil
}

func (d *CourseDBImpl) GetCoursesBySemester(semester string) ([]Course, error) {
	query := "SELECT * FROM tbl_Courses WHERE semester = $1"
	rows, err := d.DB.Query(query, semester)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.CourseCode, &course.CourseTitle,
			&course.Semester, &course.Credits, &course.Section, &course.OpenSeats,
			&course.WaitlistAvailable, &course.WaitlistCapacity, &course.CreatedAt,
			&course.UpdatedAt); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func (d *CourseDBImpl) GetCoursesBySection(section string) ([]Course, error) {
	query := "SELECT * FROM tbl_Courses WHERE section = $1"
	rows, err := d.DB.Query(query, section)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.CourseCode, &course.CourseTitle,
			&course.Semester, &course.Credits, &course.Section, &course.OpenSeats,
			&course.WaitlistAvailable, &course.WaitlistCapacity, &course.CreatedAt,
			&course.UpdatedAt); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func (d *CourseDBImpl) GetOpenCourses() ([]Course, error) {
	query := "SELECT * FROM tbl_Courses WHERE openSeats > 0"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.CourseCode, &course.CourseTitle,
			&course.Semester, &course.Credits, &course.Section, &course.OpenSeats,
			&course.WaitlistAvailable, &course.WaitlistCapacity, &course.CreatedAt,
			&course.UpdatedAt); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func (d *CourseDBImpl) CreateCourse(course Course) (int, error) {
	var id int
	query := `INSERT INTO tbl_Courses (
		courseCode,
		courseTitle,
		semester,
		credits,
		section,
		openSeats,
		waitlistAvailable,
		waitlistCapacity,
		createdAt,
		updatedAt
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err := d.DB.QueryRow(query, course.CourseCode, course.CourseTitle,
		course.Semester, course.Credits, course.Section, course.OpenSeats,
		course.WaitlistAvailable, course.WaitlistCapacity, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *CourseDBImpl) UpdateCourse(course Course) error {
	query := `UPDATE tbl_Courses SET 
		courseCode = $2,
		courseTitle = $3,
		semester = $4,
		credits = $5,
		section = $6,
		openSeats = $7,
		waitlistAvailable = $8,
		waitlistCapacity = $9,
		updatedAt = $10
		WHERE id = $1`
	_, err := d.DB.Exec(query, course.ID, course.CourseCode, course.CourseTitle,
		course.Semester, course.Credits, course.Section, course.OpenSeats,
		course.WaitlistAvailable, course.WaitlistCapacity, time.Now())
	if err != nil {
		return err
	}

	return nil
}
