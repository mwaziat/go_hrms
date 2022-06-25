package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/equimper/meetmeup/graph/database"
	"github.com/equimper/meetmeup/graph/generated"
	"github.com/equimper/meetmeup/graph/model"
)

func (r *authOpsResolver) Login(ctx context.Context, obj *model.AuthOps, email string, password string) (interface{}, error) {
	return db.UserLogin(email, password)
}

func (r *authOpsResolver) Register(ctx context.Context, obj *model.AuthOps, input model.NewUser) (interface{}, error) {
	return db.UserRegister(input), nil
}

func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.NewEmployee) (*model.Employee, error) {
	return db.SaveEmployee(input), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return db.Save(input), nil
}

func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOps, error) {
	return &model.AuthOps{}, nil
}

func (r *queryResolver) User(ctx context.Context, email string) (*model.User, error) {
	return db.FindByEmail(email), nil
}

func (r *queryResolver) UserUsername(ctx context.Context, username string) (*model.User, error) {
	return db.FindByUsername(username), nil
}

func (r *queryResolver) UserDelete(ctx context.Context, email string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.All(), nil
}

func (r *queryResolver) Employee(ctx context.Context, id string) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Employees(ctx context.Context) ([]*model.Employee, error) {
	return db.AllEmployee(), nil
}

func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	return "Success", nil
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()
