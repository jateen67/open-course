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

func (d *OrderDBImpl) GetOrders() ([]order, error) {
	query := "SELECT * FROM tbl_Orders"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order

	for rows.Next() {
		var order order
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

func (d *OrderDBImpl) GetOrder(orderID int) (*order, error) {
	query := "SELECT * FROM tbl_Orders where id = $1"
	var order order
	if err := d.DB.QueryRow(query, orderID).Scan(&order); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &order, nil
}

func (d *OrderDBImpl) GetOrdersByUserEmail(email string) ([]order, error) {
	query := "SELECT * FROM tbl_Orders WHERE email = $1"
	rows, err := d.DB.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order

	for rows.Next() {
		var order order
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

func (d *OrderDBImpl) GetOrdersByUserName(name string) ([]order, error) {
	query := "SELECT * FROM tbl_Orders WHERE name = $1"
	rows, err := d.DB.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order

	for rows.Next() {
		var order order
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

func (d *OrderDBImpl) GetOrdersByUserPhone(phone string) ([]order, error) {
	query := "SELECT * FROM tbl_Orders WHERE phone = $1"
	rows, err := d.DB.Query(query, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order

	for rows.Next() {
		var order order
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

func (d *OrderDBImpl) GetActiveOrders() ([]order, error) {
	query := "SELECT * FROM tbl_Orders WHERE is_active = 1"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order

	for rows.Next() {
		var order order
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

func (d *OrderDBImpl) GetOrdersByCourseID(courseID int) ([]order, error) {
	query := "SELECT * FROM tbl_Orders WHERE course_id = $1"
	rows, err := d.DB.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order

	for rows.Next() {
		var order order
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

func (d *OrderDBImpl) CreateOrder(name, email, phone string, courseID int) (int64, error) {
	query := `INSERT INTO tbl_Orders (
		name,
		email,
		phone,
		course_id,
		is_active,
		created_at,
		updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	res, err := d.DB.Exec(query, name, email, phone, courseID, 1, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *OrderDBImpl) UpdateOrder(name, email, phone string, courseID int, isActive bool) error {
	query := `UPDATE tbl_Orders SET 
		name = $1,
		email = $2,
		phone = $3,
		course_id = $4,
		is_active = $5,
		updated_at = $6
		WHERE name = $1 AND email = $2 AND phone = $3 AND course_id = $4`
	_, err := d.DB.Exec(query, name, email, phone, courseID, isActive, time.Now())
	if err != nil {
		return err
	}

	return nil
}
