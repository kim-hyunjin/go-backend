package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	app := MakeHandler("./test.db")
	defer app.Close()
	n := negroni.Classic()
	n.UseHandler(app.Handler)
	fmt.Println("App Started!")
	http.ListenAndServe(":3000", n)
}