// @Title 程序主包
// @Desc 程序入口，一切美好的事将从这里开始
// @Author Bearki
// @DateTime 2021/09/20 17:41
package main

import (
	"fmt"

	"github.com/Bearki/BeDisk/configs"
	"github.com/Bearki/BeDisk/routers"
	log "github.com/sirupsen/logrus"
)

// 初始化操作
func init() {
	// 初始化程序
	configs.InitApp()
}

// 正式开始
func main() {
	// 初始化路由
	app := routers.Init()
	if app == nil {
		log.Error("applicatopn routers handle is not pointer")
		panic("applicatopn routers handle is not pointer")
	}
	// 启动HTTP服务
	if err := app.Run(); err != nil {
		log.Error("http server start error: %s", err.Error())
		panic(fmt.Sprintf("http server start error: %s", err.Error()))
	}
}
