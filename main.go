package main

import (
	"github.com/JLynnLee/go-blog/config"
	"github.com/JLynnLee/go-blog/middleware"
	"github.com/JLynnLee/go-blog/repository"
	"github.com/JLynnLee/go-blog/route"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

const sqlite = "sqlite"

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("无法加载配置: " + err.Error())
	}

	var db *gorm.DB
	// 初始化数据库
	if cfg.Database.Type == sqlite {
		db = repository.InitSQLite(cfg.Sqlite.Path)
	} else {
		db = repository.InitMySql(cfg.Mysql.Dsn)
	}
	// 初始化 Gin
	engine := gin.Default()
	// 注册中间件
	engine.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("session_secret_key", cfg.Session.SecretKey)
		c.Next()
	}, middleware.ExceptionMiddleware())

	// 设置模板引擎
	engine.LoadHTMLGlob("views/*")

	// 路由
	route.SetupRouter(engine)

	engine.Run(":8080")
}
