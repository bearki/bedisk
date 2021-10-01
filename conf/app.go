package conf

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// GlobalData 程序全局配置结构
type GlobalData struct {
	IsDevMode bool   // 是否开发模式
	WorkPath  string // 工作路径
}

// App 应用程序全局配置
var App GlobalData

// InitApp 初始化应用程序
func InitApp() {
	// 初始化环境
	initEnv()
	// 初始化配置文件模块
	initConfig()
	// 初始化文件夹创建
	initMkdir()
	// 初始化日志引擎
	initLog()
}

// initEnv 初始化环境
func initEnv() {
	// 判断是否为开发环境
	if os.Getenv("BEDISK_IS_DEV") == "1" {
		App.IsDevMode = true // 赋值为开发模式
	}
	// 预定义error，减少变量开销
	var err error
	// 赋值工作路径
	if App.IsDevMode { // 开发模式
		App.WorkPath, err = os.Getwd() // 获取当前执行路径
		if err != nil {
			err = fmt.Errorf("get application work path error: %s", err.Error())
			panic(err.Error())
		}
	} else { // 正式模式
		exePath, err := os.Executable() // 获取可执行文件所在路径
		if err != nil {
			err = fmt.Errorf("get application work path error: %s", err.Error())
			panic(err.Error())
		}
		App.WorkPath = filepath.Dir(exePath) // 赋值工作目录
	}
}

// initMkdir 初始化创建文件夹
func initMkdir() {
	// 定义需要创建的文件夹
	dirList := []string{
		path.Dir(Config.Log.DirPath),
	}
	// 执行创建
	for _, dir := range dirList {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			err = fmt.Errorf("create application dir error: %s", err.Error())
			panic(err.Error())
		}
	}
}
