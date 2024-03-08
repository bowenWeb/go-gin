package main

import (
	"bo-gin/common"
	"bo-gin/routers"
  "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db := common.InitDB()
	defer db.Close()
  r := gin.Default()
  r = routers.CreateRouter(r)
  r.Run()
}


