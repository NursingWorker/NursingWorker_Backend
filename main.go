package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"nursing_work/config"
	"nursing_work/model"
	"nursing_work/router"
	"nursing_work/utils"
)

func main() {
	if err := config.Init("/home/nursing/conf/config.yaml", ""); err != nil {
		log.Println(err)
	}
	if err := model.InitDB(); err != nil {
		log.Println(err)
	}
	utils.Init()
	engine := gin.Default()
	router.Register(engine)
	engine.Run("0.0.0.0:4386")

}
