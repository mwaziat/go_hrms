package database

import (
	"context"
	"log"
	"time"

	"github.com/equimper/meetmeup/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
