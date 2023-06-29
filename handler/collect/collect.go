package collect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/utils"
)

// @Summary 收藏
// @Description 创建收藏记录
// @Tags 收藏
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param postID query string true "帖子ID"
// @Success 200 {string} utils.Response "collect success"
// @Failure 400 {string} utils.Error "错误信息"
// @Router /api/v1/collect [post]
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

// @Summary 取消收藏
// @Description 取消收藏记录
// @Tags 收藏
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param postID query string true "帖子ID"
// @Success 200 {string} utils.Response "cancel collect success"
// @Failure 400 {string} utils.Error "错误信息"
// @Router /api/v1/collect [delete]
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
