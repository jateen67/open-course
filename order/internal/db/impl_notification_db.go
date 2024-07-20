package db

import (
	"database/sql"
	"time"
)

type NotificationDBImpl struct {
	DB *sql.DB
}

func NewNotificationDBImpl(db *sql.DB) *NotificationDBImpl {
	return &NotificationDBImpl{DB: db}
}

func (d *NotificationDBImpl) GetNotifications() ([]notification, error) {
	query := "SELECT * FROM tbl_Notifications"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []notification

	for rows.Next() {
		var notification notification
		if err := rows.Scan(&notification.ID, &notification.OrderID, &notification.NotificationTypeID,
			&notification.TimeSent); err != nil {
			return notifications, err
		}
		notifications = append(notifications, notification)
	}

	if err = rows.Err(); err != nil {
		return notifications, err
	}

	return notifications, nil
}

func (d *NotificationDBImpl) GetNotification(notificationID int) (*notification, error) {
	query := "SELECT * FROM tbl_Notifications WHERE id = $1"
	var notification notification
	if err := d.DB.QueryRow(query, notificationID).Scan(&notification); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &notification, nil
}

func (d *NotificationDBImpl) GetNotificationsByNotificationTypeID(notificationTypeID int) ([]notification, error) {
	query := "SELECT * FROM tbl_Notifications WHERE notification_type_id = $1"
	rows, err := d.DB.Query(query, notificationTypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []notification

	for rows.Next() {
		var notification notification
		if err := rows.Scan(&notification.ID, &notification.OrderID, &notification.NotificationTypeID,
			&notification.TimeSent); err != nil {
			return notifications, err
		}
		notifications = append(notifications, notification)
	}

	if err = rows.Err(); err != nil {
		return notifications, err
	}

	return notifications, nil
}

func (d *NotificationDBImpl) GetNotificationsByOrderID(orderID int) ([]notification, error) {
	query := "SELECT * FROM tbl_Notifications WHERE order_id = $1"
	rows, err := d.DB.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []notification

	for rows.Next() {
		var notification notification
		if err := rows.Scan(&notification.ID, &notification.OrderID, &notification.NotificationTypeID,
			&notification.TimeSent); err != nil {
			return notifications, err
		}
		notifications = append(notifications, notification)
	}

	if err = rows.Err(); err != nil {
		return notifications, err
	}

	return notifications, nil
}

func (d *NotificationDBImpl) CreateNotification(orderID int, notificationTypeID int) (int64, error) {
	query := `INSERT INTO tbl_Notifications (
		order_id,
		notification_type_id,
		time_sent
		) VALUES ($1, $2, $3)`
	res, err := d.DB.Exec(query, orderID, notificationTypeID, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}