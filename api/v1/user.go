package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"blog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加用户
func AddUser(c *gin.Context) {
	var u model.User

	_ = c.ShouldBindJSON(&u)
	msg, scode := validator.Validate(&u)
	if scode != errmsg.Success {
		c.JSON(http.StatusOK, gin.H{
			"code" : scode,
			"msg" : msg,
		})
		return
	}

	code := model.DoesUserExist(u.Username)

	if code == errmsg.ERROR_USER_NOT_EXIST {
		errmesg := model.CreateUser(&u)
		code = errmesg
	}

	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.Errmsg(code),
		"data": u,
	})
}

//查询单个用户

//查询用户列表
func ListAllUsers(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))

	users := model.GetUsers(offset, pagesize)

	c.JSON(http.StatusOK, gin.H{
		"code":  errmsg.Success,
		"msg":   errmsg.Errmsg(errmsg.Success),
		"users": users,
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	userid,_ := strconv.Atoi(c.Param("id"))

	var user model.User
	_ = c.ShouldBindJSON(&user)

	code := model.DoesUserExist(user.Username)

	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	if code == errmsg.ERROR_USER_NOT_EXIST {
		code = model.UpdateUser(userid, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : errmsg.Errmsg(code),
	})

}

//删除用户
func DeleteUser(c *gin.Context) {

	userid, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteUser(userid)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.Errmsg(code),
	})
}
