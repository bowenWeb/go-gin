package common

import (
	"bo-gin/model"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func InitDB()*gorm.DB {
	db, err := gorm.Open("mysql", "root:admin123@tcp(localhost:3306)/ginessential?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database, err:"+ err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB()*gorm.DB{
	return DB
}