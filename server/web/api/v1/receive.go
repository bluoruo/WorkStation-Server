package v1

import (
	"github.com/gin-gonic/gin"
	"imbot/lib"
	"net/http"
	"strings"
)

var header = make(map[string]string) //Header
var strReturn string                 //返回信息

// 写入日志
func writeLog(recType string, header map[string]string, method string, str string) {
	var strLog string
	for key, value := range header {
		strLog = strLog + key + " ==> " + value + " | "
	}
	strLog = "### New [" + recType + "] Session ###\n Method: " + method + "\n Header: " + strLog +
		"\n Param: " + str + "\n======================================================================="
	lib.ReceiveLog.Println(strLog)
}

// 解析数据 写入日志 并 返回信息
func receiveParams(c *gin.Context, recType string) {
	//请求类型
	method := c.Request.Method
	//参数
	var strParam string
	if method == "POST" { //POST
		bParam, _ := c.GetRawData()
		strParam = string(bParam)
	} else { //Get
		if strings.Contains(c.Request.RequestURI, "?") { //是否存在参数
			strParam = strings.Split(c.Request.RequestURI, "?")[1]
		} else {
			strParam = ""
		}
	}
	writeLog(recType, header, method, strParam) //写入日志
	switch recType {
	case "dingtalk": //钉钉返回信息
		c.String(http.StatusOK, strReturn)
		break
	case "feishu": //飞书返回信息
		c.String(http.StatusOK, receiveFeishu(strParam))
		break
	case "workwechat": //微信返回信息
		break
	}

}

// 处理飞书返回信息
func receiveFeishu(str string) string {
	arr := lib.AnyJson([]byte(str))
	return `{"challenge":"` + arr["challenge"].(string) + `"}`
}

// Receive 通用返回数据处理
func Receive(c *gin.Context) {
	recType := c.Param("type")
	//根据访问类别 构造获取header的数据和返回信息
	header["content-type"] = c.GetHeader("content-type")
	switch recType {
	case "dingtalk":
		header["timestamp"] = c.GetHeader("timestamp")
		header["sign"] = c.GetHeader("sign")
		strReturn = "success"
		break
	case "feishu":
		break
	default:
		strReturn = "ok"
	}
	receiveParams(c, recType)
}
