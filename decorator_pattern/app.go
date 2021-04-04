package decorator_pattern

import (
	"fmt"
	"net/http"
)


func indexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World")
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	h := NewDecoHandler(mux, logger)
	return h
}