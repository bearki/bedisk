// @Title 程序主包
// @Desc 程序入口，一切美好的事将从这里开始
// @Author Bearki
// @DateTime 2021/09/20 17:41
package main

import "github.com/Bearki/BeDisk/configs"

func init() {
	// 初始化程序
	configs.InitApp()
}

func main() {
	configs.SaveConfigFile()
}
