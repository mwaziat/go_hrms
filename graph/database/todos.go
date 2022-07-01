package database

import (
	"context"
	"log"
	"time"

	"github.com/mwaziat/gqlgen-hrms/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) SaveTodo(input model.NewTodo) (*model.Todo, error) {
	collection := db.client.Database("hrms").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}

	return &model.Todo{
		ID:   res.InsertedID.(primitive.ObjectID).Hex(),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "Waziat",
		},
	}, nil
}

func (db *DB) AllTodos() []*model.Todo {
	collection := db.client.Database("hrms").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var todos []*model.Todo

	for cur.Next(ctx) {
		var todo *model.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	return todos
}
