package main

import (
	"fmt"
	"net/http"

	todoapp "github.com/kim-hyunjin/go-web/todoapp"
	"github.com/urfave/negroni"
)

func main() {
	app := todoapp.MakeHandler("./test.db")
	defer app.Close()
	n := negroni.Classic()
	n.UseHandler(app)
	fmt.Println("App Started!")
	http.ListenAndServe(":3000", n)
}