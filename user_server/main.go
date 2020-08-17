package main

import (
	"github.com/degary/logger"
)

func initLogger(name string, logPath, logName string, level string) {
	config := make(map[string]string, 8)
	config["log_path"] = logPath
	config["log_name"] = logName
	config["log_level"] = level
	config["log_split_type"] = "Size"
	err := logger.InitLogger(name, config)
	//log = logger.NewConsoleLogger(level)
	if err != nil {
		return
	}
	logger.Debug("init logger success")
	return
}
func Run() {
	for {
		logger.Warn("user server is running The id:25903,number:7956272103222810765,result:75")
		//time.Sleep(1 * time.Second)
	}
}

func main() {
	initLogger("file", "F:\\log\\", "user_server", "debug")
	Run()
}
