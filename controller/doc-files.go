package controller

import (
	"file-store/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DocFiles(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := model.GetUserInfo(openId)

	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)
	//获取文档类型文件
	docFiles := model.GetTypeFile(1, user.FileStoreId)

	c.HTML(http.StatusOK, "doc-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"docFiles":      docFiles,
		"docCount":      len(docFiles),
		"currDoc":       "active",
		"currClass":     "active",
	})
}
