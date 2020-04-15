package router

import (
	"file-store/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	router := gin.Default()

	cloud := router.Group("cloud")
	{
		cloud.GET("/index", controller.Index)
		cloud.GET("/files", controller.Files)
		cloud.GET("/upload", controller.Upload)
		cloud.GET("/doc-files", controller.DocFiles)
		cloud.GET("/image-files", controller.ImageFiles)
		cloud.GET("/video-files", controller.VideoFiles)
		cloud.GET("/music-files", controller.MusicFiles)
		cloud.GET("/other-files", controller.OtherFiles)
	}

	{
		cloud.POST("/uploadFile", controller.HandlerUpload)
	}

	return router
}
