package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

func addTestTodos() {
	todoMap[1] = &Todo{1, "Buy a milk", false, time.Now()}
	todoMap[2] = &Todo{1, "Exercise", true, time.Now()}
	todoMap[3] = &Todo{1, "Home work", true, time.Now()}
}

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	addTestTodos()
	rd = render.New()
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/todos", getTodosHandler).Methods("GET")
	return mux
}

func main() {
	n := negroni.Classic()
	n.UseHandler(MakeHandler())
	http.ListenAndServe(":3000", n)
}