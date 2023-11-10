package hire

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/utils"
)

func Create(c *gin.Context) {
	carerID := c.Query("CarerID")
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	err := model.HireCt(carerID, userID)
	if err != nil {
		utils.SendResponse(c, "雇佣失败", nil)
		return
	}
	utils.SendResponse(c, "雇佣成功", nil)
}

func Delete(c *gin.Context) {
	carerID := c.Query("CarerID")
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	err := model.HireDt(carerID, userID)
	if err != nil {
		utils.SendResponse(c, "解除雇佣失败", nil)
		return
	}
	utils.SendResponse(c, "解除雇佣成功", nil)
}

func Verify(c *gin.Context) {
	carerID := c.Query("CarerID")
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	hire := model.IsHire(carerID, userID)
	if !hire {
		utils.SendResponse(c, "未雇佣", nil)
		return
	}
	utils.SendResponse(c, "已雇佣", nil)
}
