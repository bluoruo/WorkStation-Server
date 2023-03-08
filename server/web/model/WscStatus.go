package model

type WscStatus struct {
	Model
	WsClientId int    `json:"ws_client_id"`
	Cpu        string `json:"cpu"`
	CpuTop     string `json:"cpu_top"`
	Mem        string `json:"mem"`
	MemTop     string `json:"mem_top"`
	Hd         string `json:"hd"`
	Net        string `json:"net"`
	Status     int    `json:"status"`
}

/* 获取记录 */

// ExistWscStatus 是否存在记录
func ExistWscStatus(wsId int) (int, bool) {
	var wscStatus WscStatus
	db.Select("id").Where("ws_client_id = ?", wsId).First(&wscStatus)
	if wscStatus.ID > 0 {
		return wscStatus.ID, true
	} else {
		return 0, false
	}
}

// GetWscStatus 获取客户端状态
func GetWscStatus(id int) (wscStatus WscStatus) {
	db.Where("id = ? ", id).First(&wscStatus)
	return
}

// AddWscStatus 新增客户端状态
func AddWscStatus(data map[string]interface{}) bool {
	db.Create(&WscStatus{
		WsClientId: data["ws_client_id"].(int),
		Cpu:        data["cpu"].(string),
		CpuTop:     data["cpu_top"].(string),
		Mem:        data["mem"].(string),
		MemTop:     data["mem_top"].(string),
		Hd:         data["hd"].(string),
		Net:        data["net"].(string),
		Status:     data["status"].(int),
	})
	return true
}

// EditWscStatus 更新客户端状态
func EditWscStatus(id int, maps interface{}) bool {
	db.Model(&WscStatus{}).Where("id = ?", id).Updates(maps)
	return true
}
