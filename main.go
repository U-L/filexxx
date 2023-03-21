package main

import (
	"embed"
	"filexxx/global"
	"filexxx/initalize"
	"fmt"
	"go.uber.org/zap"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
