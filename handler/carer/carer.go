package carer

import (
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/utils"
)

func Recmt(c *gin.Context) {
	//number := c.Query("number")

}

func Search(c *gin.Context) {
	point := c.Query("point")
	//start := c.Query("start")
	//number := c.Query("number")
	carers, err := model.CarerSc(point)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "search carers success", carers)
}

func Type(c *gin.Context) {
	//number := c.Query("number")
	//tp := c.Query("type")

}

func View(c *gin.Context){

}

func IsHire(c *gin.Context){

}

func ViewCt(c *gin.Context){

}

func ViewDt(c *gin.Context){

}