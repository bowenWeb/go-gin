package controller

import (
	"bo-gin/common"
	"bo-gin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码长度必须大于6位",
		})
		return
	}
	if IsPhoneExist(DB, phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号已存在",
		})
		return
	}
	// 如果用户不存在，密码符合规则，新建用户
	cryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(cryptPassword),
	}
	DB.Create(&newUser)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	var user model.User
	DB.Where("phone =?", phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"code": 422, "msg": "用户不存在"},
		)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 500, "msg": "密码错误"})
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 500, "msg": "系统错误"})
		log.Printf("token 生成 失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
		},
		"msg": "登录成功",
	})
}

func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
