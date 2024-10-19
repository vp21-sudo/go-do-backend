package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient    *mongo.Client
	TodoCollection *mongo.Collection
)

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	TodoCollection = client.Database("godo_db").Collection("todos")
}

func Disconnect(ctx context.Context) {
	if err := MongoClient.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}
