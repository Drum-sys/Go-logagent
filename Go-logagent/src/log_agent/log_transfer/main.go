package main

import (
	"github.com/astaxie/beego/logs"
)

func main()  {

	err := InitConf("ini", "./conf/logs_transfer.conf")
	if err != nil {
		panic(err)
		return
	}
	logs.Debug("init conf success")

	err = InitLogger(logConf.LogPath, logConf.LogLevel)
	if err != nil {
		panic(err)
		return
	}
	logs.Debug("init logger success")

	err = InitKafka(logConf.kafkaAddr, logConf.kafkaTopic)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}
	logs.Debug("init kafka success")

	err = InitES(logConf.ESAddr)
	if err != nil {
		logs.Error("init es failed, err:%v", err)
		return
	}
	logs.Debug("init es success")

	err = Run()
	if err != nil {
		logs.Error("RUn failed, err:%v", err)
		return
	}

	logs.Warn("warning log_transfer exited, err:%v", err)

}
