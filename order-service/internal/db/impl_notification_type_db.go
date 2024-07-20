package db

import (
	"database/sql"
)

type NotificationTypeDBImpl struct {
	DB *sql.DB
}

func NewNotificationTypeDBImpl(db *sql.DB) *NotificationTypeDBImpl {
	return &NotificationTypeDBImpl{DB: db}
}

func (d *NotificationTypeDBImpl) GetNotificationTypes() ([]notificationType, error) {
	query := "SELECT * FROM tbl_Notification_Types"
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notificationTypes []notificationType

	for rows.Next() {
		var notificationType notificationType
		if err := rows.Scan(&notificationType.ID, &notificationType.Type); err != nil {
			return notificationTypes, err
		}
		notificationTypes = append(notificationTypes, notificationType)
	}

	if err = rows.Err(); err != nil {
		return notificationTypes, err
	}

	return notificationTypes, nil
}

func (d *NotificationTypeDBImpl) CreateNotificationType(t string) (int, error) {
	var id int
	query := `INSERT INTO tbl_Notification_Types (
		type
		) VALUES ($1) RETURNING id`
	err := d.DB.QueryRow(query, t).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
