package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/pkg/qiniu"
	"nursing_work/pkg/weixin"
	"nursing_work/service"
	"nursing_work/utils"
)

func AppWeChatLogin(c *gin.Context) {
	code := c.Query("code") //  获取code
	// 根据code获取 openID 和 session_key
	wxLoginResp, err := weixin.WXLogin(code)
	if err != nil {
		utils.SendError(c, 403, err.Error())
		return
	}
	utils.SendResponse(c, "hao", wxLoginResp)
	ID := wxLoginResp.OpenId
	tokenStr, err := service.CreateToken(ID)
	//如果是首次就创建
	if !model.Exist(ID) {
		//获取头像和昵称
		accessToken, _, err1 := weixin.GetAccessToken(code)
		if err1 != nil {
			utils.SendError(c, 403, err.Error())
			return
		}
		nickname, avatar, err2 := weixin.GetUserInfo(accessToken, ID)
		if err2 != nil {
			utils.SendError(c, 403, err.Error())
			return
		}
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

func Avatar(c *gin.Context) {
	image, _, err := qiniu.UploadFile(c)
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
