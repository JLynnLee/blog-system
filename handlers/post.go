package handlers

import (
	"github.com/JLynnLee/go-blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var posts []models.Post
	db.Preload("User").Find(&posts)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Posts": posts,
	})
}

func CreatePost(context *gin.Context) {
	models.CreatePost(context)
}

func Login(context *gin.Context) {
	models.Login(context)
}

func Register(context *gin.Context) {
	models.Register(context)
}
