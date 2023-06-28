package like

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
	if err := model.LikeCt(userID, postID); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "like success", nil)
}

func Delete(c *gin.Context) {
	postID := c.Query("postID")
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	if !model.IsMyLike(userID, postID) {
		utils.SendError(c, 403, "the like is not yours")
		return
	}
	if err := model.LikeDt(userID, postID); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "cancel like success", nil)
}
