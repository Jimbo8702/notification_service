package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Jimbo8702/notification_service/types"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const notificationColl = "notification_tokens"
const hours_in_a_week = 24 * 7

type NotificationTokenStore interface {
	GetTokensByProfileID(ctx context.Context, id string) ([]*types.NotificationToken, error)
	DeviceTokenExists(ctx context.Context, token string) (bool, error)
}

type MongoNotificationStore struct {
	client *mongo.Client 
	coll   *mongo.Collection
}

func NewMongoNotificationStore(client *mongo.Client, dbName string) NotificationTokenStore {
	mns := &MongoNotificationStore{
		client: client,
		coll: client.Database(dbName).Collection(notificationColl),
	}
	if err := mns.initTimestampIndex(); err != nil {
		log.Fatal(err)
	}
	return mns
}

func (s *MongoNotificationStore) initTimestampIndex() error {
	index := mongo.IndexModel{
        Keys: bson.M{"timestamp": 1},
        Options: options.Index().SetExpireAfterSeconds(
            int32((time.Hour * 3 * hours_in_a_week).Seconds()),
        ),
    }
	_, err := s.coll.Indexes().DropAll(context.Background())
	if err != nil {
		logrus.Warnf("Drop index error while creating the mongo notification store: %s", err)
	}
	_, err = s.coll.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoNotificationStore) Insert(ctx context.Context, token *types.NotificationToken) (*types.NotificationToken, error) {
	res, err := s.coll.InsertOne(ctx, token)
	if err != nil {
		return nil, err
	}
	token.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return token, nil
}

func (s *MongoNotificationStore) DeviceTokenExists(ctx context.Context, deviceID string) (bool, error) {
	filter := bson.M{"device_id": deviceID}
	if err := s.coll.FindOne(ctx, filter).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *MongoNotificationStore) GetTokensByProfileID(ctx context.Context, id string) ([]*types.NotificationToken, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	resp, err := s.coll.Find(ctx, bson.M{"user_id": oid})
	if err != nil {
		return nil, err
	}
	var tokens []*types.NotificationToken
	if err := resp.All(ctx, &tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}
