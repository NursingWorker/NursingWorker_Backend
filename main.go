package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"nursing_work/config"
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
	engine := gin.Default()
	engine.LoadHTMLGlob("/home/nursing/static/*")
	router.Register(engine)
	if err := engine.Run("0.0.0.0:4386");err != nil {
		log.Println(err)
}
}
