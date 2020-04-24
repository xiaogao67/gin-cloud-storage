package controller

import (
	"file-store/lib"
	"file-store/model"
	"file-store/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

//上传文件页面
func Upload(c *gin.Context) {
	openId, _ := c.Get("openId")
	fId := c.DefaultQuery("fId", "0")
	//获取用户信息
	user := model.GetUserInfo(openId)
	//获取当前目录信息
	currentFolder := model.GetCurrentFolder(fId)
	//获取当前目录所有的文件夹信息
	fileFolders := model.GetFileFolder(fId, user.FileStoreId)
	//获取父级的文件夹信息
	parentFolder := model.GetParentFolder(fId)
	//获取当前目录所有父级
	currentAllParent := model.GetCurrentAllParent(parentFolder, make([]model.FileFolder, 0))
	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)

	c.HTML(http.StatusOK, "upload.html", gin.H{
		"user":             user,
		"currUpload":       "active",
		"fId":              currentFolder.Id,
		"fName":            currentFolder.FileFolderName,
		"fileFolders":      fileFolders,
		"parentFolder":     parentFolder,
		"currentAllParent": currentAllParent,
		"fileDetailUse": fileDetailUse,
	})
}

//处理上传文件
func HandlerUpload(c *gin.Context) {
	openId, _ := c.Get("openId")
	//获取用户信息
	user := model.GetUserInfo(openId)

	Fid := c.GetHeader("id")
	conf := lib.LoadServerConfig()
	//接收上传文件
	file, head, err := c.Request.FormFile("file")

	//判断当前文件夹是否有同名文件
	if ok := model.CurrFileExists(Fid, head.Filename); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 501,
		})
		return
	}

	//判断用户的容量是否足够
	if ok := model.CapacityIsEnough(head.Size, user.FileStoreId); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 503,
		})
		return
	}

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

	//将光标移至开头
	_, _ = newFile.Seek(0, 0)
	fileHash := util.GetSHA256HashCode(newFile)

	//通过hash判断文件是否已上传过oss
	if ok := model.FileOssExists(fileHash); ok {
		//上传至阿里云oss
		go lib.UploadOss(head.Filename, fileHash)
	}
	//新建文件信息
	model.CreateFile(head.Filename, fileHash, fileSize, Fid, user.FileStoreId)
	//上传成功减去相应剩余容量
	model.SubtractSize(fileSize/1024, user.FileStoreId)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
