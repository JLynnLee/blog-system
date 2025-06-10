package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "session-name")
		if err != nil {
			//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "无法获取 session"})
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		userID, ok := session.Values["user_id"]
		if !ok {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		var userIDUint uint
		switch v := userID.(type) {
		case int:
			userIDUint = uint(v)
		case int32:
			userIDUint = uint(v)
		case uint:
			userIDUint = v
		case float64:
			userIDUint = uint(v)
		default:
			c.Redirect(http.StatusFound, strconv.Itoa(int(userIDUint)))
			c.Abort()
			return
		}

		c.Set("user_id", userIDUint)
		c.Next()
	}
}
