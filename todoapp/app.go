package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kim-hyunjin/go-web/todoapp/model"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

type Success struct {
	Success bool `json:"success"`
}

var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	list := model.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("name")
	todo := model.AddTodo(value)
	rd.JSON(w, http.StatusCreated, todo)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := model.DeleteTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}	
}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := model.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func MakeHandler() http.Handler {
	rd = render.New()
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/todos", getTodosHandler).Methods("GET")
	mux.HandleFunc("/todos", addTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", deleteTodoHandler).Methods("DELETE")
	mux.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")

	return mux
}

func main() {
	n := negroni.Classic()
	n.UseHandler(MakeHandler())
	http.ListenAndServe(":3000", n)
}