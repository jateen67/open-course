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
        CREATE TABLE IF NOT EXISTS tbl_Users (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
			phone INT UNIQUE NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW()
        )

		CREATE TABLE IF NOT EXISTS tbl_Courses (
            id SERIAL PRIMARY KEY,
            course_code VARCHAR(10) NOT NULL,
            course_title NVARCHAR(200) NOT NULL,
			semester VARCHAR(10) NOT NULL,
			section VARCHAR(10) NOT NULL,
			credits VARCHAR(5) NOT NULL,
			open_seats INT NOT NULL,
			waitlist_available INT NOT NULL,
			waitlist_capacity INT NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW()
			updated_at TIMESTAMPTZ DEFAULT NOW()
        )

		CREATE TABLE IF NOT EXISTS tbl_Orders (
            id SERIAL PRIMARY KEY,
            user_id INT REFERENCES tbl_Users (id),
            course_id INT REFERENCES tbl_Courses (id),
			is_active BIT NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW()
			updated_at TIMESTAMPTZ DEFAULT NOW()
        )

		CREATE TABLE IF NOT EXISTS tbl_Notifications (
            id SERIAL PRIMARY KEY,
            order_id INT REFERENCES tbl_Orders (id),
            notification_type_id INT REFERENCES tbl_Notification_Types (id),
			time_sent TIMESTAMPTZ DEFAULT NOW()
        )

		CREATE TABLE IF NOT EXISTS tbl_Notification_Types (
            id SERIAL PRIMARY KEY,
            type VARCHAR(10) NOT NULL,
        )
    `
	_, err := db.Exec(query)
	return err
}

func GetUsers(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Users"
	res, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUser(db *sql.DB, userID int) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Users where id = $1"
	res, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserByName(db *sql.DB, name string) (*sql.Rows, error) {
	query := "SELECT TOP(1) * FROM tbl_Users WHERE name = $1"
	res, err := db.Query(query, name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserByEmail(db *sql.DB, email string) (*sql.Rows, error) {
	query := "SELECT TOP(1) * FROM tbl_Users WHERE email = $1"
	res, err := db.Query(query, email)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UserExists(db *sql.DB, email string) (bool, error) {
	query := "SELECT COUNT(*) FROM tbl_Users WHERE email= $1"
	var count int
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetUserByPhone(db *sql.DB, phone string) (*sql.Rows, error) {
	query := "SELECT TOP(1) * FROM tbl_Users WHERE phone = $1"
	res, err := db.Query(query, phone)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateUser(db *sql.DB, name, email string, phone int) error {
	query := "INSERT INTO tbl_Users (name, email, phone, created_at) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query, name, email, phone, time.Now())
	return err
}

func GetCourses(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Courses"
	res, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetCourse(db *sql.DB, courseID int) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Courses where id = $1"
	res, err := db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetCourseByCourseCode(db *sql.DB, courseCode string) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Users WHERE course_code = $1"
	res, err := db.Query(query, courseCode)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetCoursesBySemester(db *sql.DB, semester string) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Users WHERE semester = $1"
	res, err := db.Query(query, semester)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetCoursesBySection(db *sql.DB, section string) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Users WHERE section = $1"
	res, err := db.Query(query, section)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetOpenCourses(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Users WHERE open_seats > 0"
	res, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateCourse(db *sql.DB, courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) error {
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
	_, err := db.Exec(query, courseCode, courseTitle, semester, credits, section, openSeats, wa, wc, time.Now(), time.Now())
	return err
}

func UpdateCourse(db *sql.DB, id, courseCode, courseTitle, semester, credits, section string, openSeats, wa, wc int) error {
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
	_, err := db.Exec(query, id, courseCode, courseTitle, semester, credits, section, openSeats, wa, wc, time.Now())
	return err
}

func GetOrders(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Orders"
	res, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetOrder(db *sql.DB, orderID int) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Orders where id = $1"
	res, err := db.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetOrdersByUserID(db *sql.DB, userID int) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Orders WHERE user_id = $1"
	res, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetOrdersByCourseID(db *sql.DB, courseID int) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Orders WHERE course_id = $1"
	res, err := db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetActiveOrders(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Orders WHERE is_active = 1"
	res, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateOrder(db *sql.DB, userID, courseID int) error {
	query := `INSERT INTO tbl_Orders (
	user_id,
	course_id,
	is_active,
	created_at,
	updated_at
	) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, userID, courseID, 1, time.Now(), time.Now())
	return err
}

func UpdateOrder(db *sql.DB, id, isActive int) error {
	query := `UPDATE tbl_Orders SET 
		is_active = $2,
		updated_at = $3
		WHERE id = $1`
	_, err := db.Exec(query, id, isActive, time.Now())
	return err
}

func GetNotifications(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Notifications"
	res, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetNotification(db *sql.DB, notificationID int) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Notifications where id = $1"
	res, err := db.Query(query, notificationID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetNotificationsByOrderID(db *sql.DB, orderID int) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Notifications WHERE order_id = $1"
	res, err := db.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetNotificationsByNotificationTypeID(db *sql.DB, notificationTypeID int) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Notifications WHERE Notification_type_id = $1"
	res, err := db.Query(query, notificationTypeID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateNotification(db *sql.DB, orderID, notificationTypeID int) error {
	query := `INSERT INTO tbl_Notifications (
	order_id,
	notification_type_id,
	time_sent
	) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, orderID, notificationTypeID, time.Now())
	return err
}

func GetNotificationTypes(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT * FROM tbl_Notification_Types"
	res, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateNotificationType(db *sql.DB, t string) error {
	query := `INSERT INTO tbl_Notification_Types (
	type
	) VALUES ($1)`
	_, err := db.Exec(query, t)
	return err
}
