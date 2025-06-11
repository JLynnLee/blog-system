package main

import (
	"github.com/JLynnLee/go-blog/handlers"
	"github.com/JLynnLee/go-blog/middleware"
	"github.com/JLynnLee/go-blog/models"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	// 初始化 Gin
	r := gin.Default()

	// 使用绝对路径确保 SQLite 正确创建/读取数据库文件
	db, err := gorm.Open(sqlite.Open("D:/workspace/go/blog-system/blog.db"), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库: " + err.Error())
	}

	// 自动迁移模型
	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}

	// 注册中间件
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// 设置模板引擎
	r.LoadHTMLGlob("templates/*")

	// 路由
	r.GET("/", middleware.AuthMiddleware(), handlers.GetPosts)
	r.GET("/create", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", nil)
	})
	r.POST("/create", middleware.AuthMiddleware(), handlers.CreatePost)

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", handlers.Login)

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	r.POST("/register", handlers.Register)
	r.POST("/logout", middleware.AuthMiddleware(), handlers.Logout)

	r.Run(":8080")
}
