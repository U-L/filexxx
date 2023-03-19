package main

import (
	"filexxx/global"
	"filexxx/initalize"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//初始化日志库
	initalize.InitLogger()
	//init ip
	initalize.InitGetIp()
	//初始化路由
	router := initalize.Routers()

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
