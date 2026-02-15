package session

import (
    "github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func InitStore(key string) {
    Store = sessions.NewCookieStore([]byte(key))
}
