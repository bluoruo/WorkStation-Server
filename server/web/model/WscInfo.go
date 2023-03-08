package model

type WscInfo struct {
	Model
	WsClientId int    `json:"ws_client_id"`
	WscVer     string `json:"wsc_ver"`
	HostName   string `json:"host_name"`
	OS         string `json:"os"`
	Arch       string `json:"arch"`
	Kernel     string `json:"kernel"`
	CPU        string `json:"cpu"`
	Mem        string `json:"mem"`
	DISK       string `json:"disk"`
	Remark     string `json:"remark"`
}

/* 获取记录 */

// ExistWscIdWscInfo 是否存在记录
func ExistWscIdWscInfo(wsId int) (int, bool) {
	var wscInfo WscInfo
	db.Select("id").Where("ws_client_id = ?", wsId).First(&wscInfo)
	if wscInfo.ID > 0 {
		return wscInfo.ID, true
	} else {
		return 0, false
	}
}

/* 添加记录 */

// AddWscInfo 新增
func AddWscInfo(data map[string]interface{}) bool {
	db.Create(&WscInfo{
		WsClientId: data["ws_client_id"].(int),
		WscVer:     data["wsc_ver"].(string),
		HostName:   data["host_name"].(string),
		OS:         data["os"].(string),
		Arch:       data["arch"].(string),
		Kernel:     data["kernel"].(string),
		CPU:        data["cpu"].(string),
		Mem:        data["mem"].(string),
		DISK:       data["disk"].(string),
	})
	return true
}

/* 修改记录 */

// EditWscInfo 修改
func EditWscInfo(id int, maps interface{}) bool {
	db.Model(&WscInfo{}).Where("id = ?", id).Updates(maps)
	return true
}
