package db

import (
	"database/sql"
)

type dbImpl struct {
	DB *sql.DB
}

func NewDBImpl(db *sql.DB) *dbImpl {
	return &dbImpl{DB: db}
}

// CreateUser implements UserDB.
func (d *dbImpl) CreateUser(name string, email string, phone int) error {
	panic("unimplemented")
}

// GetUser implements UserDB.
func (d *dbImpl) GetUser() (*user, error) {
	panic("unimplemented")
}

// GetUserByEmail implements UserDB.
func (d *dbImpl) GetUserByEmail(email string) (*user, error) {
	panic("unimplemented")
}

// GetUserByName implements UserDB.
func (d *dbImpl) GetUserByName(name string) ([]*user, error) {
	panic("unimplemented")
}

// GetUserByPhone implements UserDB.
func (d *dbImpl) GetUserByPhone(phone int) (*user, error) {
	panic("unimplemented")
}

// GetUsers implements UserDB.
func (d *dbImpl) GetUsers() ([]*user, error) {
	panic("unimplemented")
}
