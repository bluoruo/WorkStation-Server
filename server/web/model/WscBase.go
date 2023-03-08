package model

type WscBase struct {
	Model
	WsAccountId   int    `json:"ws_account_id"`
	Name          string `json:"name"`
	HostName      string `json:"host_name"`
	HostId        string `json:"host_id"`
	Password      string `json:"password"`
	Type          int    `json:"type"`
	Tag           string `json:"tag"`
	DdnsServerId  int    `json:"ddns_server_id"`
	DdnsSubDomain string `json:"ddns_sub_domain"`
	DdnsType      int    `json:"ddns_type"`
	Services      string `json:"services"`
	Remark        string `json:"remark"`
	Status        int    `json:"status"`
}

/* 获取记录 */

// ExistWscBase 是否存在相同记录
func ExistWscBase(name string, hostId string) (int, bool) {
	var wscBase WscBase
	db.Select("id").Where("name = ? And host_id = ?", name, hostId).First(&wscBase)
	if wscBase.ID > 0 {
		return wscBase.ID, true
	} else {
		return 0, false
	}
}

// hostId 是否存在
func existWscBaseHostId(hostId string) bool {
	var wscBase WscBase
	db.Select("id").Where("host_id = ?", hostId).First(&wscBase)
	if wscBase.ID > 0 {
		return true
	} else {
		return false
	}
}

// 用户名是否存在
func existWscBaseName(name string) bool {
	var wscBase WscBase
	db.Select("id").Where("name = ? ", name).First(&wscBase)
	if wscBase.ID > 0 {
		return true
	} else {
		return false
	}
}

// 客户端 Base信息
func wscBaseInfo(name string, hostId string) (wscBase WscBase) {
	db.Where("name = ? And host_id = ?", name, hostId).First(&wscBase)
	return
}

// 识别名
func getWscBaseName(hostId string) (wscBase WscBase) {
	db.Where("host_id = ?", hostId).First(&wscBase)
	return
}

// CheckWscBaseName 取识别名
func CheckWscBaseName(name string, hostId string) string {
	if existWscBaseHostId(hostId) { //已存在设备
		//已有名称
		arr := getWscBaseName(hostId)
		return arr.Name
	} else { //不存在设备
		if existWscBaseName(name) { //是否重名
			return "error"
		}
	}
	return "ok"
}

// GetWscBase 获取基础信息
func GetWscBase(data map[string]interface{}) (wscBase WscBase) {
	if !existWscBaseHostId(data["host_id"].(string)) { //不存在
		addWscBase(data) //新增记录
	}
	// 默认返回
	return wscBaseInfo(data["name"].(string), data["host_id"].(string))
}

/* 添加记录 */
func addWscBase(data map[string]interface{}) bool {
	db.Create(&WscBase{
		Name:     data["name"].(string),
		HostName: data["host_name"].(string),
		HostId:   data["host_id"].(string),
		Status:   2,
	})
	return true
}

/* 编辑记录 */

// EditWscBase 编辑基础信息
func EditWscBase(id int, maps interface{}) bool {
	db.Model(&WscBase{}).Where("id = ?", id).Updates(maps)
	return true
}
