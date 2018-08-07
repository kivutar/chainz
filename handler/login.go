package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"

	"golang.org/x/oauth2"

	gcontext "github.com/kivutar/chainz/context"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	domain := os.Getenv("AUTH0_DOMAIN")
	aud := os.Getenv("AUDIENCE")

	conf := &oauth2.Config{
		ClientID:     os.Getenv("AUTH0_ID"),
		ClientSecret: os.Getenv("AUTH0_SECRET"),
		RedirectURL:  os.Getenv("http://localhost:3000/callback"),
		Scopes:       []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	if aud == "" {
		aud = "https://" + domain + "/userinfo"
	}

	// Generate random state
	b := make([]byte, 32)
	rand.Read(b)
	state := base64.StdEncoding.EncodeToString(b)

	session, err := gcontext.Store.Get(r, "state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	audience := oauth2.SetAuthURLParam("audience", aud)
	url := conf.AuthCodeURL(state, audience)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
