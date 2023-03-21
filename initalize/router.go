package initalize

import (
	"filexxx/router"
	"github.com/gin-gonic/gin"
)


func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()

	//Router.MaxMultipartMemory = 1 << 20

	//router group
	//Router.Routes()

	ApiGroup := Router.Group("/v1")
	router.InitBaseRouter(ApiGroup)
	router.InitQrRouter(ApiGroup)
	router.InitFileRouter(ApiGroup)

	return Router
}