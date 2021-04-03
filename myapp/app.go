package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string	`json:"first_name"`
	LastName string		`json:"last_name"`
	Email string		`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	rw.Header().Add("content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, string(data))
}

func indexHandler (rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello world")
}

func barHandler (rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(rw, "hello, %s!", name)
}
func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	fmt.Println("Server Started!")
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	return mux
}