package model

import (
	"file-store/model/mysql"
	"file-store/util"
	"strings"
	"time"
)

type Share struct {
	Id       int
	Code     string
	FileId   int
	Username string
	Hash     string
}

//创建分享
func CreateShare(code, username string, fId int) string {
	share := Share{
		Code:     strings.ToLower(code),
		FileId:   fId,
		Username: username,
		Hash:     util.EncodeMd5(code + string(time.Now().Unix())),
	}
	mysql.DB.Create(&share)

	return share.Hash
}

//查询分享
func GetShareInfo(f string) (share Share) {
	mysql.DB.Find(&share, "hash = ?", f)
	return
}

//校验提取码
func VerifyShareCode(fId, code string) bool {
	var share Share
	mysql.DB.Find(&share, "file_id = ? and code = ?", fId, code)
	if share.Id == 0 {
		return false
	}
	return true
}