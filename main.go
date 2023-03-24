package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"filexxx/api/wsocket"
	"filexxx/global"
	"filexxx/initalize"
)

//go:embed template/*
var f embed.FS

func main() {
	//创建文件夹
	initalize.InitMkd()
	//init命令
	initalize.InitArgs()

	//初始化日志库
	initalize.InitLogger()
	//init ip
	initalize.InitGetIp()
	//初始化路由
	router := initalize.Routers()

	hub := wsocket.NewHub()

	go hub.Run()
	router.GET("/ws", func(c *gin.Context) {
		wsocket.WsApi(c, hub)
	})

	sfile, _ := fs.Sub(f,"template")
	router.StaticFS("/static/", http.FS(sfile))
	//router.StaticFile("/favicon.ico","./template/favicon.ico")
	router.StaticFile("/favicon.ico","./template/favicon.ico")
	//open
	initalize.InitOpen()
	//优雅退出
	go func() {
		if err := router.Run(fmt.Sprintf(":%d",global.MyConfig.Port)); err != nil {
			zap.S().Info("程序启动失败: ", err.Error())
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("程序退出成功")
}
