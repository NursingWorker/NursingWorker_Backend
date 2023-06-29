package post

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/pkg/qiniu"
	"nursing_work/service"
	"nursing_work/utils"
)

// @Summary  获取推荐的帖子
// @Description 获取推荐的帖子
// @Tags 帖子管理
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param number query string true "数量"
// @Success 200 {object} []model.Post "成功获取推荐的帖子"
// @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/post/recommendation [get]
func Recmt(c *gin.Context) {
	number := c.Query("number")
	var old []string
	//反序列化得到已刷新过的数组
	posts, err := model.ReCmt(old, number)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get posts success", posts)
}

// @Summary  获取评论
// @Description 获取评论
// @Tags 帖子管理
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param postID query string true "帖子ID"
// @Success 200 {object} []model.Comment "成功获取评论"
// @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/post/comment [get]
func Cmt(c *gin.Context) {
	postID := c.Query("postID")
	cmts, err := model.Cmt(postID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get comment success", cmts)
}

// @Summary  获取评论的回复
// @Description 获取评论的回复
// @Tags 帖子管理
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param commentID query string true "评论ID"
// @Success 200 {object} []model.Reply "成功获取回复"
// @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/post/comment/reply [get]
func Rep(c *gin.Context) {
	commentID := c.Query("commentID")
	reps, err := model.Rep(commentID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get replies success", reps)
}

// @Summary  发布帖子
// @Description 发布帖子
// @Tags 帖子管理
// @Accept multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param title formData string true "标题"
// @Param content formData string true "内容"
// @Success 200 {object} utils.Response "成功发布帖子"
// @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/post [post]
func PtCt(c *gin.Context) {
	images, videos, err := qiniu.UploadFile(c)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	var post model.Post
	if err = c.ShouldBind(&post); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	openID, _ := c.Get("openID")
	post.OpenID = fmt.Sprint(openID)
	service.CreatePost(post, images, videos)
	utils.SendResponse(c, "send post success", nil)
}

// @Summary  删除帖子
// @Description 删除帖子
// @Tags 帖子管理
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param postID query string true "帖子ID"
// @Success 200 {object} utils.Response "成功删除帖子"
// @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/post [delete]
func PtDt(c *gin.Context) {
	postID := c.Query("postID")
	tmp, _ := c.Get("openID")
	openID := fmt.Sprint(tmp)
	if !model.IsMyPost(postID, openID) {
		utils.SendError(c, 403, "the post is not yours")
		return
	}
	model.PtDt(postID)
	utils.SendResponse(c, "delete post success", nil)
}

// @Summary  发布评论
// @Description 发布评论
// @Tags 帖子管理
// @Accept multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param postID formData string true "帖子ID"
// @Param content formData string true "评论内容"
// @Success 200 {object} utils.Response "成功发布评论"
// @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/post/comment [post]
func CmtCt(c *gin.Context) {
	var cmt model.Comment
	err := c.ShouldBind(&cmt)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	tmp, _ := c.Get("openID")
	cmt.UserID = model.UserID(fmt.Sprint(tmp))
	if err = model.CmtCt(cmt); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "send comment success", nil)
}

// @Summary  删除评论
// @Description 删除评论
// @Tags 帖子管理
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param commentID query string true "评论ID"
// @Success 200 {object} utils.Response "成功删除评论"
// @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/post/comment [delete]
func CmtDt(c *gin.Context) {
	cmtID := c.Query("commentID")
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	if !model.IsMyComment(userID, cmtID) {
		utils.SendError(c, 403, "the comment is not yours")
		return
	}
	model.CmtDt(cmtID)
	utils.SendResponse(c, "delete comment success", nil)
}

// @Summary  发布回复
// @Description 发布回复
// @Tags 帖子管理
// @Accept multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param commentID formData string true "评论ID"
// @Param objectID formData string true "对象ID"
// @Param content formData string true "回复内容"
// @Success 200 {object} utils.Response "成功发布回复"
// @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/post/comment/reply [post]
func ReplyCt(c *gin.Context) {
	var rep model.Reply
	if err := c.ShouldBind(&rep); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	tmp, _ := c.Get("openID")
	rep.UserID = model.UserID(fmt.Sprint(tmp))
	if err := model.RepCt(rep); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "send reply success", nil)
}

// @Summary  删除回复
// @Description 删除回复
// @Tags 帖子管理
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param replyID query string true "回复ID"
// @Success 200 {object} utils.Response "成功删除回复"
// @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/post/comment/reply [delete]
func ReplyDt(c *gin.Context) {
	repID := c.Query("replyID")
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	if !model.IsMyReply(userID, repID) {
		utils.SendError(c, 403, "the reply is not yours")
		return
	}
	model.RepDt(repID)
	utils.SendResponse(c, "delete reply success", nil)
}
