package db

import (
	"database/sql"
	"time"
)

type UserDBImpl struct {
	DB *sql.DB
}

func NewUserDBImpl(db *sql.DB) *UserDBImpl {
	return &UserDBImpl{DB: db}
}

func (d *UserDBImpl) GetUsers() ([]user, error) {
	query := "SELECT * FROM tbl_Users"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email,
			&user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (d *UserDBImpl) GetUser(userID int) (*user, error) {
	query := "SELECT * FROM tbl_Users where id = $1"
	var user user
	if err := d.DB.QueryRow(query, userID).Scan(&user); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (d *UserDBImpl) GetUsersByName(name string) ([]user, error) {
	query := "SELECT TOP(1) * FROM tbl_Users WHERE name = $1"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email,
			&user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (d *UserDBImpl) GetUsersByEmail(email string) ([]user, error) {
	query := "SELECT TOP(1) * FROM tbl_Users WHERE email = $1"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email,
			&user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (d *UserDBImpl) GetUsersByPhone(phone int) ([]user, error) {
	query := "SELECT TOP(1) * FROM tbl_Users WHERE phone = $1"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email,
			&user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (d *UserDBImpl) CreateUser(name, email string, phone int) (int64, error) {
	query := "INSERT INTO tbl_Users (name, email, phone, created_at) VALUES ($1, $2, $3, $4)"
	res, err := d.DB.Exec(query, name, email, phone, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
