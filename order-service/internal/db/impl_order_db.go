package db

import (
	"database/sql"
	"time"
)

type OrderDBImpl struct {
	DB *sql.DB
}

func NewOrderDBImpl(db *sql.DB) *OrderDBImpl {
	return &OrderDBImpl{DB: db}
}

func (d *OrderDBImpl) GetOrders() ([]Order, error) {
	query := "SELECT * FROM tbl_Orders"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Name, &order.Email,
			&order.Phone, &order.CourseID, &order.IsActive, &order.CreatedAt,
			&order.UpdatedAt); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func (d *OrderDBImpl) GetOrder(orderID int) (*Order, error) {
	query := "SELECT * FROM tbl_Orders where id = $1"
	var order Order
	if err := d.DB.QueryRow(query, orderID).Scan(&order); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &order, nil
}

func (d *OrderDBImpl) GetOrdersByUserEmail(email string) ([]Order, error) {
	query := "SELECT * FROM tbl_Orders WHERE email = $1"
	rows, err := d.DB.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Name, &order.Email,
			&order.Phone, &order.CourseID, &order.IsActive, &order.CreatedAt,
			&order.UpdatedAt); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func (d *OrderDBImpl) GetOrdersByUserName(name string) ([]Order, error) {
	query := "SELECT * FROM tbl_Orders WHERE name = $1"
	rows, err := d.DB.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Name, &order.Email,
			&order.Phone, &order.CourseID, &order.IsActive, &order.CreatedAt,
			&order.UpdatedAt); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func (d *OrderDBImpl) GetOrdersByUserPhone(phone string) ([]Order, error) {
	query := "SELECT * FROM tbl_Orders WHERE phone = $1"
	rows, err := d.DB.Query(query, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Name, &order.Email,
			&order.Phone, &order.CourseID, &order.IsActive, &order.CreatedAt,
			&order.UpdatedAt); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func (d *OrderDBImpl) GetActiveOrders() ([]Order, error) {
	query := "SELECT * FROM tbl_Orders WHERE isActive = 1"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Name, &order.Email,
			&order.Phone, &order.CourseID, &order.IsActive, &order.CreatedAt,
			&order.UpdatedAt); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func (d *OrderDBImpl) GetOrdersByCourseID(courseID int) ([]Order, error) {
	query := "SELECT * FROM tbl_Orders WHERE courseId = $1"
	rows, err := d.DB.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Name, &order.Email,
			&order.Phone, &order.CourseID, &order.IsActive, &order.CreatedAt,
			&order.UpdatedAt); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func (d *OrderDBImpl) CreateOrder(order Order) (int, error) {
	var id int
	query := `INSERT INTO tbl_Orders (
		name,
		email,
		phone,
		courseId,
		isActive,
		createdAt,
		updatedAt
		) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := d.DB.QueryRow(query, order.Name, order.Email, order.Phone, order.CourseID, 1, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *OrderDBImpl) UpdateOrder(order Order) error {
	query := `UPDATE tbl_Orders SET 
		phone = $1,
		courseId = $2,
		isActive = $3,
		updatedAt = $4
		WHERE phone = $1 AND courseId = $2`
	_, err := d.DB.Exec(query, order.Phone, order.CourseID, order.IsActive, time.Now())
	if err != nil {
		return err
	}

	return nil
}
