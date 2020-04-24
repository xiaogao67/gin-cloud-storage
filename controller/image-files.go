package controller

import (
	"file-store/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImageFiles(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := model.GetUserInfo(openId)

	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)
	//获取图像类型文件
	imgFiles := model.GetTypeFile(2, user.FileStoreId)

	c.HTML(http.StatusOK, "image-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"imgFiles":      imgFiles,
		"imgCount":      len(imgFiles),
		"currImg":       "active",
		"currClass":     "active",
	})
}
