package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OtherFiles(c *gin.Context)  {
	c.HTML(http.StatusOK, "other-files.html", nil)
}
