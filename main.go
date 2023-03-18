package main

import (
	"filexxx/initalize"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	//初始化日志库
	initalize.InitLogger()

	//初始化路由
	router := initalize.Routers()
	cmd:=exec.Command(`xdg-open`, `http://127.0.0.1:8090/static`)
	err := cmd.Start()
	if err != nil {
		zap.S().Info("浏览器打开失败 ：",err)
	}
	zap.S().Info("浏览器打开成功 ：","http://127.0.0.1:8090/static")
	//优雅退出
	go func() {
		if err := router.Run(":8090"); err != nil {
			zap.S().Info("程序启动失败: ", err.Error())
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.S().Info("程序退出成功")
}
