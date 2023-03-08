package route

import (
	"github.com/gin-gonic/gin"
	bot "imbot/server/web/api/v1"
	wss "imbot/server/web/wss/v2"
)

func InitRouter() *gin.Engine {
	//r := gin.New()
	r := gin.Default()

	/* 通用参数记录 */
	// Path:/robot/receive	- 通用返回信息接受
	r.GET("/robot/receive/:type", bot.Receive)
	r.POST("/robot/receive/:type", bot.Receive)

	/* WorkStation Server */
	//基本信息
	r.GET("/v2/config", wss.WssConfig)
	r.POST("/v2/config", wss.WssConfig)
	r.POST("/v2/config.php", wss.WssConfig)
	//Base信息 相关
	r.POST("/v2/base", wss.WssBase)
	//ddns信息 相关
	r.POST("/v2/ddns", wss.WssDDns)
	//repo上报 基本信息 相关
	r.POST("/v2/repo", wss.WssRepo)
	//status 上报 客户端状态 相关
	r.POST("/v2/status", wss.WssStatus)

	/* im Robot */
	// Path:/robot/v1/dingtalk	- 钉钉接口
	imbot := r.Group("/robot/v1/dingtalk")
	{
		// 接收用户发来的消息
		imbot.POST("/receiveMessage", bot.ReceiveDingTalkMessage)
	}

	// 结束返回 r
	return r

}
