// @Title 程序主包
// @Desc 程序入口，一切美好的事将从这里开始
// @Author bearki
// @DateTime 2021/09/20 17:41
package main

import (
	"fmt"

	"github.com/bearki/bedisk/conf"
	"github.com/bearki/bedisk/routers"
	"github.com/gin-gonic/gin"
)

// 初始化操作
func init() {
	// 初始化程序
	conf.InitApp()
}

// 正式开始
func main() {
	// 赋值当前GIN运行模式为release
	if !conf.App.IsDevMode {
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化路由
	app := routers.Init()
	if app == nil {
		panic("applicatopn routers handle is not pointer")
	}
	// 拼接监听地址
	listenAddress := conf.Config.HttpServe.Host
	if conf.Config.HttpServe.Port > 0 {
		listenAddress += fmt.Sprintf(":%d", conf.Config.HttpServe.Port)
	}
	conf.Log.Info("listen HTTP serve address: %s", listenAddress)
	// 启动HTTP服务
	if err := app.Run(listenAddress); err != nil {
		panic(fmt.Sprintf("http server start error: %s", err.Error()))
	}
}
