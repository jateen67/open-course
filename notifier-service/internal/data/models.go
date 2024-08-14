package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		LogNotification: LogNotification{},
	}
}

type Models struct {
	LogNotification LogNotification
}

type LogNotification struct {
	ID                 string             `bson:"_id,omitempty" json:"id,omitempty"`
	OrderID            int                `bson:"orderId" json:"orderId"`
	NotificationTypeId primitive.ObjectID `bson:"notificationTypeId" json:"notificationTypeId"`
	TimeSent           time.Time          `bson:"timeSent" json:"timeSent"`
}

func (l *LogNotification) Insert(n LogNotification) error {
	collection := client.Database("notificationsdb").Collection("notifications")

	_, err := collection.InsertOne(context.TODO(), LogNotification{
		OrderID:            n.OrderID,
		NotificationTypeId: n.NotificationTypeId,
		TimeSent:           time.Now(),
	})

	if err != nil {
		log.Println("error logging notification: ", err)
		return err
	}

	return nil
}
