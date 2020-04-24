package model

import (
	"file-store/model/mysql"
	"file-store/util"
	"path"
	"strconv"
	"strings"
	"time"
)

//文件表
type MyFile struct {
	Id             int
	FileName       string //文件名
	FileHash       string //文件哈希值
	FileStoreId    int    //文件仓库id
	FilePath       string //文件存储路径
	DownloadNum    int    //下载次数
	UploadTime     string //上传时间
	ParentFolderId int    //父文件夹id
	Size           int64  //文件大小
	SizeStr        string //文件大小单位
	Type           int    //文件类型
	Postfix        string //文件后缀
}

//添加文件数据
func CreateFile(filename, fileHash string, fileSize int64, fId string, fileStoreId int) {
	var sizeStr string
	//获取文件后缀
	fileSuffix := path.Ext(filename)
	//获取文件名
	filePrefix := filename[0 : len(filename)-len(fileSuffix)]
	fid, _ := strconv.Atoi(fId)

	if fileSize < 1048576 {
		sizeStr = strconv.FormatInt(fileSize/1024, 10) + "KB"
	} else {
		sizeStr = strconv.FormatInt(fileSize/102400, 10) + "MB"
	}

	myFile := MyFile{
		FileName:       filePrefix,
		FileHash:       fileHash,
		FileStoreId:    fileStoreId,
		FilePath:       "",
		DownloadNum:    0,
		UploadTime:     time.Now().Format("2006-01-02 15:04:05"),
		ParentFolderId: fid,
		Size:           fileSize / 1024,
		SizeStr:        sizeStr,
		Type:           util.GetFileTypeInt(fileSuffix),
		Postfix:        strings.ToLower(fileSuffix),
	}
	mysql.DB.Create(&myFile)
}

//获取用户的文件
func GetUserFile(parentId string, storeId int) (files []MyFile) {
	mysql.DB.Find(&files, "file_store_id = ? and parent_folder_id = ?", storeId, parentId)
	return
}

//文件上传成功减去相应容量
func SubtractSize(size int64, fileStoreId int) {
	var fileStore FileStore
	mysql.DB.First(&fileStore, fileStoreId)

	fileStore.CurrentSize = fileStore.CurrentSize + size/1024
	fileStore.MaxSize = fileStore.MaxSize - size/1024
	mysql.DB.Save(&fileStore)
}

//获取用户文件数量
func GetUserFileCount(fileStoreId int) (fileCount int) {
	var file []MyFile
	mysql.DB.Find(&file, "file_store_id = ?", fileStoreId).Count(&fileCount)
	return
}

//获取用户文件使用明细情况
func GetFileDetailUse(fileStoreId int) map[string]int64 {
	var files []MyFile
	var (
		docCount   int64
		imgCount   int64
		videoCount int64
		musicCount int64
		otherCount int64
	)

	fileDetailUseMap := make(map[string]int64, 0)

	//文档类型
	docCount = mysql.DB.Find(&files, "file_store_id = ? AND type = ?", fileStoreId, 1).RowsAffected
	fileDetailUseMap["docCount"] = docCount
	////图片类型
	imgCount = mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 2).RowsAffected
	fileDetailUseMap["imgCount"] = imgCount
	//视频类型
	videoCount = mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 3).RowsAffected
	fileDetailUseMap["videoCount"] = videoCount
	//音乐类型
	musicCount = mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 4).RowsAffected
	fileDetailUseMap["musicCount"] = musicCount
	//其他类型
	otherCount = mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 5).RowsAffected
	fileDetailUseMap["otherCount"] = otherCount

	return fileDetailUseMap
}

//根据文件类型获取文件
func GetTypeFile(fileType, fileStoreId int) (files []MyFile) {
	mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, fileType)
	return
}

//判断当前文件夹是否有同名文件
func CurrFileExists(fId, filename string) bool {
	var file MyFile
	//获取文件后缀
	fileSuffix := strings.ToLower(path.Ext(filename))
	//获取文件名
	filePrefix := filename[0 : len(filename)-len(fileSuffix)]

	mysql.DB.Find(&file, "parent_folder_id = ? and file_name = ? and postfix = ?", fId, filePrefix, fileSuffix)

	if file.Size > 0 {
		return false
	}
	return true
}

//通过hash判断文件是否已上传过oss
func FileOssExists(fileHash string) bool {
	var file MyFile
	mysql.DB.Find(&file, "file_hash = ?", fileHash)
	if file.FileHash != "" {
		return false
	}
	return true
}

//通过fileId获取文件信息
func GetFileInfo(fId string) (file MyFile) {
	mysql.DB.First(&file, fId)
	return
}

//文件下载次数+1
func DownloadNumAdd(fId string) {
	var file MyFile
	mysql.DB.First(&file, fId)
	file.DownloadNum = file.DownloadNum + 1
	mysql.DB.Save(&file)
}

//删除数据库文件数据
func DeleteUserFile(fId, folderId string, storeId int) {
	mysql.DB.Where("id = ? and file_store_id = ? and parent_folder_id = ?", fId, storeId, folderId).Delete(MyFile{})
}
