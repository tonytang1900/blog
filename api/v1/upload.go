package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, header, _ := c.Request.FormFile("file")
	filesize := header.Size

	fmt.Println("=================", header.Size)
	imgUrl, code := model.UploadFile(file, filesize)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : errmsg.Errmsg(code),
		"imgUrl" : imgUrl,
	})
}
