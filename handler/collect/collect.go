package collect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/utils"
)

func Create(c *gin.Context) {
	postID := c.Query("postID")
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	if err := model.CltCt(userID, postID); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "collect success", nil)
}

func Delete(c *gin.Context) {
	postID := c.Query("postID")
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	if !model.IsMyClt(userID, postID) {
		utils.SendError(c, 403, "the collection is not yours")
		return
	}
	if err := model.CltDt(userID, postID); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "cancel collect success", nil)
}
