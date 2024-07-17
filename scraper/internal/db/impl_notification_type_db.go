package db

import (
	"database/sql"
)

type NotificationTypeDBImpl struct {
	DB *sql.DB
}

// CreateNotificationType implements NotificationTypeDB.
func (n *NotificationTypeDBImpl) CreateNotificationType(t string) error {
	panic("unimplemented")
}

// GetNotificationTypes implements NotificationTypeDB.
func (n *NotificationTypeDBImpl) GetNotificationTypes() ([]*notificationType, error) {
	panic("unimplemented")
}

func NewNotificationTypeDBImpl(db *sql.DB) *NotificationTypeDBImpl {
	return &NotificationTypeDBImpl{DB: db}
}
