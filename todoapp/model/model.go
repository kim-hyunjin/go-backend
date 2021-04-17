package model

import "time"

var handler dbHandler

func init() {
	handler = newMemoryHandler()
	// handler = new SqliteHandler()
}

type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type dbHandler interface {
	getTodos() []*Todo
	addTodo(name string) *Todo
	deleteTodo(id int) bool
	completeTodo(id int, complete bool) bool
}

func GetTodos() []*Todo {
	return handler.getTodos()
}

func AddTodo(name string) *Todo {
	return handler.addTodo(name);
}

func DeleteTodo(id int) bool {
	return handler.deleteTodo(id)
}

func CompleteTodo(id int, complete bool) bool {
	return handler.completeTodo(id, complete);
}