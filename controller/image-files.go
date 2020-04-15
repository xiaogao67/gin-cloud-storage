package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImageFiles(c *gin.Context)  {
	c.HTML(http.StatusOK, "image-files.html", nil)
}
