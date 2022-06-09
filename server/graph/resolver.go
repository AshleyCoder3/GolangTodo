package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/AshleyCoder3/GolangTodo/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// best practice is to store state in this file

type Resolver struct {
	TodosList []*model.Todo
	UserList  []*model.User
}
