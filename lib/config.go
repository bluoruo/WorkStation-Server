package lib

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var ConfigFile = ".env" //配置文件

type serverConfig struct {
	Listen   string
	HttpPort string
}
type sqlConfig struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type redisConfig struct {
	Host     string
	Port     string
	Password string
}

type openAiConfig struct {
	ApiKey       string
	Organization string
}

var ServerConfig = &serverConfig{}
var SqlConfig = &sqlConfig{}
var RedisConfig = &redisConfig{}
var OpenAiConfig = &openAiConfig{}

// LoadConfig 加载配置
func LoadConfig() bool {
	cfg, err := ini.Load(ConfigFile)
	if err != nil {
		fmt.Println("配置文件加载错误: ", err)
		return false
	}
	//加载Server配置
	err = cfg.Section("Server").MapTo(ServerConfig)
	if err != nil {
		fmt.Println("获取Server Config错误: ", err)
		return false
	}
	//加载 SQL配置
	err = cfg.Section("SQL").MapTo(SqlConfig)
	if err != nil {
		fmt.Println("获取SQL Config错误: ", err)
		return false
	}
	//加载 Redis配置
	err = cfg.Section("Redis").MapTo(RedisConfig)
	if err != nil {
		fmt.Println("获取Redis Config错误: ", err)
		return false
	}
	//加载 OpenAi配置
	err = cfg.Section("OpenAI").MapTo(OpenAiConfig)
	if err != nil {
		fmt.Println("获取OpenAI Config错误: ", err)
		return false
	}
	return true
}
