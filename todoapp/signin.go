package todoapp

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var googleOauthConfig oauth2.Config;

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := generateStateOauthCookie(w)
	url := googleOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{
		Name:"oauthstate", 
		Value: state, 
		Expires: time.Now().Add(1 * 24 * time.Hour),
	}
	http.SetCookie(w, &cookie)
	return state
}

func googleOAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	oauthstate, _ := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthstate.Value {
		log.Printf("invalid google oauth state cookie: %s, state:%s\n", oauthstate.Value, r.FormValue("state"))
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	userInfo, err := getGoogleUserInfo(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	fmt.Fprint(w, string(userInfo))
}

var oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func getGoogleUserInfo(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("Failed to exchange %s\n", err.Error())
	}
	res, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to get user info %s\n", err.Error())
	}

	return ioutil.ReadAll(res.Body)
}
