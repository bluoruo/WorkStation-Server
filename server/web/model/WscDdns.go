package model

type WscDdns struct {
	Model
	WsClientId int    `json:"ws_client_id"`
	WsToken    string `json:"ws_token"`
	DdnsId     int    `json:"ddns_id"`
	Domain     string `json:"domain"`
	IpType     int    `json:"ip_type"`
	Ip         string `json:"ip"`
	Status     int    `json:"status"`
}

/* 查询记录 */
func getWscDDnsId(name string) int {
	var wscDDns WscDdns
	db.Select("id").Where("id = ?", name).First(&wscDDns)
	return wscDDns.ID
}
