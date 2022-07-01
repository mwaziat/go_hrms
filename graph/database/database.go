package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:wanasaba10@wanasaba.h4rbt.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	return &DB{
		client: client,
	}

}
