package router

import (
	"filexxx/api"
	"filexxx/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup)  {
	BaseRouter := Router.Group("base").Use(middleware.GinLogger(),middleware.Cors())
	{
		BaseRouter.GET("/ipview",api.GetIpData)
		BaseRouter.POST("/myip",api.SelectIp)
	}
}