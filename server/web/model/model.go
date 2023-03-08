package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"imbot/lib"
	"log"
	"sync"
)

var (
	db       *gorm.DB
	connOnce sync.Once
)

type Model struct {
	ID         int `gorm:"primaryKey" json:"id"`
	AddTime    int `gorm:"autoCreateTime" json:"add_time"`
	UpdateTime int `gorm:"autoUpdateTime" json:"update_time"`
}

// Model 默认连接
func connDb() {
	connOnce.Do(func() {
		lib.LoadConfig()
		var err error
		//dsn := "harry:1q2w3e4r..@tcp([240e:379:1b2e:2f00::b9d]:3306)/wss?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := lib.SqlConfig.User + ":" + lib.SqlConfig.Password +
			"@tcp(" + lib.SqlConfig.Host + ":" + lib.SqlConfig.Port + ")/" + lib.SqlConfig.Database +
			"?charset=utf8mb4&parseTime=True&loc=Local"
		fmt.Println(dsn)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //不需要s结尾
			}})
		if err != nil {
			log.Println(err)
		}
	})
}

func init() {
	connDb()
}
