package model

type WscExec struct {
	Model
	WsClientId int `json:"ws_client_id"`
	ShellId    int `json:"shell_id"`
	Status     int `json:"status"`
}

/* 获取记录 */

// ExistWscExecShellId 获取执行命令ID
func ExistWscExecShellId(wsId int) (int, bool) {
	var wscExec WscExec
	db.Where("shell_id = ? And status = 1 ", wsId).First(&wscExec)
	if wscExec.ID > 0 {
		return wscExec.ShellId, true
	} else {
		return 0, false
	}
}
