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
            courseID INT PRIMARY KEY,
            termCode INT NOT NULL,
            session VARCHAR(5) NOT NULL,
            subject VARCHAR(4) NOT NULL,
            catalog VARCHAR(4) NOT NULL,
            section INT NOT NULL,
            componentCode VARCHAR(3) NOT NULL,
            componentDescription VARCHAR(20) NOT NULL,
            classNumber INT NOT NULL,
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
            orderId SERIAL PRIMARY KEY,
            email TEXT NOT NULL,
			phone TEXT NOT NULL,
			courseId INT REFERENCES tbl_Courses (courseId),
			isActive BOOLEAN NOT NULL,
			createdAt TIMESTAMPTZ DEFAULT NOW(),
			updatedAt TIMESTAMPTZ DEFAULT NOW()
        );
    `
	_, err := db.Exec(query)
	return err
}

func CourseExists(db *sql.DB, courseId int) (bool, error) {
	query := "SELECT COUNT(*) FROM tbl_Courses WHERE courseId = $1"
	var count int
	err := db.QueryRow(query, courseId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateDefaultCourse(db *sql.DB, courseID, termCode int, session, subject, catalog string, section int,
	componentCode, componentDescription string, classNumber, classAssociation int, courseTitle, classStartTime, classEndTime string,
	mondays, tuesdays, wednesdays, thursdays, fridays, saturdays, sundays bool, classStartDate, classEndDate string,
	enrollmentCapacity, currentEnrollment, waitlistCapacity, currentWaitlistTotal int) error {
	query := `INSERT INTO tbl_Courses (courseID, termCode, session, subject, catalog, section, componentCode, componentDescription,
			  classNumber, classAssociation, courseTitle, classStartTime, classEndTime, mondays, tuesdays, wednesdays, thursdays,
			  fridays, saturdays, sundays, classStartDate, classEndDate, enrollmentCapacity, currentEnrollment, waitlistCapacity,
			  currentWaitlistTotal) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, 
			  $18, $19, $20, $21, $22, $23, $24, $25, $26)`
	_, err := db.Exec(query, courseID, termCode, session, subject, catalog, section, componentCode, componentDescription,
		classNumber, classAssociation, courseTitle, classStartTime, classEndTime, mondays, tuesdays, wednesdays, thursdays,
		fridays, saturdays, sundays, classStartDate, classEndDate, enrollmentCapacity, currentEnrollment, waitlistCapacity,
		currentWaitlistTotal)
	return err
}

func OrderExists(db *sql.DB, email, phone string, courseID int) (bool, error) {
	query := "SELECT COUNT(*) FROM tbl_Orders WHERE (email = $1 OR phone = $2) AND courseId = $3"
	var count int
	err := db.QueryRow(query, email, phone, courseID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateDefaultOrder(db *sql.DB, email, phone string, courseID int) error {
	query := "INSERT INTO tbl_Orders (email, phone, courseId, isActive, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := db.Exec(query, email, phone, courseID, 1, time.Now(), time.Now())
	return err
}
