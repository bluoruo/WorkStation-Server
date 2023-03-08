package v2

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"imbot/lib"
	"imbot/server/web/model"
	"log"
	"net/http"
	"strconv"
)

type returnExecJson struct {
	Code string        `json:"code"`
	Msg  string        `json:"msg"`
	Data structComData `json:"data"`
}

type structComData struct {
	Type int    `json:"type"`
	Com  string `json:"com"`
}

var reExecJson = &returnExecJson{}

// WssStatus 接收客户端上报状态
func WssStatus(c *gin.Context) {
	byteJson, _ := c.GetRawData()
	var data map[string]interface{}
	err := json.Unmarshal(byteJson, &data)
	if err != nil {
		log.Println("参数错误！")
	}
	//处理 int默认变成float64的字段
	wsId, _ := strconv.Atoi(fmt.Sprintf("%1.0f", data["ws_client_id"].(float64)))
	status, _ := strconv.Atoi(fmt.Sprintf("%1.0f", data["status"].(float64)))
	delete(data, "ws_client_id")
	delete(data, "status")
	data["ws_client_id"] = wsId
	data["status"] = status
	//插入执行客户端命令
	var reMsg string
	execInfo, est := checkExec(wsId)
	if est {
		reMsg = "success"
	} else {
		reMsg = "ok"
	}
	//处理上报数据
	id, st := model.ExistWscStatus(wsId)
	if st { //存在，修改
		if model.EditWscStatus(id, data) {
			if st { //有客户端指令
				reExecJson.Code = "0"
				reExecJson.Msg = reMsg
				reExecJson.Data = execInfo
				c.JSON(http.StatusOK, reExecJson)
			}
		} else {
			c.String(http.StatusOK, lib.ReturnErr("提交 status 失败！"))
		}
	} else { //不存在，添加
		if model.AddWscStatus(data) {
			if st { //有客户端指令
				reExecJson.Code = "0"
				reExecJson.Msg = reMsg
				reExecJson.Data = execInfo
				c.JSON(http.StatusOK, reExecJson)
			}
		} else {
			c.String(http.StatusOK, lib.ReturnErr("提交 status 失败！"))
		}
	}
}

// 检查是否客户端执行命令
func checkExec(wId int) (structComData, bool) {
	shellId, st := model.ExistWscExecShellId(wId)
	var res structComData
	if st {
		id, est := model.ExistWscExecInfo(shellId)
		if est {
			execInfoData := model.GetWscExecInfo(id)
			res.Type = execInfoData.Type
			res.Com = execInfoData.Com
			return res, true
		}
	}
	return res, false
}
