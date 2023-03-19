package qrcode

import (
	"encoding/base64"
	"filexxx/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
)
func SetUrl(host string,post int,api string) string{
	if host == ""{
		host= global.Address.Addr[0]
	}
	return fmt.Sprintf("http://%s:%d%s",host,post,api)
}

func GetQRFunc(c string) string {
	encode, _ := qrcode.Encode(c, qrcode.Medium, 256)
	return base64.StdEncoding.EncodeToString(encode)
}

func GetQrBase(c *gin.Context) {
	url := SetUrl(global.MyIP,global.MyConfig.Port,"/static")
	fmt.Println("url",url)
	purl :=GetQRFunc(url)
	c.JSON(http.StatusOK,gin.H{
		"base64": "data:image/png;base64,"+purl,
	})
}

func QrCodesDown(c *gin.Context) {
	if pa:=c.Param("purl");pa != ""{
		url := SetUrl(global.MyIP,global.MyConfig.Port,"/v1/file/download/")
		purl :=GetQRFunc((url) +pa)
		c.JSON(http.StatusOK,gin.H{
			"base64": "data:image/png;base64,"+purl,
		})
	}else {
		c.JSON(http.StatusNotFound,gin.H{
			"error":"没有链接",
		})
	}
}
