package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context) {
	var categ model.Category

	_ = c.ShouldBindJSON(&categ)
	code := model.DoesCategoryExist(categ.Name)

	if code == errmsg.ERROR_CATEGORY_NOT_EXIST {
		errmesg := model.CreateCategory(&categ)
		code = errmesg
	}

	if code == errmsg.ERROR_CATEGORYNAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.Errmsg(code),
		"data": categ,
	})
}

//查询单个分类下的文章

//查询分类列表
func ListCategories(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))

	categs := model.GetCategories(offset, pagesize)

	c.JSON(http.StatusOK, gin.H{
		"code":  errmsg.Success,
		"msg":   errmsg.Errmsg(errmsg.Success),
		"users": categs,
	})
}
//编辑分类
func EditCategory(c *gin.Context) {
	cid,_ := strconv.Atoi(c.Param("id"))

	var categ model.Category
	_ = c.ShouldBindJSON(&categ)

	code := model.DoesCategoryExist(categ.Name)

	if code == errmsg.ERROR_CATEGORYNAME_USED {
		c.Abort()
	}
	if code == errmsg.ERROR_CATEGORY_NOT_EXIST {
		code = model.UpdateCategory(cid, categ)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : errmsg.Errmsg(code),
	})

}
//删除分类
func DeleteCategory(c *gin.Context) {

	cid, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteCategory(cid)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.Errmsg(code),
	})
}