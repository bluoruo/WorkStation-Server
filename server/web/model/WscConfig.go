package model

type WscConfig struct {
	Model
	Name      string `json:"name"`
	Server    string `json:"server"`
	BaseApi   string `json:"base_api"`
	DdnsApi   string `json:"ddns_api"`
	RepoApi   string `json:"repo_api"`
	StatusApi string `json:"status_api"`
	ComApi    string `json:"com_api"`
	Remark    string `json:"remark"`
}

/* 获取记录 */

// ExistConfig 是否存在配置
func ExistConfig(name string) bool {
	var config WscConfig
	db.Select("id").Where("name = ?", name).First(&config)
	if config.ID > 0 {
		return true
	}
	return false
}

// 获取默认配置信息
func defConfig() (config WscConfig) {
	db.Where("id = 1").First(&config)
	return
}

// 获取用户配置信息
func userConfig(name string) (config WscConfig) {
	db.Where("name = ?", name).First(&config)
	return
}

// GetConfig 获取配置信息
func GetConfig(name string) (config WscConfig) {
	if ExistConfig(name) {
		return userConfig(name)
	} else {
		return defConfig()
	}
}

/* 分页相关 */

// GetConfigTotal 获取总数
func GetConfigTotal(maps interface{}) (count int64) {
	db.Model(&WscConfig{}).Where(maps).Count(&count)
	return
}

// GetConfigList 分页查询
func GetConfigList(pageNum int, pageSize int, maps interface{}) (config []WscConfig) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&config)
	return
}

/* 添加记录 */
func addConfig(data map[string]interface{}) bool {
	db.Create(&WscConfig{
		Name:      data["name"].(string),
		Server:    data["server"].(string),
		BaseApi:   data["base_api"].(string),
		DdnsApi:   data["ddns_api"].(string),
		RepoApi:   data["repo_api"].(string),
		StatusApi: data["status_api"].(string),
		ComApi:    data["com_api"].(string),
		Remark:    data["remark"].(string),
	})
	return true
}

/* 编辑记录 */
func editConfig(id int, maps interface{}) bool {
	db.Model(&WscConfig{}).Where("id = ?", id).Updates(maps)
	return true
}

/* 删除记录 */
func delConfig(id int) bool {
	db.Where("id = ?", id).Delete(&WscConfig{})
	return true
}
