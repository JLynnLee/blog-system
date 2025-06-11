package route

import (
	"github.com/JLynnLee/go-blog/handler"
	"github.com/JLynnLee/go-blog/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(r *gin.Engine) {
	r.Use(middleware.ExceptionMiddleware())

	// web页面
	// 主页
	r.GET("/", middleware.AuthMiddleware(), handler.GetPosts)
	// 创建文章页
	r.GET("/create", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", nil)
	})
	// 登录页
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// 注册页
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	// api接口
	// 登录
	r.POST("/login", handler.Login)
	// 注册
	r.POST("/register", handler.Register)
	// 登出
	r.POST("/logout", middleware.AuthMiddleware(), handler.Logout)
	// 创建文章
	r.POST("/create", middleware.AuthMiddleware(), handler.CreatePost)
	// 测试
	r.GET("/test", func(c *gin.Context) {
		panic("something went wrong")
	})
}
