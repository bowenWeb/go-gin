package controller

import (
	"bo-gin/common"
	"bo-gin/dto"
	"bo-gin/model"
	"bo-gin/response"
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
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于6位")
		return
	}
	if IsPhoneExist(DB, phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已存在")
		return
	}
	// 如果用户不存在，密码符合规则，新建用户
	cryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(cryptPassword),
	}
	DB.Create(&newUser)

	response.Success(
		ctx,
		gin.H{
			"msg": "",
		},
		"注册成功",
	)
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	var user model.User
	DB.Where("phone =?", phone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 500, nil, "密码错误")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 500, nil, "系统错误")
		log.Printf("token 生成 失败")
		return
	}

	response.Success(
		ctx,
		gin.H{
			"token": token,
		},
		"登录成功",
	)
}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(
		ctx,
		gin.H{"user": dto.ToUserDto(user.(model.User))},
		"获取成功",
	)
}

func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
