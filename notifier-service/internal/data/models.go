package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

type NotificationType struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type int                `bson:"type" json:"type"`
}

func (l *LogNotification) Insert(orderID int, notificationType string) error {
	collection := client.Database("notificationsdb").Collection("notifications")

	filter := bson.D{{Key: "type", Value: notificationType}}
	var result NotificationType
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err
	}

	var logs []interface{}
	logs = append(logs, LogNotification{OrderID: orderID, NotificationTypeId: result.ID, TimeSent: time.Now()})

	_, err = collection.InsertOne(context.TODO(), logs)

	if err != nil {
		return err
	}

	return nil
}
