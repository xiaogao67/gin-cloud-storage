package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MusicFiles(c *gin.Context)  {
	c.HTML(http.StatusOK, "music-files.html", nil)
}
