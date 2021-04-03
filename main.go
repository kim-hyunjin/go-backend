package main

import (
	"fmt"
	"net/http"

	"github.com/kim-hyunjin/go-web/restful"
)

func main() {
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", restful.NewHandler())
}