package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

// GraphQl in Go - GQLGen Tutorial
//https://www.youtube.com/watch?v=O6jYy421tGw&t=783s

import (
	"context"
	"fmt"
	"github.com/AshleyCoder3/GolangTodo/graph/generated"
	"github.com/AshleyCoder3/GolangTodo/graph/model"
	"strconv"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// add to the list of Todos
	todo := &model.Todo{
		ID:   strconv.Itoa(len(r.TodosList) + 1),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: fmt.Sprintf("Ashley%s", strconv.Itoa(len(r.TodosList)+1)),
		},
	}

	r.TodosList = append(r.TodosList, todo)
	return todo, nil

}

// Todos * = pointer - memory address of the variable
//then & needed to show/go to location
//then dereference with *

//		this object	<== function of
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
