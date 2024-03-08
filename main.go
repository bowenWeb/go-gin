package main

import (
	"fmt"
	"net/http"
  "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)
type User struct {
  gorm.Model
  Name string
	Phone string
  Password string
}

func main() {
	db := InitDB()
	defer db.Close()
  r := gin.Default()
  r.POST("/api/user/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		phone := ctx.PostForm("phone")
		password := ctx.PostForm("password")

		// fmt.Println(name,"name")
		// fmt.Println(phone,"phone")
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":422,
				"msg":"密码长度必须大于6位",
			})
			return
		}
		if IsPhoneExist(db,phone) {
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":422,
				"msg":"手机号已存在",
			})
			return
		}
		// 如果用户不存在，密码符合规则，新建用户
		newUser := User {
			Name:name,
			Phone:phone,
			Password:password,
		}
		db.Create(&newUser)

    ctx.JSON(http.StatusOK, gin.H{
      "msg": "注册成功",
    })
  })
  r.Run()
}

func IsPhoneExist(db *gorm.DB,phone string) bool{
	var user User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func InitDB()*gorm.DB {
	db, err := gorm.Open("mysql", "root:admin123@tcp(localhost:3306)/ginessential?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database, err:"+ err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}