package handler

import (
	"github.com/JLynnLee/go-blog/model"
	"github.com/gin-gonic/gin"
)

func GetPosts(context *gin.Context) {
	model.GetPosts(context)
}

func CreatePost(context *gin.Context) {
	model.CreatePost(context)
}

func Login(context *gin.Context) {
	model.Login(context)
}

func Register(context *gin.Context) {
	model.Register(context)
}

func Logout(context *gin.Context) {
	model.Logout(context)
}
