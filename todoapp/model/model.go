package model

import "time"

type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type DbHandler interface {
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	DeleteTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

func NewDbHandler(filepath string) DbHandler {
	return newSqliteHandler(filepath)
}
