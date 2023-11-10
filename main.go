package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"nursing_work/config"
	"nursing_work/handler/chat"
	"nursing_work/model"
	"nursing_work/pkg/qiniu"
	"nursing_work/router"
	"nursing_work/utils"
)

// @title Nursing
// @version 1.1.0
// @description NursingAPI
// @termsOfService http://swagger.io/terrms/
// @contact.name  BIG_DUST
// @contact.email 3264085417@qq.com
// @host 43.138.61.49
// @BasePath /api/v1
// @Schemes http
func main() {
	if err := config.Init("./conf/config.yaml", ""); err != nil {
		log.Println(err)
	}
	if err := model.InitDB(); err != nil {
		log.Println(err)
	}
	utils.Init()
	qiniu.Load()
	chat.Chs = make(map[string]chan string, 1000)
	chat.Conns = make(map[string]*chat.WsConnection, 1000)
	engine := gin.Default()
	router.Register(engine)
	if err := engine.Run(viper.GetString("server_ip")); err != nil {
		log.Println(err)
	}
}
