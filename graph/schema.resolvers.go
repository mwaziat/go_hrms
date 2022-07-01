package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mwaziat/gqlgen-hrms/graph/database"
	"github.com/mwaziat/gqlgen-hrms/graph/generated"
	"github.com/mwaziat/gqlgen-hrms/graph/model"
)

func (r *authOpsResolver) Login(ctx context.Context, obj *model.AuthOps, email string, password string) (interface{}, error) {
	return db.UserLogin(email, password)
}

func (r *authOpsResolver) Register(ctx context.Context, obj *model.AuthOps, input model.NewUser) (interface{}, error) {
	return db.UserRegister(input), nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return db.SaveTodo(input)
}

func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.NewEmployee) (*model.Employee, error) {
	return db.SaveEmployee(input), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return db.SaveUser(input), nil
}

func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOps, error) {
	return &model.AuthOps{}, nil
}

func (r *mutationResolver) UpdateEmployee(ctx context.Context, id string, input model.UpdateEmployee) (*model.Employee, error) {
	return db.UpdateEmployee(id, input)
}

func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (*bool, error) {
	return db.DeleteEmployee(id), nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return db.AllTodos(), nil
}

func (r *queryResolver) Todo(ctx context.Context, text *string) (*model.Todo, error) {
	return &model.Todo{
		ID:   "1234",
		Text: "Text test",
		Done: false,
		User: &model.User{
			ID:   "12",
			Name: "waziat",
		},
	}, nil
}

func (r *queryResolver) User(ctx context.Context, email string) (*model.User, error) {
	return db.FindUserByEmail(email), nil
}

func (r *queryResolver) UserUsername(ctx context.Context, username string) (*model.User, error) {
	return db.FindUserByUsername(username), nil
}

func (r *queryResolver) UserDelete(ctx context.Context, email string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.AllUsers(), nil
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
