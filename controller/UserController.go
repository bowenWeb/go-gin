package controller
import (
	"bo-gin/model"
	"bo-gin/common"
	"net/http"
	"github.com/jinzhu/gorm"
  "github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

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
	if IsPhoneExist(DB,phone) {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"msg":"手机号已存在",
		})
		return
	}
	// 如果用户不存在，密码符合规则，新建用户
	newUser := model.User {
		Name:name,
		Phone:phone,
		Password:password,
	}
	DB.Create(&newUser)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}


func IsPhoneExist(db *gorm.DB,phone string) bool{
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}