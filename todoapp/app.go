package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}
func MakeHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
	return mux
}

func main() {
	n := negroni.Classic()
	n.UseHandler(MakeHandler())
	http.ListenAndServe(":3000", n)
}