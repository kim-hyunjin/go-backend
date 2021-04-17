package model

type sqliteHandler struct {

}

func (s *sqliteHandler) getTodos() []*Todo {
	return nil
}

func (s *sqliteHandler) addTodo(name string) *Todo {
	return nil
}

func (s *sqliteHandler) deleteTodo(id int) bool {
	return false
}

func (s *sqliteHandler) completeTodo(id int, complete bool) bool {
	return false
}