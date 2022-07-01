package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/mwaziat/gqlgen-hrms/graph/model"
	"github.com/mwaziat/gqlgen-hrms/graph/services"
	"github.com/mwaziat/gqlgen-hrms/graph/tools"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) UserRegister(input model.NewUser) *model.User {
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

func (db *DB) UserLogin(email string, password string) (interface{}, error) {
	user := db.FindUserByEmail(email)
	if user == nil {
		return nil, &gqlerror.Error{
			Message: "Email not found",
		}
	}

	if !tools.ComparePassword(password, user.Password) {
		return nil, errors.New("Password not match")
	}

	expiredAt := time.Now().Add(time.Hour * 72).Unix()
	token, err := services.JwtGenerate(user.ID, expiredAt)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":     token,
		"ExpiredAt": int(expiredAt),
	}, nil
}
