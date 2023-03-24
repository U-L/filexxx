package api

import (
	"filexxx/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetHost 设置Ip初始值 和 刷新保留Ip
func SetHost(host string) string {
	if host == "" {
		host = global.Address.Addr[0]
	} else if host == "undefined" {
		host = global.Address.MyIPTmp
	}
	return host
}

//GetIpData 获取所有ip api
func GetIpData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":  global.Address.Addr,
		"tmpIp": SetHost(global.Address.MyIP),
	})
}

// SelectIp 设置ip
func SelectIp(c *gin.Context) {
	if url := c.Query("url"); url != "" {
		global.Address.MyIP = SetHost(url)
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求成功",
		})
		global.Address.MyIPTmp = global.Address.MyIP
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求出错",
		})
	}
}


