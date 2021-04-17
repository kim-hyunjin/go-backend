package todoapp

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kim-hyunjin/go-web/todoapp/model"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()
var dbHandler *model.DbHandler

type Success struct {
	Success bool `json:"success"`
}

type AppHandler struct {
	http.Handler // http.Handler를 embed
	db model.DbHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("name")
	todo := a.db.AddTodo(value)
	rd.JSON(w, http.StatusCreated, todo)
}

func (a *AppHandler) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.DeleteTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}	
}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) Close() {
	a.db.Close()
}

func MakeHandler(filepath string) *AppHandler {
	mux := mux.NewRouter()
	a := &AppHandler{
		Handler: mux,
		db: model.NewDbHandler(filepath),
	}
	mux.HandleFunc("/", a.indexHandler)
	mux.HandleFunc("/todos", a.getTodosHandler).Methods("GET")
	mux.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", a.deleteTodoHandler).Methods("DELETE")
	mux.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	mux.HandleFunc("/auth/google/login", googleLoginHandler)
	mux.HandleFunc("/auth/google/callback", googleOAuthCallbackHandler)

	return a
}