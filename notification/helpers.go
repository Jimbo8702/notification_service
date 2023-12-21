package main

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(conStr string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conStr))
	if err != nil {
		return nil, err
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func NewFirebaseMessageClient(fbAccountID string) (*messaging.Client, error) {
	conf := &firebase.Config{
		ServiceAccountID: fbAccountID,
	}
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		return nil, err
	}
	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}
	return client, err
}