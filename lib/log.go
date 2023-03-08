package lib

import (
	"fmt"
	"log"
	"os"
)

var Logger *log.Logger
var ReceiveLog *log.Logger

func init() {
	var logPath = "./log"
	if !Exists(logPath) {
		//常见日志文件夹
		err := os.Mkdir(logPath, 0777)
		if err != nil {
			fmt.Println("创建日志目录错误: ", err)
		}
	}
	//Server.log
	logFile, err := os.OpenFile(logPath+"/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开Server日志文件异常: ", err)
	}
	Logger = log.New(logFile, "[Server]", log.Ldate|log.Ltime)

	//Receive.log
	receiveFile, err := os.OpenFile(logPath+"/receive.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开Receive日志异常: ", err)
	}
	ReceiveLog = log.New(receiveFile, "[Receive]", log.Ldate|log.Ltime)
}
