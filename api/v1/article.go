package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context) {
	var article model.Article

	_ = c.ShouldBindJSON(&article)

	code := model.CreateArticle(&article)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.Errmsg(code),
		"data": article,
	})
}

//查询单个文章
func ListOneArticle(c *gin.Context) {
	aid,_ := strconv.Atoi(c.Param("id"))
	article,code := model.GetOneArticle(aid)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : errmsg.Errmsg(code),
		"article" : article,
	})

}

//查询分类下的所有文章
func ListArticlesUnderCategory(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))

	articles, code := model.GetArticlesUnderCategory(cid, offset, pagesize)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : errmsg.Errmsg(code),
		"articlesUnderCateg": articles,
	})

}
//查询文章列表
func ListArticles(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))

	articles := model.GetArticles(offset, pagesize)

	c.JSON(http.StatusOK, gin.H{
		"code":  errmsg.Success,
		"msg":   errmsg.Errmsg(errmsg.Success),
		"users": articles,
	})
}

//编辑文章
func EditArticle(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("id"))
	var article model.Article

	_ = c.ShouldBindJSON(&article)

	code := model.UpdateArticle(aid, article)


	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.Errmsg(code),
	})

}

//删除文章
func DeleteArticle(c *gin.Context) {

	aid, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteArticle(aid)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.Errmsg(code),
	})
}
