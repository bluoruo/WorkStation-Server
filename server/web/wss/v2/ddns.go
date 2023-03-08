package v2

import (
	"github.com/gin-gonic/gin"
	"imbot/lib"
	"imbot/server/web/model"
	"net/http"
	"strconv"
)

func WssDDns(c *gin.Context) {
	strId := c.PostForm("ddns")
	intId, _ := strconv.Atoi(strId)
	//获取ddns接口参数信息
	id, st := model.ExistWscDDnsServer(intId)
	if st {
		res := model.GetDDnsApi(id)
		c.String(http.StatusOK, lib.ReturnOK(res))
	} else {
		c.String(http.StatusOK, lib.ReturnErr("没有这个ddns接口！"))
	}
}
