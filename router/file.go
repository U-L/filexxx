package router

import (
	"filexxx/api/file"
	"filexxx/middleware"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup)  {
	FileRouter := Router.Group("file").Use(middleware.GinLogger())
	{
		FileRouter.POST("upload",file.UploadFile)
		FileRouter.GET("list",file.ListFile)
		FileRouter.GET("download/:name",file.DownloadFile)
		FileRouter.GET("delete/:name",file.RemoveFile)
	}

}