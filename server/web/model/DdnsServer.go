package model

type DdnsServer struct {
	Model
	Name         string `json:"name"`
	Account      string `json:"account"`
	Password     string `json:"password"`
	ApiKey       string `json:"api_key"`
	ApiSecret    string `json:"api_secret"`
	ApiToken     string `json:"api_token"`
	MasterDomain string `json:"master_domain"`
	LostTime     int    `json:"lost_time"`
	Url          string `json:"url"`
	ClassId      int    `json:"class_id"`
	Remark       string `json:"remark"`
	Status       int    `json:"status"`
}

/* 获取记录 */

// ExistWscDDnsServer 是否存在相同记录
func ExistWscDDnsServer(id int) (int, bool) {
	var ddnsServer DdnsServer
	db.Where("id = ? ", id).First(&ddnsServer)
	if ddnsServer.ID > 0 {
		return ddnsServer.ID, true
	} else {
		return 0, false
	}
}

// GetDDnsApi 获取所有信息
func GetDDnsApi(id int) (ddnsServer DdnsServer) {
	db.Where("id = ? ", id).First(&ddnsServer)
	return
}
