package like

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/utils"
)

// @Summary 点赞
// @Description 创建点赞记录
// @Tags 点赞
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param postID query string true "帖子ID"
// @Success 200 {string} utils.Response "like success"
// @Failure 400 {string} utils.Error "错误信息"
// @Router /api/v1/like [post]
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

// @Summary 取消点赞
// @Description 取消点赞记录
// @Tags 点赞
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param postID query string true "帖子ID"
// @Success 200 {string} utils.Response "cancel like success"
// @Failure 400 {string} utils.Error "错误信息"
// @Router /api/v1/like [delete]
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
