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

func (d *NotificationDBImpl) GetNotifications() ([]Notification, error) {
	query := "SELECT * FROM tbl_Notifications"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []Notification

	for rows.Next() {
		var notification Notification
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

func (d *NotificationDBImpl) GetNotification(notificationID int) (*Notification, error) {
	query := "SELECT * FROM tbl_Notifications WHERE id = $1"
	var notification Notification
	if err := d.DB.QueryRow(query, notificationID).Scan(&notification); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &notification, nil
}

func (d *NotificationDBImpl) GetNotificationsByNotificationTypeID(notificationTypeID int) ([]Notification, error) {
	query := "SELECT * FROM tbl_Notifications WHERE notificationTypeId = $1"
	rows, err := d.DB.Query(query, notificationTypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []Notification

	for rows.Next() {
		var notification Notification
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

func (d *NotificationDBImpl) GetNotificationsByOrderID(orderID int) ([]Notification, error) {
	query := "SELECT * FROM tbl_Notifications WHERE orderId = $1"
	rows, err := d.DB.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []Notification

	for rows.Next() {
		var notification Notification
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

func (d *NotificationDBImpl) CreateNotification(notification Notification) (int, error) {
	var id int
	query := `INSERT INTO tbl_Notifications (
		orderId,
		notificationTypeId,
		timeSent
		) VALUES ($1, $2, $3) RETURNING id`
	err := d.DB.QueryRow(query, notification.OrderID, notification.NotificationTypeID, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
