package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Files(c *gin.Context)  {
	c.HTML(http.StatusOK, "files.html", nil)
}
