package router

import (
	"filexxx/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup)  {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("",api.GetView)
	}
}