package db

import (
	"database/sql"

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
		if err := rows.Scan(
			&course.ClassNumber, &course.CourseID, &course.TermCode, &course.Session, &course.Subject, &course.Catalog,
			&course.Section, &course.ComponentCode, &course.ComponentDescription, &course.ClassAssociation, &course.CourseTitle,
			&course.ClassStartTime, &course.ClassEndTime, &course.Mondays, &course.Tuesdays, &course.Wednesdays, &course.Thursdays,
			&course.Fridays, &course.Saturdays, &course.Sundays, &course.ClassStartDate, &course.ClassEndDate,
			&course.EnrollmentCapacity, &course.CurrentEnrollment, &course.WaitlistCapacity, &course.CurrentWaitlistTotal,
		); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func (d *CourseDBImpl) GetCoursesByInput(input string, termCode int) ([]Course, error) {
	query := "SELECT DISTINCT ON (courseId) * FROM tbl_Courses WHERE termCode = $1 AND LOWER(subject || ' ' || catalog || ' ' || coursetitle) LIKE '%' || LOWER($2) || '%' LIMIT 5"
	rows, err := d.DB.Query(query, termCode, input)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		if err := rows.Scan(
			&course.ClassNumber, &course.CourseID, &course.TermCode, &course.Session, &course.Subject, &course.Catalog,
			&course.Section, &course.ComponentCode, &course.ComponentDescription, &course.ClassAssociation, &course.CourseTitle,
			&course.ClassStartTime, &course.ClassEndTime, &course.Mondays, &course.Tuesdays, &course.Wednesdays, &course.Thursdays,
			&course.Fridays, &course.Saturdays, &course.Sundays, &course.ClassStartDate, &course.ClassEndDate,
			&course.EnrollmentCapacity, &course.CurrentEnrollment, &course.WaitlistCapacity, &course.CurrentWaitlistTotal,
		); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func (d *CourseDBImpl) GetCourseInfo(courseID, termCode int) ([]Course, error) {
	query := "SELECT * FROM tbl_Courses WHERE termCode = $1 AND courseId = $2"
	rows, err := d.DB.Query(query, termCode, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		if err := rows.Scan(
			&course.ClassNumber, &course.CourseID, &course.TermCode, &course.Session, &course.Subject, &course.Catalog,
			&course.Section, &course.ComponentCode, &course.ComponentDescription, &course.ClassAssociation, &course.CourseTitle,
			&course.ClassStartTime, &course.ClassEndTime, &course.Mondays, &course.Tuesdays, &course.Wednesdays, &course.Thursdays,
			&course.Fridays, &course.Saturdays, &course.Sundays, &course.ClassStartDate, &course.ClassEndDate,
			&course.EnrollmentCapacity, &course.CurrentEnrollment, &course.WaitlistCapacity, &course.CurrentWaitlistTotal,
		); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func (d *CourseDBImpl) GetCoursesByMultpleIDs(classNumbers []int) ([]Course, error) {
	query := "SELECT * FROM tbl_Courses WHERE classNumber = ANY($1)"
	rows, err := d.DB.Query(query, pq.Array(classNumbers))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		if err := rows.Scan(
			&course.ClassNumber, &course.CourseID, &course.TermCode, &course.Session, &course.Subject, &course.Catalog,
			&course.Section, &course.ComponentCode, &course.ComponentDescription, &course.ClassAssociation, &course.CourseTitle,
			&course.ClassStartTime, &course.ClassEndTime, &course.Mondays, &course.Tuesdays, &course.Wednesdays, &course.Thursdays,
			&course.Fridays, &course.Saturdays, &course.Sundays, &course.ClassStartDate, &course.ClassEndDate,
			&course.EnrollmentCapacity, &course.CurrentEnrollment, &course.WaitlistCapacity, &course.CurrentWaitlistTotal,
		); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func (d *CourseDBImpl) GetCoursesBySemester(termCode int) ([]Course, error) {
	query := "SELECT * FROM tbl_Courses WHERE termCode = $1"
	rows, err := d.DB.Query(query, termCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		if err := rows.Scan(
			&course.ClassNumber, &course.CourseID, &course.TermCode, &course.Session, &course.Subject, &course.Catalog,
			&course.Section, &course.ComponentCode, &course.ComponentDescription, &course.ClassAssociation, &course.CourseTitle,
			&course.ClassStartTime, &course.ClassEndTime, &course.Mondays, &course.Tuesdays, &course.Wednesdays, &course.Thursdays,
			&course.Fridays, &course.Saturdays, &course.Sundays, &course.ClassStartDate, &course.ClassEndDate,
			&course.EnrollmentCapacity, &course.CurrentEnrollment, &course.WaitlistCapacity, &course.CurrentWaitlistTotal,
		); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}
