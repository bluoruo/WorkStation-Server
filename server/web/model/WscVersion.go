package model

type WscVersion struct {
	Model
	Ver     string `json:"ver"`
	Md5     string `json:"md5"`
	Os      string `json:"os"`
	Arch    string `json:"arch"`
	DownUrl string `json:"down_url"`
	Status  string `json:"-"`
}

/* 取出 */

func ExistNewVersion(os, arch string) (int, bool) {
	var wscVersion WscVersion
	db.Select("id").Where("os = ? And arch = ? And status = 1", os, arch).First(&wscVersion)
	if wscVersion.ID > 0 {
		return wscVersion.ID, true
	} else {
		return 0, false
	}
}

// NewVersion 取出操作系统的版本信息
func NewVersion(id int) (wscVersion WscVersion) {
	db.Where("id = ? ", id).First(&wscVersion)
	return
}
