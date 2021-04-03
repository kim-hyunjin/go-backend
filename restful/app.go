package restful

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello, world")
}

func usersHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Get user info by /users/:id")
}

func getUserInfoHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(rw, "User Id : ", vars["id"])
}

// NewHandler make a new handler
func NewHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)
	return mux
}