package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/utils"
)

func GetHistory(c *gin.Context) {
	tmp, _ := c.Get("openID")
	Mid := model.UserID(fmt.Sprint(tmp))

	Oid := c.Query("object_id")
	msgs, err := model.FindByObject(Oid, Mid)
	if err != nil {
		utils.SendError(c, 200, "历史消息查询失败")
		return
	}
	utils.SendResponse(c, "历史消息查询成功", msgs)
}

func GetHistoryUser(c *gin.Context) {
	tmp, _ := c.Get("openID")
	Mid := model.UserID(fmt.Sprint(tmp))

	oids, err := model.FindObject(Mid)
	if err != nil {
		utils.SendError(c, 200, "聊天列表查询失败")
		return
	}

	utils.SendResponse(c, "聊天列表查询成功", oids)

}
