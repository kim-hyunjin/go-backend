package main

import (
	"net/http"

	"github.com/kim-hyunjin/go-web/restful"
)

func main() {
	http.ListenAndServe(":3000", restful.NewHandler())
}