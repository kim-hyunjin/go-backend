package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello Foo!")
}

func main() {
	fmt.Println("Server Started!")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "hello world")
	})

	http.HandleFunc("/bar", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "heelo Bar")
	})

	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil)
}