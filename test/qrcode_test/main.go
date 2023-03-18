package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
)

func GetQRCodeFile(content, outPath string, level qrcode.RecoveryLevel, size int) interface{} {
	//固定方法
	err := qrcode.WriteFile(content, level, size, outPath)
	if err != nil {
		return err.Error()
	}
	return nil
}
func GetQrCodes(c *gin.Context) {
	if pa:=c.Param("purl");pa != ""{
		purl :="" + GetQRFunc(pa)
		c.JSON(http.StatusOK,gin.H{
			"base64": purl,
		})
	}else {
		c.JSON(http.StatusNotFound,gin.H{
			"error":"没有链接",
		})
	}
}

func GetQRFunc(c string) string {
	encode, _ := qrcode.Encode(c, qrcode.Medium, 256)
	return base64.StdEncoding.EncodeToString(encode)
}

func main()  {
	r:= gin.Default()
	r.GET("/qrcode/:purl",GetQrCodes)
	r.Run(":8888")
}
