package router

import (
	"filexxx/api/qrcode"
	"filexxx/middleware"
	"github.com/gin-gonic/gin"
)

func InitQrRouter(Router *gin.RouterGroup)  {
	QrRouter := Router.Group("qr").Use(middleware.GinLogger())
	{
		QrRouter.GET("/qrcode/:purl",qrcode.QrCodesDown)
		QrRouter.GET("/qrview",qrcode.GetQrBase)
	}
}