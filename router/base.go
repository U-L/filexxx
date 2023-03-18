package router

import (
	"filexxx/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup)  {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("/ipview",api.GetIpData)
		BaseRouter.GET("/myip",api.SelectIp)
	}
}