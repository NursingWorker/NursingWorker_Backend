package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/utils"
)

func View(c *gin.Context) {
	userID := c.Param("id")
	if userID == "0" {
		tmp, _ := c.Get("openID")
		userID = model.UserID(fmt.Sprint(tmp))
	}
	view, err := model.FindView(userID)
	if err != nil {
		utils.SendError(c, 200, err.Error())
		return
	}
	utils.SendResponse(c, "查询到对象做出的评价", view)
}

func Other(c *gin.Context) {
	//carerID := c.Query("carerID")

}
func Create(c *gin.Context) {
	var view model.View
	if err := c.ShouldBind(&view); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	view.UserID = userID
	err := model.ViewCt(view)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "add view success", nil)
}

func Delete(c *gin.Context) {
	viewID := c.Query("viewID")
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	err := model.ViewDt(viewID, userID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "delete view success", nil)
}
