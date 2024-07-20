package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	connString := os.Getenv("POSTGRES_CONNECTION_STRING")
	count := 1

	for {
		db, err := sql.Open("postgres", connString)
		if err != nil {
			log.Println("could not connect to postgres. retrying... ")
			count++
		} else {
			err = db.Ping()
			if err != nil {
				log.Println("postgres connection test failed. retrying...")
				count++
				db.Close()
			} else {
				return db, nil
			}
		}

		if count > 10 {
			return nil, err
		}

		log.Println("retrying in 1 second...")
		time.Sleep(1 * time.Second)
	}
}

func CreateTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS tbl_Courses (
            id SERIAL PRIMARY KEY,
            course_code VARCHAR(10) NOT NULL,
            course_title TEXT NOT NULL,
			semester VARCHAR(10) NOT NULL,
			section VARCHAR(10) NOT NULL,
			credits VARCHAR(5) NOT NULL,
			open_seats INT NOT NULL,
			waitlist_available INT NOT NULL,
			waitlist_capacity INT NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
        );

		CREATE TABLE IF NOT EXISTS tbl_Orders (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT NOT NULL,
			phone TEXT NOT NULL,
			course_id INT REFERENCES tbl_Courses (id),
			is_active BOOLEAN NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
        );

		CREATE TABLE IF NOT EXISTS tbl_Notification_Types (
			id SERIAL PRIMARY KEY,
			type VARCHAR(10) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tbl_Notifications (
            id SERIAL PRIMARY KEY,
            order_id INT REFERENCES tbl_Orders (id),
            notification_type_id INT REFERENCES tbl_Notification_Types (id),
			time_sent TIMESTAMPTZ DEFAULT NOW()
        );
    `
	_, err := db.Exec(query)
	return err
}

func CourseExists(db *sql.DB, courseCode, semester, section string) (bool, error) {
	query := "SELECT COUNT(*) FROM tbl_Courses WHERE course_code = $1 AND semester = $2 AND section = $3"
	var count int
	err := db.QueryRow(query, courseCode, semester, section).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateDefaultCourse(db *sql.DB, courseCode, courseTitle, semester, section, credits string, openSeats, wa, wc int) error {
	query := `INSERT INTO tbl_Courses (course_code, course_title, 
			  semester, section, credits, open_seats, waitlist_available, waitlist_capacity, created_at, updated_at)
	 		  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := db.Exec(query, courseCode, courseTitle, semester, section, credits, openSeats, wa, wc, time.Now(), time.Now())
	return err
}

func OrderExists(db *sql.DB, name, email, phone string, courseID int) (bool, error) {
	query := "SELECT COUNT(*) FROM tbl_Orders WHERE (email = $1 OR name = $2 OR phone = $3) AND course_id = $4"
	var count int
	err := db.QueryRow(query, email, name, phone, courseID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateDefaultOrder(db *sql.DB, name, email, phone string, courseID int) error {
	query := "INSERT INTO tbl_Orders (name, email, phone, course_id, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := db.Exec(query, name, email, phone, courseID, 1, time.Now(), time.Now())
	return err
}

func NotificationTypeExists(db *sql.DB, t string) (bool, error) {
	query := "SELECT COUNT(*) FROM tbl_Notification_Types WHERE type = $1"
	var count int
	err := db.QueryRow(query, t).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateDefaultNotificationType(db *sql.DB, t string) error {
	query := "INSERT INTO tbl_Notification_Types (type) VALUES ($1)"
	_, err := db.Exec(query, t)
	return err
}
