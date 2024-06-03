package router

import (
	"github.com/gin-gonic/gin"
	"lottery_wechat/api"
)

func SetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)        //不用debug模式，打印垃圾信息很多
	r := gin.Default()                  //视图层进行参数校验，没问题继续走
	group := r.Group("/lottery_wechat") //进行分组
	group.GET("/hello", api.Hello)
	group.POST("/add_prize", api.InitPrize)
	group.GET("/get_prize_info", api.GetPrizeInfo)
	group.GET("/lottery", api.Lottery)
	return r
}
