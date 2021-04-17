package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	todoapp "github.com/kim-hyunjin/go-web/todoapp"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	todoapp.GoogleOauthConfig = oauth2.Config{
		RedirectURL: "http://localhost:3000/auth/google/callback",
		ClientID: os.Getenv("client_id"),
		ClientSecret: os.Getenv("client_secret"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	todoapp.Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	app := todoapp.MakeHandler("./test.db")
	defer app.Close()
	
	fmt.Println("App Started!")
	http.ListenAndServe(":3000", app)
}