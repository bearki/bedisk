package configs

import (
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

// initLog 初始化日志引擎
func initLog() {
	// 截取日志路径到文件夹部分
	Config.Log.Path = path.Dir(Config.Log.Path)
	// 创建日志储存文件夹
	err := os.MkdirAll(Config.Log.Path, 0755)
	if err != nil {
		fmt.Printf("create log save dir error: %s\n", err.Error())
		panic(err.Error())
	}
	// 转换日志保存天数
	logSaveDay := time.Duration(Config.Log.SaveDay)
	// 实例化一个文件分割工具
	writer, err := rotatelogs.New(
		fmt.Sprintf(
			"%s/bedisk.%s.log",
			Config.Log.Path, // 文件名(不含后缀)
			`%Y%m%d`,        // 分割日期格式
		), // 最新日志文件名
		rotatelogs.WithLinkName(Config.Log.Path),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour*24*logSaveDay), // 日志最大保存时间
		rotatelogs.WithRotationTime(time.Hour*24),      // 日志切割时间间隔(24小时分割一次)
	)
	if err != nil {
		fmt.Printf("log init error: %s\n", err.Error())
		panic(err.Error())
	}
	// 添加事件绑定,为不同级别设置不同的输出目的
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, nil)
	// 绑定logrus事件
	log.AddHook(lfHook)

	log.Info("log init success")
}
