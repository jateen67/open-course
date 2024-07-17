package db

import (
	"database/sql"
)

type NotificationDBImpl struct {
	DB *sql.DB
}

// CreateNotification implements NotificationDB.
func (n *NotificationDBImpl) CreateNotification(orderID int, notificationTypeID int) error {
	panic("unimplemented")
}

// GetNotification implements NotificationDB.
func (n *NotificationDBImpl) GetNotification() (*notification, error) {
	panic("unimplemented")
}

// GetNotifications implements NotificationDB.
func (n *NotificationDBImpl) GetNotifications() ([]*notification, error) {
	panic("unimplemented")
}

// GetNotificationsByNotificationTypeID implements NotificationDB.
func (n *NotificationDBImpl) GetNotificationsByNotificationTypeID(notificationTypeID int) ([]*notification, error) {
	panic("unimplemented")
}

// GetNotificationsByOrderID implements NotificationDB.
func (n *NotificationDBImpl) GetNotificationsByOrderID(orderID int) ([]*notification, error) {
	panic("unimplemented")
}

func NewNotificationDBImpl(db *sql.DB) *NotificationDBImpl {
	return &NotificationDBImpl{DB: db}
}
