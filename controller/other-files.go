package controller

import (
	"file-store/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OtherFiles(c *gin.Context)  {
	openId, _ := c.Get("openId")
	user := model.GetUserInfo(openId)

	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)
	//获取音频类型文件
	otherFiles := model.GetTypeFile(5, user.FileStoreId)

	c.HTML(http.StatusOK, "other-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"otherFiles":    otherFiles,
		"otherCount":    len(otherFiles),
		"currOther":       "active",
		"currClass":     "active",
	})
}
