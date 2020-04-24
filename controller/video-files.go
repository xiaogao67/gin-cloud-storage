package controller

import (
	"file-store/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VideoFiles(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := model.GetUserInfo(openId)

	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)
	//获取视频类型文件
	videoFiles := model.GetTypeFile(3, user.FileStoreId)

	c.HTML(http.StatusOK, "video-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"videoFiles":    videoFiles,
		"videoCount":    len(videoFiles),
		"currVideo":     "active",
		"currClass":     "active",
	})
}
