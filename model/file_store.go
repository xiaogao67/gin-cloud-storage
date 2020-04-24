package model

import (
	"file-store/model/mysql"
)

type FileStore struct {
	Id          int
	UserId      int
	CurrentSize int64
	MaxSize     int64
}

//根据用户id获取仓库信息
func GetUserFileStore(userId int) (fileStore FileStore) {
	mysql.DB.Find(&fileStore, "user_id = ?", userId)
	return
}

//判断用户容量是否足够
func CapacityIsEnough(fileSize int64, fileStoreId int) bool {
	var fileStore FileStore
	mysql.DB.First(&fileStore, fileStoreId)
	if fileStore.MaxSize - (fileSize/1024) < 0 {
		return false
	}

	return true
}