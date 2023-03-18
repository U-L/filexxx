package api

import (
	"filexxx/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIpData(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"data": global.Address.Addr,
	})
}

func SelectIp(c *gin.Context){
	if url:=c.Query("url"); url!=""{
		global.MyIP = url
		c.JSON(http.StatusOK,gin.H{
			"msg":"请求成功",
		})
	}else{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"请求出错",
		})
	}

}