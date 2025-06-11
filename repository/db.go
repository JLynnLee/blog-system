package repository

import (
	"fmt"
	"github.com/JLynnLee/go-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("migrations/"+path), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库: " + err.Error())
	}
	// 自动迁移模型
	DBMigrate(db)
	return db
}

func InitMySql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库: " + err.Error())
	}
	fmt.Println(db.Config)
	// 自动迁移模型
	DBMigrate(db)
	return db
}

func DBMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{}, &model.Post{})
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
}
