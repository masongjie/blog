package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var SugarLog *zap.SugaredLogger

func Initlog() (err error) {
	// 将路由信息输出到gin.log
	gin.DisableConsoleColor()
	f, err := os.Create("./log/gin.log")
	if err != nil {
		fmt.Println("create blogger.log failed,err:", err)
		return
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	return

}

func InitSagurLog() (err error) {
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	f, err := os.OpenFile("./log/test.log",  os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	writeSyncer := zapcore.AddSync(f)
	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)

	logger := zap.New(core)
	SugarLog = logger.Sugar()
	return
}
