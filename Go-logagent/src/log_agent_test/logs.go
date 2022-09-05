package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

// 热加载

func main()  {
	config := make(map[string]interface{})
	config["filename"] = "/home/drum/go_workspace/src/laonanhai/src/log_agent/1.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	logs.Debug("this is a test, my name is %s", "ljw")
	logs.Trace("this is a test, my name is %s", "ljp")
	logs.Warning("this is a test, my name is %s", "lyp")

}
