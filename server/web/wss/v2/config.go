package v2

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"imbot/lib"
	"imbot/server/web/model"
	"log"
	"net/http"
)

type structReturnConfig struct {
	Server    string
	BaseApi   string
	RepoApi   string
	DDnsApi   string
	StatusApi string
	ComApi    string
}

// WssConfig 返回 WorkStation Server 配置信息
func WssConfig(c *gin.Context) {
	//请求类型
	method := c.Request.Method
	var strType string
	if method == "POST" {
		strType = c.PostForm("type")
		switch strType {
		case "config":
			makeConfig(c)
			break
		case "ip":
			c.String(http.StatusOK, c.ClientIP())
		default:
			c.String(http.StatusOK, "Welcome to WorkStation Server.")
		}
	} else {
		strType = c.Query("type")
		switch strType {
		case "ip":
			c.String(http.StatusOK, c.ClientIP())
		case "update":
			c.String(http.StatusOK, updateInfo(c))
			break
		default:
			c.String(http.StatusOK, "Welcome to WorkStation Server.")
		}
	}

}

func makeConfig(c *gin.Context) {
	data := model.GetConfig("def")
	//log.Println(data)
	var res = &structReturnConfig{}
	res.Server = data.Server
	res.BaseApi = data.BaseApi
	res.DDnsApi = data.DdnsApi
	res.RepoApi = data.RepoApi
	res.StatusApi = data.StatusApi
	res.ComApi = data.ComApi
	//log.Println(res)
	//转成json
	resJson, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	//log.Println(string(resJson))
	str := lib.CmcheEncode(string(resJson))
	c.String(http.StatusOK, str)
}

func updateInfo(c *gin.Context) string {
	os := c.Query("os")
	arch := c.Query("arch")
	if os == "" || arch == "" {
		return lib.ReturnErr("缺少必要的参数！")
	}
	id, st := model.ExistNewVersion(os, arch)
	if st {
		data := model.NewVersion(id)
		res, _ := json.Marshal(data)
		return string(res)
	} else {
		return lib.ReturnErr("没有这个系统的版本！")
	}

}
