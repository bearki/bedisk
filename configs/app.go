package configs

// GlobalData 程序全局配置结构
type GlobalData struct {
}

// App 应用程序全局配置
var App GlobalData

// InitApp 初始化应用程序
func InitApp() {
	// 初始化配置文件模块
	initConfig()
	// 初始化日志引擎
	initLog()
}
