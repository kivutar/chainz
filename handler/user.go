package handler

import (
	"fmt"
	"net/http"

	gcontext "github.com/kivutar/chainz/context"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := gcontext.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", session.Values["profile"])
}
