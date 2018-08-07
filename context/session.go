package context

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

var (
	Store *sessions.CookieStore
)

func InitSessionStore() {
	Store = sessions.NewCookieStore([]byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
}
