package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 处理异常
				// 记录日志
				log.Printf("Panic：%v", err)
				// 返回统一错误相应
				c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
					Detail:  fmt.Sprintf("%v", err),
				})
			}
		}()
		c.Next()
	}
}
