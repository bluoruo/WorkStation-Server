package v2

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"wsserver/lib"
	"wsserver/server/web/model"
)

func WssRepo(c *gin.Context) {
	byteJson, _ := c.GetRawData()
	var data map[string]interface{}
	err := json.Unmarshal(byteJson, &data)
	if err != nil {
		log.Println("提交参数错误！")
	}
	//处理 int默认变成float64的字段
	wsId, _ := strconv.Atoi(fmt.Sprintf("%1.0f", data["ws_client_id"].(float64)))
	delete(data, "ws_client_id")
	data["ws_client_id"] = wsId
	//log.Println(data)
	id, st := model.ExistWscIdWscInfo(wsId)
	if st { //存在，修改
		//log.Println("存在，修改")
		if model.EditWscInfo(id, data) {
			c.String(http.StatusOK, lib.ReturnOK("更新info成功！"))
		} else {
			c.String(http.StatusOK, lib.ReturnErr("提交 info 失败！"))
		}
	} else { //不存在，添加
		//log.Println("不存在，添加")
		if model.AddWscInfo(data) {
			c.String(http.StatusOK, lib.ReturnOK("新增info成功！"))
		} else {
			c.String(http.StatusOK, lib.ReturnErr("提交 info 失败！"))
		}
	}
}
