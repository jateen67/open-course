package data

import (
	"context"
	"log"
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

func (l *LogNotification) Insert(orderIDs []int, notificationType string) error {
	collection := client.Database("notificationsdb").Collection("notifications")

	filter := bson.D{{Key: "type", Value: notificationType}}
	var result NotificationType
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("error getting notification type: ", err)
		return err
	}

	var logs []interface{}
	for _, i := range orderIDs {
		logs = append(logs, LogNotification{OrderID: i, NotificationTypeId: result.ID, TimeSent: time.Now()})
	}

	_, err = collection.InsertMany(context.TODO(), logs)

	if err != nil {
		log.Println("error logging notification: ", err)
		return err
	}

	return nil
}
