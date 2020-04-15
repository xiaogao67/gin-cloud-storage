package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VideoFiles(c *gin.Context)  {
	c.HTML(http.StatusOK, "video-files.html", nil)
}
