package routers

import (
	"bo-gin/controller"
	"bo-gin/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRouter(r *gin.Engine) *gin.Engine {

	r.POST("/api/user/register", controller.Register)
	r.POST("/api/user/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.UserInfo)

	return r
}
