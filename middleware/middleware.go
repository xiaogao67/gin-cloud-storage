package middleware

import (
	"file-store/lib"
	"file-store/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//检查是否登录中间件
func CheckLogin(c *gin.Context)  {
	token, err := c.Cookie("Token")
	if err != nil {
		fmt.Println("cookie", err.Error())
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}

	openId, err := lib.GetKey(token)
	if err != nil {
		fmt.Println("Get Redis Err:", err.Error())
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}

	user := model.GetUserInfo(openId)

	if user.Id == 0 {
		//校验失败 返回登录页面
		c.Redirect(http.StatusFound, "/")
	} else {
		//校验成功 继续执行
		c.Set("openId", openId)
		c.Next()
	}
}
