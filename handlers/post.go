package handlers

import (
	"github.com/JLynnLee/go-blog/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(context *gin.Context) {
	models.GetPosts(context)
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

func Logout(context *gin.Context) {
	models.Logout(context)
}
