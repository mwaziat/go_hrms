package database

import (
	"context"
	"log"
	"time"

	"github.com/mwaziat/gqlgen-hrms/graph/model"
	"github.com/mwaziat/gqlgen-hrms/graph/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) SaveUser(input model.NewUser) *model.User {
	collection := db.client.Database("hrms").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userHashPassword, err := tools.HashPassword(input.Password)
	if err != nil {
		log.Fatal(err)
	}

	input.Password = userHashPassword
	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}

	return &model.User{
		ID:       res.InsertedID.(primitive.ObjectID).Hex(),
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: userHashPassword,
	}

}

func HashPassword(s string) {
	panic("unimplemented")
}

func (db *DB) FindUserByID(ID string) *model.User {
	ObjectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("hrms").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := collection.FindOne(ctx, bson.M{"_id": ObjectId})
	user := model.User{}
	res.Decode(&user)
	return &user
}

func (db *DB) FindUserByEmail(email string) *model.User {
	collection := db.client.Database("hrms").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := model.User{}
	res := collection.FindOne(ctx, bson.M{"email": email})
	res.Decode(&user)

	return &user
}

func (db *DB) FindUserByUsername(username string) *model.User {
	collection := db.client.Database("hrms").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := collection.FindOne(ctx, bson.M{"username": username})
	user := model.User{}
	res.Decode(&user)
	return &user
}

func (db *DB) AllUsers() []*model.User {
	collection := db.client.Database("hrms").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var users []*model.User

	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}
