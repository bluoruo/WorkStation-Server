package main

import (
	"wsserver/lib"
	"wsserver/server/web/route"
)

func init() {
	/* 加载基础配置信息 */
	lib.LoadConfig()
}

// http server
func runHttpServer() {
	r := route.InitRouter()
	err := r.Run(lib.ServerConfig.Listen + ":" + lib.ServerConfig.HttpPort)
	if err != nil {
		return
	}
}

// 入口
func main() {
	// 启动http
	runHttpServer()

}
