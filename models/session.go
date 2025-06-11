package models

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("secret-key"))
