package controller

import (
	"file-store/lib"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

//上传文件页面
func Upload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

//处理上传文件
func HandlerUpload(c *gin.Context) {
	conf := lib.LoadServerConfig()
	//接收上传文件
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("文件上传错误", err.Error())
		return
	}
	defer file.Close()

	//文件保存本地的路径
	location := conf.UploadLocation + head.Filename

	//在本地创建一个新的文件
	newFile, err := os.Create(location)
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer newFile.Close()

	//将上传文件拷贝至新创建的文件中
	fileSize, err := io.Copy(newFile, file)
	if err != nil {
		fmt.Println("文件拷贝错误", err.Error())
		return
	}
	fmt.Println(fileSize)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
