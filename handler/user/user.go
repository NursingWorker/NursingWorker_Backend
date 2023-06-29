package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/pkg/qiniu"
	"nursing_work/service"
	"nursing_work/utils"
)




 // @Summary 登录
 // @Description 要一个code
 // @Tags 用户管理
// @Accept  multipart/form-data
 // @Produce application/json
 // @Param code query string true "code"
 // @Success 200 {object} utils.Response "成功获取用户信息"
   // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/login [post]
func AppWeChatLogin(c *gin.Context) {
	//code := c.Query("code") //  获取code
	// 根据code获取 openID 和 session_key
	//wxLoginResp, err := weixin.WXLogin(code)
	//if err != nil {
	//	utils.SendError(c, 403, err.Error())
	//	return
	//}
	//ID := wxLoginResp.OpenId

	ID := c.Query("code")
	tokenStr, err := service.CreateToken(ID)
	//如果是首次就创建
	if !model.Exist(ID) {
		//获取头像和昵称
		//accessToken, _, err1 := weixin.GetAccessToken(code)
		//if err1 != nil {
		//	utils.SendError(c, 403, err.Error())
		//	return
		//}
		//nickname, avatar, err2 := weixin.GetUserInfo(accessToken, ID)
		//if err2 != nil {
		//	utils.SendError(c, 403, err.Error())
		//	return
		//}
		nickname := "一个默认的name"
		avatar := "一个默认的链接"
		user := model.User{
			OpenID:   ID,
			NickName: nickname,
			Avatar:   avatar,
		}
		token := model.Token{
			OpenID: ID,
			Token:  tokenStr,
		}
		if err = model.CreateUser(user, token); err != nil {
			utils.SendError(c, 403, err.Error())
			return
		}
	}

	if err = model.TokenUpdate(tokenStr, ID); err != nil {
		utils.SendError(c, 403, err.Error())
		return
	}
	utils.SendResponse(c, "login success", gin.H{
		"token": tokenStr,
	})
}


 // @Summary  获取用户基本信息
  // @Description 获取用户基本信息
  // @Tags 用户管理
 // @Accept  multipart/form-data
  // @Produce application/json
// @Param Authorization header string true "token"
  // @Param id path string true "用户ID"
  // @Success 200 {object} model.User "成功获取用户信息"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/{id} [get]
