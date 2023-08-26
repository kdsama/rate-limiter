package mongo

import (
	// "context"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
	Db     string
}

func GetMongoClient(uri string, db string) *MongoClient {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	log.Print("Connected to MongoDB!")
	return &MongoClient{Client: client, Db: db}
}
