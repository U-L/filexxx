package router

import (
	"filexxx/api"
	"filexxx/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup)  {
	BaseRouter := Router.Group("base").Use(middleware.GinLogger())
	{
		BaseRouter.GET("/ipview",api.GetIpData)
		BaseRouter.GET("/myip",api.SelectIp)
	}
}