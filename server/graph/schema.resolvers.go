package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AshleyCoder3/GolangTodo/graph/generated"
	"github.com/AshleyCoder3/GolangTodo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// add to the list of Todos
	todo := &model.Todo{
		ID:    strconv.Itoa(len(r.TodosList) + 1),
		Title: input.Title,
		Body:  input.Body,
		Done:  false,
		User: &model.User{
			ID:   input.UserID,
			Name: fmt.Sprintf("Ashley%s", strconv.Itoa(len(r.TodosList)+1)),
		},
	}

	r.TodosList = append(r.TodosList, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// return list of todos
	/* return []*model.Todo{
		{
			ID:   "123",
			Text: "testing",
			Done: false,
			User: &model.User{
				ID:   "user1",
				Name: "Ashley",
			},
		},
	}, nil
	*/
	return r.TodosList, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	return r.UserList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type todoResolver struct{ *Resolver }
