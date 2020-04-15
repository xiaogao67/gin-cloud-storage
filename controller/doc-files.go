package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DocFiles(c *gin.Context)  {
	c.HTML(http.StatusOK, "doc-files.html", nil)
}
