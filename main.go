package main

import (
	"fmt"
	"net/http"

	"github.com/kim-hyunjin/go-web/decorator_pattern"
)



func main() {
	fmt.Println("Server Starting...")
	mux := decorator_pattern.NewHandler()

	http.ListenAndServe(":3000", mux)
}