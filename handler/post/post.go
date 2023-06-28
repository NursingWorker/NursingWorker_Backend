package post

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/pkg/qiniu"
	"nursing_work/service"
	"nursing_work/utils"
)

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

func Cmt(c *gin.Context) {
	postID := c.Query("postID")
	cmts, err := model.Cmt(postID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get comment success", cmts)
}

func Rep(c *gin.Context) {
	commentID := c.Query("commentID")
	reps, err := model.Rep(commentID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get replies success", reps)
}

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
