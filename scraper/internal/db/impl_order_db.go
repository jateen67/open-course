package db

import (
	"database/sql"
)

type OrderDBImpl struct {
	DB *sql.DB
}

// CreateOrder implements OrderDB.
func (o *OrderDBImpl) CreateOrder(userID int, courseID int) error {
	panic("unimplemented")
}

// GetActiveOrders implements OrderDB.
func (o *OrderDBImpl) GetActiveOrders() ([]*order, error) {
	panic("unimplemented")
}

// GetOrder implements OrderDB.
func (o *OrderDBImpl) GetOrder() (*order, error) {
	panic("unimplemented")
}

// GetOrders implements OrderDB.
func (o *OrderDBImpl) GetOrders() ([]*order, error) {
	panic("unimplemented")
}

// GetOrdersByCourseID implements OrderDB.
func (o *OrderDBImpl) GetOrdersByCourseID(courseID int) ([]*order, error) {
	panic("unimplemented")
}

// GetOrdersByUserID implements OrderDB.
func (o *OrderDBImpl) GetOrdersByUserID(userID int) (*order, error) {
	panic("unimplemented")
}

// UpdateOrder implements OrderDB.
func (o *OrderDBImpl) UpdateOrder(id int, userID int, courseID int) error {
	panic("unimplemented")
}

func NewOrderDBImpl(db *sql.DB) *OrderDBImpl {
	return &OrderDBImpl{DB: db}
}
