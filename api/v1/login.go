package v1

import (
	"blog/middleware"
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	var user model.User
	var tokenstr string

	_ = c.ShouldBindJSON(&user)

	code := model.ValidateLogin(user.Username, user.Password)

	if code == errmsg.Success {
		tokenstr, _ = middleware.GenToken(user.Username, user.Password)
	}

	c.Set("authorization", tokenstr)
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : errmsg.Errmsg(code),
		"token" : tokenstr,
	})
}
