package initalize

import (
	"filexxx/global"
	"fmt"
	"go.uber.org/zap"
	"os/exec"
	"runtime"
)

// OsSwitch 选择操作系统对应命令 低分支我使用switch不使用map
func OsSwitch() string{
	switch runtime.GOOS {
	case "linux":
		return "xdg-open"
	case "window":
		return "cmd /c start"
	case "darwin":
		return "open"
	default:
		return ""
	}
}

func InitOpen() {
	url := fmt.Sprintf("http://%s:%d/static",global.MyConfig.Host,global.MyConfig.Port)
	cmd:=exec.Command(OsSwitch(), url)
	err := cmd.Start()
	if err != nil {
		zap.S().Info("浏览器打开失败 ：",err)
	}
	zap.S().Info("浏览器打开成功 ：",url)
}