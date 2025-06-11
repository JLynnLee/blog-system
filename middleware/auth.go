package middleware

import (
	"github.com/JLynnLee/go-blog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := model.GetSession(c)
		if err != nil {
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
