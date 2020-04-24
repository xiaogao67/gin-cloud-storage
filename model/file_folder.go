package model

import (
	"file-store/model/mysql"
	"fmt"
	"strconv"
	"time"
)

//文件夹表
type FileFolder struct {
	Id             int
	FileFolderName string
	ParentFolderId int
	FileStoreId    int
	Time           string
}

//新建文件夹
func CreateFolder(folderName, parentId string, fileStoreId int) {
	parentIdInt, err := strconv.Atoi(parentId)
	if err != nil {
		fmt.Println("父类id错误")
		return
	}
	fileFolder := FileFolder{
		FileFolderName: folderName,
		ParentFolderId: parentIdInt,
		FileStoreId:    fileStoreId,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
	}
	mysql.DB.Create(&fileFolder)
}

//获取父类的id
func GetParentFolder(fId string) (fileFolder FileFolder) {
	mysql.DB.Find(&fileFolder, "id = ?", fId)
	return
}

//获取目录所有文件夹
func GetFileFolder(parentId string, fileStoreId int) (fileFolders []FileFolder) {
	mysql.DB.Order("time desc").Find(&fileFolders, "parent_folder_id = ? and file_store_id = ?", parentId, fileStoreId)
	return
}

//获取当前的目录信息
func GetCurrentFolder(fId string) (fileFolder FileFolder) {
	mysql.DB.Find(&fileFolder, "id = ?", fId)
	return
}

//获取当前路径所有的父级
func GetCurrentAllParent(folder FileFolder, folders []FileFolder) []FileFolder {
	var parentFolder FileFolder
	if folder.ParentFolderId != 0 {
		mysql.DB.Find(&parentFolder, "id = ?", folder.ParentFolderId)
		folders = append(folders, parentFolder)
		//递归查找当前所有父级
		return GetCurrentAllParent(parentFolder, folders)
	}

	//反转切片
	for i, j := 0, len(folders)-1; i < j; i, j = i+1, j-1 {
		folders[i], folders[j] = folders[j], folders[i]
	}

	return folders
}

//获取用户文件夹数量
func GetUserFileFolderCount(fileStoreId int) (fileFolderCount int) {
	var fileFolder []FileFolder
	mysql.DB.Find(&fileFolder, "file_store_id = ?", fileStoreId).Count(&fileFolderCount)
	return
}

//删除文件夹信息
func DeleteFileFolder(fId string) bool {
	var fileFolder FileFolder
	var fileFolder2 FileFolder
	//删除文件夹信息
	mysql.DB.Where("id = ?", fId).Delete(FileFolder{})
	//删除文件夹中文件信息
	mysql.DB.Where("parent_folder_id = ?", fId).Delete(MyFile{})
	//删除文件夹中文件夹信息
	mysql.DB.Find(&fileFolder, "parent_folder_id = ?", fId)
	mysql.DB.Where("parent_folder_id = ?", fId).Delete(FileFolder{})

	mysql.DB.Find(&fileFolder2, "parent_folder_id = ?", fileFolder.Id)
	if fileFolder2.Id != 0 {  //递归删除文件下的文件夹
		return DeleteFileFolder(strconv.Itoa(fileFolder.Id))
	}

	return true
}

//修改文件夹名
func UpdateFolderName(fId, fName string)  {
	var fileFolder FileFolder
	mysql.DB.Model(&fileFolder).Where("id = ?", fId).Update("file_folder_name", fName)
}