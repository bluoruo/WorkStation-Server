package v2

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"imbot/lib"
	"imbot/server/web/model"
	"log"
	"net/http"
)

// 检查识别名 返回结构
type returnCheckName struct {
	Status string
	Name   string
}

var ginBase *gin.Context

// WssBase 入口
func WssBase(c *gin.Context) {
	ginBase = c
	strType := c.Query("type")
	switch strType {
	case "checkname": //检查识别名
		setWssBase()
		break
	case "edit": //修改识别名
		edtWssBase()
		break
	case "getinfo": //获取基本信息
		getWscBase()
		break
	default:
		c.String(http.StatusOK, "Welcome to WorkStation Server.")
	}
}

// 获取基本信息
func getWscBase() {
	var data map[string]interface{}
	body, _ := ginBase.GetRawData()
	err := json.Unmarshal(body, &data)
	if err != nil {
		ginBase.String(http.StatusOK, lib.ReturnErr("提交的数据格式错误!"))
	}
	if data["name"].(string) == "" || data["host_id"].(string) == "" {
		ginBase.String(http.StatusOK, lib.ReturnErr("缺少必须的参数!"))
	} else {
		res := model.GetWscBase(data)
		ginBase.String(http.StatusOK, lib.ReturnOK(res))
	}
}

// 检查识别名
func setWssBase() {
	var data map[string]interface{}
	body, _ := ginBase.GetRawData()
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("提交的数据格式错误！")
	}
	if data["name"].(string) == "" || data["host_id"].(string) == "" {
		log.Println("缺少必须的参数!")
	} else {
		str := model.CheckWscBaseName(data["name"].(string), data["host_id"].(string))
		var res = &returnCheckName{}
		// ok = 可用 | err = 换个名称 | other = 使用已存在的名称
		if str == "ok" {
			res.Status = str
			res.Name = data["name"].(string)
		} else if str == "error" {
			res.Status = str
		} else {
			res.Status = "server"
			res.Name = str
		}
		ginBase.String(http.StatusOK, lib.ReturnOK(res))
	}
}

// 修改识别名
func edtWssBase() {
	var data map[string]interface{}
	body, _ := ginBase.GetRawData()
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("提交的数据格式错误！")
	}
	if data["name"].(string) == "" || data["host_id"].(string) == "" {
		log.Println("缺少必须的参数!")
	} else {
		id, st := model.ExistWscBase(data["name"].(string), data["host_id"].(string))
		if st {
			//修改数据
			fmt.Println("修改的ID: ", id)
			fmt.Println(data["data"])
			if model.EditWscBase(id, data["data"]) {
				ginBase.String(http.StatusOK, lib.ReturnOK(data["data"]))
			} else {
				ginBase.String(http.StatusOK, lib.ReturnErr("修改失败"))
			}
		} else {
			ginBase.String(http.StatusOK, lib.ReturnErr("客户端不存在"))
		}
	}
}
