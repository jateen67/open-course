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
		DO $$ DECLARE
			r RECORD;
		BEGIN
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
				EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
			END LOOP;
		END $$;

		CREATE TABLE IF NOT EXISTS tbl_Courses (
            classNumber INT PRIMARY KEY,
            courseId INT NOT NULL,
            termCode INT NOT NULL,
            session VARCHAR(5) NOT NULL,
            subject VARCHAR(4) NOT NULL,
            catalog VARCHAR(4) NOT NULL,
            section VARCHAR(10) NOT NULL,
            componentCode VARCHAR(3) NOT NULL,
            componentDescription VARCHAR(40) NOT NULL,
            classAssociation INT NOT NULL,
            courseTitle TEXT NOT NULL,
            classStartTime VARCHAR(10) NOT NULL,
            classEndTime VARCHAR(10) NOT NULL,
			mondays BOOLEAN NOT NULL,
			tuesdays BOOLEAN NOT NULL,
			wednesdays BOOLEAN NOT NULL,
			thursdays BOOLEAN NOT NULL,
			fridays BOOLEAN NOT NULL,
			saturdays BOOLEAN NOT NULL,
			sundays BOOLEAN NOT NULL,
            classStartDate VARCHAR(10) NOT NULL,
            classEndDate VARCHAR(10) NOT NULL,
			enrollmentCapacity INT NOT NULL,
			currentEnrollment INT NOT NULL,
			waitlistCapacity INT NOT NULL,
			currentWaitlistTotal INT NOT NULL
        );

		CREATE TABLE IF NOT EXISTS tbl_Orders (
            id SERIAL PRIMARY KEY,
            email TEXT NOT NULL,
			phone TEXT NOT NULL,
			classNumber INT REFERENCES tbl_Courses (classNumber),
			isActive BOOLEAN NOT NULL,
			createdAt TIMESTAMPTZ DEFAULT NOW(),
			updatedAt TIMESTAMPTZ DEFAULT NOW()
        );
    `
	_, err := db.Exec(query)
	return err
}

func CreateIndexes(db *sql.DB) error {
	query := `
		CREATE INDEX idx_course_id ON tbl_Courses(courseId);
		CREATE INDEX idx_class_number ON tbl_Orders(classNumber);
		CREATE INDEX idx_phone ON tbl_Orders(phone);
	`
	_, err := db.Exec(query)
	return err
}

func CoursesTablePopulated(db *sql.DB) (bool, error) {
	query := "SELECT COUNT(1) WHERE EXISTS (SELECT * FROM tbl_Courses)"
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateDefaultCourse(db *sql.DB, classNumber, courseID, termCode int, session, subject, catalog, section string,
	componentCode, componentDescription string, classAssociation int, courseTitle, classStartTime, classEndTime string,
	mondays, tuesdays, wednesdays, thursdays, fridays, saturdays, sundays bool, classStartDate, classEndDate string,
	enrollmentCapacity, currentEnrollment, waitlistCapacity, currentWaitlistTotal int) error {
	query := `INSERT INTO tbl_Courses (classNumber, courseId, termCode, session, subject, catalog, section, componentCode, componentDescription,
			  classAssociation, courseTitle, classStartTime, classEndTime, mondays, tuesdays, wednesdays, thursdays,
			  fridays, saturdays, sundays, classStartDate, classEndDate, enrollmentCapacity, currentEnrollment, waitlistCapacity,
			  currentWaitlistTotal) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, 
			  $18, $19, $20, $21, $22, $23, $24, $25, $26)`
	_, err := db.Exec(query, classNumber, courseID, termCode, session, subject, catalog, section, componentCode, componentDescription,
		classAssociation, courseTitle, classStartTime, classEndTime, mondays, tuesdays, wednesdays, thursdays, fridays, saturdays,
		sundays, classStartDate, classEndDate, enrollmentCapacity, currentEnrollment, waitlistCapacity, currentWaitlistTotal)
	return err
}

func OrdersTablePopulated(db *sql.DB) (bool, error) {
	query := "SELECT COUNT(1) WHERE EXISTS (SELECT * FROM tbl_Orders)"
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateDefaultOrder(db *sql.DB, email, phone string, classNumber int) error {
	query := "INSERT INTO tbl_Orders (email, phone, classNumber, isActive, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := db.Exec(query, email, phone, classNumber, 1, time.Now(), time.Now())
	return err
}

func ContainsClassNumber(db *sql.DB, classNumber int) (bool, error) {
	query := "SELECT COUNT(1) WHERE EXISTS (SELECT * FROM tbl_Courses WHERE classNumber = $1)"
	var count int
	err := db.QueryRow(query, classNumber).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
