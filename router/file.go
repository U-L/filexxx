package router

import (
	"filexxx/api/file"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup)  {
	FileRouter := Router.Group("file")
	{
		FileRouter.POST("upload",file.UploadFile)
		FileRouter.GET("list",file.ListFile)
		FileRouter.GET("download/:name",file.DownloadFile)
	}

}