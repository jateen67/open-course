package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port     = "80"
	mongoUrl = "mongodb://mongo:27017"
)

type NotificationType struct {
	ID   string `bson:"_id,omitempty" json:"id,omitempty"`
	Type string `bson:"type" json:"type"`
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	seed(mongoClient)

	log.Printf("starting mailer service on port %s\n", port)
	srv := newServer(mongoClient).Router

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), srv)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("mailer service closed")
	} else if err != nil {
		log.Println("error starting mailer service: ", err)
		os.Exit(1)
	}
}

func seed(client *mongo.Client) {
	var result bson.M
	opts := options.FindOne()
	ntExists := true

	collection := client.Database("notificationsdb").Collection("notificationTypes")
	err := collection.FindOne(context.TODO(), bson.D{{Key: "type", Value: "Email"}}, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ntExists = false
		}
	}

	if !ntExists {
		_, err = collection.InsertOne(context.TODO(), NotificationType{
			Type: "Email",
		})
		if err != nil {
			log.Fatalf("error seeding notificationTypes collection: %s", err)
		}
	}

	ntExists = true

	err = collection.FindOne(context.TODO(), bson.D{{Key: "type", Value: "SMS"}}, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ntExists = false
		}
	}

	if !ntExists {
		_, err = collection.InsertOne(context.TODO(), NotificationType{
			Type: "SMS",
		})
		if err != nil {
			log.Fatalf("error seeding notificationTypes collection: %s", err)
		}
	}

	log.Println("notificationType collection seeded successfully")
}

func connectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUrl)

	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting: ", err)
		return nil, err
	}

	log.Println("connected to mongo")

	return c, nil
}
