// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool   `json:"done"`
	User  *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
