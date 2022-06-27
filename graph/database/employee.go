package database

import (
	"context"
	"log"
	"time"

	"github.com/equimper/meetmeup/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) SaveEmployee(input model.NewEmployee) *model.Employee {
	collection := db.client.Database("hrms").Collection("employee")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}

	return &model.Employee{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Address:   input.Address,
		Position:  input.Position,
	}
}

func (db *DB) UpdateEmployee(id string, input model.UpdateEmployee) (*model.Employee, error) {
	collection := db.client.Database("hrms").Collection("employee")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	res := collection.FindOne(ctx, bson.M{"_id": ObjectId})
	if res != nil {
		filter := bson.M{"_id": ObjectId}
		update := bson.M{
			"firstName": input.FirstName,
			"lastName":  input.LastName,
			"email":     input.Email,
			"address":   input.Address,
			"position":  input.Position,
		}
		var updatedDocument bson.M
		err := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}).Decode(&updatedDocument)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, &gqlerror.Error{
					Message: "Email not found",
				}
			}
			log.Fatal(err)
		}

		return &model.Employee{
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
			Address:   input.Address,
			Position:  input.Position,
		}, nil
	}

	return nil, &gqlerror.Error{
		Message: "Email not found",
	}
}

func (db *DB) AllEmployee() []*model.Employee {
	collection := db.client.Database("hrms").Collection("employee")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var employees []*model.Employee

	for cur.Next(ctx) {
		var employee *model.Employee
		err := cur.Decode(&employee)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
	}

	return employees
}
