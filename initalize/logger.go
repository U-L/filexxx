package initalize

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	logger *zap.Logger
	//sugarLogger *zap.SugaredLogger
)

func InitLogger() {
	//logger, _ := zap.NewDevelopment()
	//zap.ReplaceGlobals(logger)
	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel),
		zapcore.NewCore(getConsoleEncoder(),zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
	)
	logger = zap.New(core, zap.AddCaller())
	//sugarLogger = logger.Sugar()
	zap.ReplaceGlobals(logger)
}
func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getLogWriter() zapcore.WriteSyncer {
	//时间定义文件名称
	t := time.Unix(time.Now().Unix(),0).Format("2006-01-02-150405")

	file, _ := os.Create(fmt.Sprintf("./tmp/log/%s.log",t))
	//file, _ := os.Create("./tmp/log/mylog.log")
	return zapcore.AddSync(file)
}
