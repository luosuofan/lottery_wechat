package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"lottery_wechat/config"
	"lottery_wechat/router"
)

func Init() {
	config.InitGlobalConfig()
}

func main() {
	config1 := config.GetGlobalConfig()
	fmt.Println(config1)
	Init()
	log.Infof("11111111")
	r := router.SetRouter()
	if err := r.Run(fmt.Sprintf(":%d", config1.AppConfig.Port)); err != nil {
		log.Errorf("sever run err: %v", err)
	}
}
