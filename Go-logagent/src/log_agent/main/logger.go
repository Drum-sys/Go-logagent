package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func convertLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func InitLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = appConf.logPath
	config["level"] = convertLevel(appConf.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("init logger  failed err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}