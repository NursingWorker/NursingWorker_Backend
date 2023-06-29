package mid

import (
	"github.com/gin-gonic/gin"
	"nursing_work/model"
	"nursing_work/service"
)

func TokenMiddleWare(c *gin.Context) {
	openID := service.GetToken(c)
	if openID == "" || model.TokenLatest(openID) != c.GetHeader("Authorization") {
		c.JSON(200, gin.H{"code": 200, "msg": "权限不足"})
		c.Abort()
		return
	} else {
		c.Set("openID", openID)
		c.Next()
		//根据 id 找到对应用户信息并返回
	}
}
