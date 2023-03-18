package main

import (
	"filexxx/global"
	"filexxx/initalize"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"
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
	url := fmt.Sprintf("http://%s:%d/static",global.MyConfig.Host,global.MyConfig.Port)
	cmd:=exec.Command(`xdg-open`, url)
	err := cmd.Start()
	if err != nil {
		zap.S().Info("浏览器打开失败 ：",err)
	}
	zap.S().Info("浏览器打开成功 ：",url)

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
