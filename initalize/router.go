package initalize

import (
	"filexxx/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.StaticFS("/static",http.Dir("./template"))
	//Router.StaticFile("/favicon.ico","./rescoure/favicon.ico")
	//Router.MaxMultipartMemory = 1 << 20
	//router group
	Router.Routes()
	ApiGroup := Router.Group("/v1")
	router.InitBaseRouter(ApiGroup)
	router.InitFileRouter(ApiGroup)

	return Router
}