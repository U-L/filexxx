package initalize

import (
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func InitMkd()  {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	err := os.MkdirAll(dir+"/tmp/upload", 0777)
	err = os.MkdirAll(dir+"/tmp/log", 0777)
	if err != nil {
		zap.S().Info("创建文件夹失败",err.Error())
	}
}