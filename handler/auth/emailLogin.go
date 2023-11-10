package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nursing_work/utils"
)

type EmailLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func EmailLogin(c *gin.Context) {
	req := new(EmailLoginReq)
	err := c.Bind(req)
	if err != nil {
		utils.SendError(c, http.StatusOK, err.Error())
		return
	}

}
