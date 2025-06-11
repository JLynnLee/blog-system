package model

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func GetSession(c *gin.Context) (*sessions.Session, error) {
	Store := sessions.NewCookieStore([]byte(c.MustGet("session_secret_key").(string)))
	return Store.Get(c.Request, "session-name")
}
