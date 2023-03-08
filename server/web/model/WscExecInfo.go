package model

type WscExecInfo struct {
	Model
	Name   string `json:"name"`
	Type   int    `json:"type"`
	Com    string `json:"com"`
	Remark string `json:"remark"`
	Status int    `json:"status"`
}

/* 获取记录 */

// ExistWscExecInfo 获取执行命令ID
func ExistWscExecInfo(id int) (int, bool) {
	var wscExecInfo WscExecInfo
	db.Where("id = ? And status = 1 ", id).First(&wscExecInfo)
	if wscExecInfo.ID > 0 {
		return wscExecInfo.ID, true
	} else {
		return 0, false
	}
}

// GetWscExecInfo 获取指令
func GetWscExecInfo(id int) (wscExecInfo WscExecInfo) {
	db.Where("id = ? ", id).First(&wscExecInfo)
	return
}
