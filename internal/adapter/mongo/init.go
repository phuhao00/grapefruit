package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientInstance *Client

func InitMongo() {
	ctx := context.Background()
	c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}
	clientInstance = &Client{}
	clientInstance.client = c
}
