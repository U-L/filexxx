package router

import (
	"filexxx/api/qrcode"
	"github.com/gin-gonic/gin"
)

func InitQrRouter(Router *gin.RouterGroup)  {
	QrRouter := Router.Group("qr")
	{
		QrRouter.GET("/qrcode/:purl",qrcode.QrCodesDown)
		QrRouter.GET("/qrview",qrcode.GetQrBase)
	}
}