package conf

import (
	"github.com/Bearki/BeDisk/tools"
	"github.com/bearki/belog"
)

// Log 全局日志对象
var Log *belog.BeLog

// initLog 初始化日志引擎
func initLog() {
	// 初始化文件日志引擎
	Log = belog.New(belog.EngineFile, belog.FileEngineOption{
		LogPath: tools.JoinPath(Config.Log.DirPath, "bedisk.log"), // 日志文件储存路径
		MaxSize: Config.Log.MaxSize,                               // 单文件最大容量（单位：MB）
		SaveDay: Config.Log.SaveDay,                               // 日志保存天数
	}).SetSkip(1).OpenFileLine() // 配置函数栈层级 及 开启文件行数打印

	// 判断当前是开发模式还是正式模式
	if App.IsDevMode {
		// 开启控制台日志输出
		Log.SetEngine(belog.EngineConsole, nil)
	}
}
