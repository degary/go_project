package main

import (
	"github.com/degary/logger"
	"time"
)

func initLogger(name string, logPath, logName string, level string) {
	config := make(map[string]string, 8)
	config["log_path"] = logPath
	config["log_name"] = logName
	config["log_level"] = level
	err := logger.InitLogger(name, config)
	//log = logger.NewConsoleLogger(level)
	if err != nil {
		return
	}
	logger.Debug("init logger success")
	return
}
func Run() {
	for i := 0; i < 10; i++ {
		logger.Warn("user server is running")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	initLogger("console", "F:\\log\\", "user_server", "debug")
	Run()
}
