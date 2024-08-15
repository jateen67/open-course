package db

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
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
		if err := rows.Scan(&order.ID, &order.Email,
			&order.Phone, &order.ClassNumber, &order.IsActive, &order.CreatedAt,
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
	if err := d.DB.QueryRow(query, orderID).Scan(&order.ID,
		&order.Email, &order.Phone, &order.ClassNumber,
		&order.IsActive, &order.CreatedAt, &order.UpdatedAt); err != nil {
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
		if err := rows.Scan(&order.ID, &order.Email,
			&order.Phone, &order.ClassNumber, &order.IsActive, &order.CreatedAt,
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
		if err := rows.Scan(&order.ID, &order.Email,
			&order.Phone, &order.ClassNumber, &order.IsActive, &order.CreatedAt,
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
	query := "SELECT * FROM tbl_Orders WHERE isActive = TRUE"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Email,
			&order.Phone, &order.ClassNumber, &order.IsActive, &order.CreatedAt,
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

func (d *OrderDBImpl) GetOrdersByClassNumber(classNumber int) ([]Order, error) {
	query := "SELECT * FROM tbl_Orders WHERE classNumber = $1"
	rows, err := d.DB.Query(query, classNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Email,
			&order.Phone, &order.ClassNumber, &order.IsActive, &order.CreatedAt,
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
		email,
		phone,
		classNumber,
		isActive,
		createdAt,
		updatedAt
		) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := d.DB.QueryRow(query, order.Email, order.Phone, order.ClassNumber, 1, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *OrderDBImpl) UpdateOrder(order Order) error {
	query := `UPDATE tbl_Orders SET 
		phone = $1,
		classNumber = $2,
		isActive = $3,
		updatedAt = $4
		WHERE phone = $1 AND classNumber = $2`
	_, err := d.DB.Exec(query, order.Phone, order.ClassNumber, order.IsActive, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (d *OrderDBImpl) UpdateOrderStatus(orderIds []int) error {
	query := `UPDATE tbl_Orders SET 
		isActive = 0,
		updatedAt = $2
		WHERE id = ANY($1)`
	_, err := d.DB.Exec(query, pq.Array(orderIds), time.Now())
	if err != nil {
		return err
	}

	return nil
}