func Base(c *gin.Context) {
	userID := c.Param("id")
	if userID == "0" {
		tmp, _ := c.Get("openID")
		userID = model.UserID(fmt.Sprint(tmp))
	}
	user, err := model.Base(userID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get userinfo success", user)
}


 // @Summary  关注用户
  // @Description 关注用户
  // @Tags 用户管理
 // @Accept  multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param idolID query string true "被关注用户ID"
  // @Success 200 {object} utils.Response "成功关注用户"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/subscribe [post]
func Subscribe(c *gin.Context) {
	idolID := c.Query("idolID")
	tmp, _ := c.Get("openID")
	fanID := model.UserID(fmt.Sprint(tmp))
	err := model.FlCt(idolID, fanID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "subscribe success", nil)
}


 // @Summary  取消关注用户
  // @Description 取消关注用户
  // @Tags 用户管理
 // @Accept  multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param idolID query string true "被取消关注用户ID"
  // @Success 200 {object} utils.Response "成功取消关注用户"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/subscribe [delete]
func DisSubscribe(c *gin.Context) {
	idolID := c.Query("idolID")
	tmp, _ := c.Get("openID")
	fanID := model.UserID(fmt.Sprint(tmp))
	err := model.FlDt(idolID, fanID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "dis subscribe success", nil)
}


 // @Summary  获取关注的人
  // @Description 获取关注的人
  // @Tags 用户管理
 // @Accept  multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param id path string true "用户ID"
  // @Success 200 {object} []model.User "成功获取关注的人"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/{id}/subscribe [get]
func Following(c *gin.Context) {
	userID := c.Param("id")
	if userID == "0" {
		tmp, _ := c.Get("openID")
		userID = model.UserID(fmt.Sprint(tmp))
	}
	idols, err := model.Idols(userID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get idols success", idols)
}


 // @Summary  获取粉丝
  // @Description 获取粉丝
  // @Tags 用户管理
 // @Accept  multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param id path string true "用户ID"
  // @Success 200 {object} []model.User "成功获取粉丝"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/{id}/fans [get]
func Fans(c *gin.Context) {
	userID := c.Param("id")
	if userID == "0" {
		tmp, _ := c.Get("openID")
		userID = model.UserID(fmt.Sprint(tmp))
	}
	fans, err := model.Fans(userID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get fans success", fans)
}


 // @Summary  获取已发布的帖子
  // @Description 获取已发布的帖子
  // @Tags 用户管理
 // @Accept  multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param id path string true "用户ID"
  // @Success 200 {object} []model.Post "成功获取已发布的帖子"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/{id}/post [get]
func Post(c *gin.Context) {
	userID := c.Param("id")
	if userID == "0" {
		tmp, _ := c.Get("openID")
		userID = model.UserID(fmt.Sprint(tmp))
	}
	posts, err := model.HisPost(model.OpenID(userID))
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get posts success", posts)
}


 // @Summary  获取收藏的帖子
  // @Description 获取收藏的帖子
  // @Tags 用户管理
 // @Accept  multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param id path string true "用户ID"
  // @Success 200 {object} []model.Post "成功获取收藏的帖子"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/{id}/collection [get]
func Collection(c *gin.Context) {
	userID := c.Param("id")
	if userID == "0" {
		tmp, _ := c.Get("openID")
		userID = model.UserID(fmt.Sprint(tmp))
	}
	hisColl, err := model.HisColl(userID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get posts of collection success", hisColl)
}


 // @Summary  获取点赞的帖子
  // @Description 获取点赞的帖子
  // @Tags 用户管理
 // @Accept  multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param id path string true "用户ID"
  // @Success 200 {object} []model.Post "成功获取点赞的帖子"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/{id}/like [get]
func Like(c *gin.Context) {
	userID := c.Param("id")
	if userID == "0" {
		tmp, _ := c.Get("openID")
		userID = model.UserID(fmt.Sprint(tmp))
	}
	hisLike, err := model.HisLike(userID)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "get posts of like success", hisLike)
}


 // @Summary  更新头像
  // @Description 更新头像
  // @Tags 用户管理
  // @Accept multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param avatar formData file true "头像文件"
  // @Success 200 {object} utils.Response "成功更新头像"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/avatar [put]
func Avatar(c *gin.Context) {
	image,_,err := qiniu.UploadFile(c)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	tmp, _ := c.Get("openID")
	err = model.Avatar(fmt.Sprint(tmp), image[0])
	if err != nil {
		utils.SendError(c, 422, err.Error())
		return
	}
	utils.SendResponse(c, "update avatar success", nil)
}


 // @Summary  更新基本信息
  // @Description 更新基本信息
  // @Tags 用户管理
  // @Accept multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param nickname formData string true "昵称"
  // @Param intro formData string true "简介"
  // @Param age formData string true "年龄"
  // @Param gender formData string true "性别"
  // @Param stature formData string true "身高"
  // @Param address formData string true "地址"
  // @Param experience formData string true "经验/经历"
  // @Success 200 {object} utils.Response "成功更新基本信息"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/info [put]
func InfoUpt(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	tmp, _ := c.Get("openID")
	userID := model.UserID(fmt.Sprint(tmp))
	err := model.InfoUpt(userID, user)
	if err != nil {
		utils.SendError(c, 400, err.Error())
		return
	}
	utils.SendResponse(c, "update userInfo success", nil)
}


 // @Summary  更新身份
  // @Description 更新身份
  // @Tags 用户管理
 // @Accept  multipart/form-data
  // @Produce application/json 
// @Param Authorization header string true "token"
  // @Param identity query string true "身份"
  // @Success 200 {object} utils.Response "成功更新身份"
    // @Failure 400 {object} utils.Error "失败"
// @Router /api/v1/user/identity [put]
func Identity(c *gin.Context) {
	identity := c.Query("identity")
	tmp, _ := c.Get("openID")
	err := model.Identity(fmt.Sprint(tmp), identity)
	if err != nil {
		utils.SendError(c, 422, err.Error())
		return
	}
	utils.SendResponse(c, "update identity success", nil)
}
