package db

import (
	"database/sql"
	"time"
)

type CourseDBImpl struct {
	DB *sql.DB
}

func NewCourseDBImpl(db *sql.DB) *CourseDBImpl {
	return &CourseDBImpl{DB: db}
}

func (d *CourseDBImpl) GetCourses() ([]course, error) {
	query := "SELECT * FROM tbl_Courses"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []course

	for rows.Next() {
		var course course
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

func (d *CourseDBImpl) GetCourse(courseID int) (*course, error) {
	query := "SELECT * FROM tbl_Courses where id = $1"
	var course course
	if err := d.DB.QueryRow(query, courseID).Scan(&course); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &course, nil
}

func (d *CourseDBImpl) GetCourseByCourseCode(courseCode string) (*course, error) {
	query := "SELECT * FROM tbl_Users WHERE course_code = $1"
	var course course
	if err := d.DB.QueryRow(query, courseCode).Scan(&course); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &course, nil
}

func (d *CourseDBImpl) GetCoursesBySemester(semester string) ([]course, error) {
	query := "SELECT * FROM tbl_Users WHERE semester = $1"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []course

	for rows.Next() {
		var course course
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

func (d *CourseDBImpl) GetCoursesBySection(section string) ([]course, error) {
	query := "SELECT * FROM tbl_Users WHERE section = $1"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []course

	for rows.Next() {
		var course course
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

func (d *CourseDBImpl) GetOpenCourses() ([]course, error) {
	query := "SELECT * FROM tbl_Users WHERE open_seats > 0"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []course

	for rows.Next() {
		var course course
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

func (d *CourseDBImpl) CreateCourse(courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) (int64, error) {
	query := `INSERT INTO tbl_Courses (
	course_code,
	course_title,
	semester,
	credits,
	section,
	open_seats,
	waitlist_available,
	waitlist_capacity,
	created_at,
	updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	res, err := d.DB.Exec(query, courseCode, courseTitle, semester, credits, section, openSeats, wa, wc, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *CourseDBImpl) UpdateCourse(id int, courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) error {
	query := `UPDATE tbl_Courses SET 
	course_code = $2,
	course_title = $3,
	semester = $4,
	credits = $5,
	section = $6,
	open_seats = $7,
	waitlist_available = $8,
	waitlist_capacity = $9,
	updated_at = $10
	WHERE id = $1`
	_, err := d.DB.Exec(query, id, courseCode, courseTitle, semester, credits, section, openSeats, wa, wc, time.Now())
	if err != nil {
		return err
	}

	return nil
}
