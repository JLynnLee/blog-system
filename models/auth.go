package models

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func GetPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var posts []Post
	db.Preload("User").Where("user_id = ?", c.MustGet("user_id")).Find(&posts)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Posts": posts,
	})
}

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 校验用户名和密码格式
	if !isValidUsername(user.Username) || !isValidPassword(user.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户名或密码不符合要求" + user.Username + "——" + user.Password})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	// 检查用户名是否已存在
	var existingUser User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "用户名已存在"})
		return
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}

// isValidUsername 校验用户名格式
func isValidUsername(username string) bool {
	return len(username) >= 3 && len(username) <= 32
}

// isValidPassword 校验密码强度
func isValidPassword(password string) bool {
	return len(password) >= 6
}

func Login(c *gin.Context) {
	var input struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var user User
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	session, _ := Store.Get(c.Request, "session-name")
	session.Values["user_id"] = user.ID
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  "无法保存 session",
			"detail": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func CreatePost(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var post Post
	if err := c.ShouldBind(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID = userID
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&post)

	c.Redirect(http.StatusSeeOther, "/")
}

func Logout(c *gin.Context) {
	session, _ := Store.Get(c.Request, "session-name")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusSeeOther, "/login")
}
