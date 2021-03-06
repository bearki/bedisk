package routers

import (
	"github.com/bearki/bedisk/conf"
	"github.com/gin-gonic/gin"
)

// Init 路由初始化入口
// @return *gin.Engine 路由句柄指针
func Init() *gin.Engine {

	conf.Log.Info("router initializing···")

	// 创建一个空路由
	r := gin.New()

	// 初始化API接口路由
	initApiRouter(r)

	// 初始化WEB网页路由
	initWebRouter(r)

	conf.Log.Info("router init success")

	// 返回初始化好的路由
	return r
}
