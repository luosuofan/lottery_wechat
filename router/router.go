package router

import (
	"github.com/gin-gonic/gin"
	"lottery_wechat/api"
)

func SetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //不用debug，打印垃圾信息很多
	r := gin.Default()
	group := r.Group("/lottery_wechat") //进行分组
	group.GET("/hello", api.Hello)
	return r
}
