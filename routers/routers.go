package routers
import (
	"bo-gin/controller"
  "github.com/gin-gonic/gin"
)

func CreateRouter(r *gin.Engine)*gin.Engine {
	
  r.POST("/api/user/register", controller.Register)

	return r
}